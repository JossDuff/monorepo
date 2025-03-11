use alloy_primitives::{hex, B256};
use anyhow::{Context, Result};
use log::{error, info};
use op_succinct_client_utils::{boot::BootInfoStruct, types::u32_to_u8};
use op_succinct_host_utils::{
    fetcher::{CacheMode, OPSuccinctDataFetcher, RunContext},
    get_agg_proof_stdin, get_proof_stdin, start_server_and_native_client, ProgramType,
};
use op_succinct_local_proposer::SpanProofRequest;
use sp1_sdk::{utils, HashableKey, Prover, ProverClient, SP1Proof, SP1ProofWithPublicValues};
use std::{fs, str::FromStr};

pub const RANGE_ELF: &[u8] = include_bytes!("../../../elf/range-elf");
pub const AGG_ELF: &[u8] = include_bytes!("../../../elf/aggregation-elf");

const CHECKPOINTED_BLOCKHASH: &str = "0x00";

#[tokio::main]
async fn main() -> Result<()> {
    //  start: 27772, end: 28072
    let l2_start_block = 27772;
    let l2_end_block = 28072;

    info!(
        "Will be proving from block {} to block {}",
        l2_start_block, l2_end_block
    );

    // dummy payload
    let payload = SpanProofRequest {
        start: l2_start_block,
        end: l2_end_block,
    };

    utils::setup_logger();

    dotenv::dotenv().ok();

    let prover = ProverClient::builder().cuda().build();
    let (range_pk, range_vk) = prover.setup(RANGE_ELF);
    // let (_agg_pk, agg_vk) = prover.setup(AGG_ELF);
    let multi_block_vkey_u8 = u32_to_u8(range_vk.vk.hash_u32());
    let _range_vkey_commitment = B256::from(multi_block_vkey_u8);
    // let _agg_vkey_hash = B256::from_str(&agg_vk.bytes32()).unwrap();
    let fetcher = match OPSuccinctDataFetcher::new_with_rollup_config(RunContext::Docker).await {
        Ok(f) => f,
        Err(e) => {
            error!("Failed to create data fetcher: {}", e);
            todo!();
        }
    };

    let host_args = fetcher
        .get_host_args(
            payload.start,
            payload.end,
            None,
            ProgramType::Multi,
            CacheMode::DeleteCache,
        )
        .await
        .context("Failed to get host CLI args")?;

    let mem_kv_store = start_server_and_native_client(host_args).await?;

    let sp1_stdin = get_proof_stdin(mem_kv_store).context("Failed to get proof stdin")?;

    info!("executing span proof");

    let start_time = tokio::time::Instant::now();
    let proof = prover
        .prove(&range_pk, &sp1_stdin)
        .compressed()
        .run()
        .unwrap();

    info!("done with span proof");
    let minutes = start_time.elapsed().as_secs_f64() / 60.0;
    info!("Time to compute: {} minutes", minutes);
    panic!("done");

    // Create a proof directory for the chain ID if it doesn't exist.
    let proof_dir = "proofs/".to_string();
    if !std::path::Path::new(&proof_dir).exists() {
        fs::create_dir_all(&proof_dir).unwrap();
    }
    let proof_path = format!("{}/{}-{}.bin", proof_dir, l2_start_block, l2_end_block);

    // Save the proof to the proof directory corresponding to the chain ID.
    proof.save(&proof_path).expect("saving proof failed");

    let (proofs, boot_infos) = load_aggregation_proof_data(proof_path);

    info!("loaded saved proof");

    let l1_head_string = CHECKPOINTED_BLOCKHASH
        .strip_prefix("0x")
        .context("Invalid L1 head format: missing 0x prefix")?;
    let l1_head_bytes =
        hex::decode(l1_head_string).context("Failed to decode L1 head hex string")?;

    let l1_head: [u8; 32] = l1_head_bytes
        .clone()
        .try_into()
        .expect("Invalid L1 head length, expected 32 bytes");

    let fetcher = OPSuccinctDataFetcher::new_with_rollup_config(RunContext::Docker)
        .await
        .context("failed to create fetcher")?;

    let headers = fetcher
        .get_header_preimages(&boot_infos, l1_head.into())
        .await
        .context("Failed to get header preimages")?;

    let sp1_stdin = get_agg_proof_stdin(proofs, boot_infos, headers, &range_vk, l1_head.into())
        .context("Failed to get agg proof stdin")?;

    let (agg_pk, _) = prover.setup(AGG_ELF);
    // println!("Aggregate ELF Verification Key: {:?}", agg_vk.vk.bytes32());

    info!("executing agg proof");
    let _proof_res = prover
        .prove(&agg_pk, &sp1_stdin)
        .groth16()
        .run()
        .expect("proving failed");

    info!("done with agg proof");

    Ok(())
}

/// Load the aggregation proof data.
fn load_aggregation_proof_data(proof_path: String) -> (Vec<SP1Proof>, Vec<BootInfoStruct>) {
    if fs::metadata(&proof_path).is_err() {
        panic!("Proof file not found: {}", proof_path);
    }

    let mut deserialized_proof =
        SP1ProofWithPublicValues::load(proof_path).expect("loading proof failed");

    // The public values are the BootInfoStruct.
    let boot_info = deserialized_proof.public_values.read();

    (vec![deserialized_proof.proof], vec![boot_info])
}
