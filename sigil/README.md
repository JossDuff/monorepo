# Sigil

Extracted source and build code for all the binaries that make up Sigil.

# /op-succinct

original repo: <https://github.com/succinctlabs/op-succinct/>
forked version: `op-succinct-v1.0.0-rc6`

Contains binaries `op-succinct-proposer` and `op-succinct-server` ("proof server"").
`op-succinct-proposer` monitors the L2 chain and periodically sends a request for
a proof to `op-succinct-server` in the form of a block range.  `op-succinct-server`
is a small server that accepts requests for proofs and delegates the actual proof
generation to either a local CUDA accelerated prover (a docker image that is
started by the server) or the sp1 prover network (currently only runs the CUDA
prover).  When a proof is done, it returns the proof to `op-succinct-proposer`
who then posts the proof on-chain.

`op-succinct/proposer/succinct/bin/single-block-proof.rs` is a binary used for
debugging isolated proof requests and isn't used in prod.

# TODO /op-geth

Contains the execution node binary for the op-stack chain.

# TODO /optimism

Contains the op-stack chain binaries `op-batcher` and `op-node`.
