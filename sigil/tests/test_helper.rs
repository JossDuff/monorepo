use anyhow::{anyhow, Context, Result};
use testcontainers::{
    core::{ContainerAsync, IntoContainerPort, WaitFor},
    runners::AsyncRunner,
    GenericImage, ImageExt,
};

// TODO: there are a lot of constansts in here.  They will either be params or real CONST vars.

pub struct SigilTestInstance {
    container: ContainerAsync<GenericImage>,
    host_port: u16,
    reqwest_client: reqwest::Client,
}

impl SigilTestInstance {
    pub async fn new(config_file: &str, port: u16) -> Self {
        let config_toml_path = format!("test_configs/{config_file}");

        // TODO: how can we pass in different sigil.toml files to test different configurations?
        // TODO: also how can we run sigil in the container with RUST_LOG=priory=trace,warn ?
        let container = GenericImage::new("sigil", "dev")
            .with_exposed_port(port.tcp())
            .with_wait_for(WaitFor::message_on_stdout("Sigil is alive."))
            .with_env_var("RUST_LOG", "priory=trace,warn")
            .with_env_var("CONFIG_TOML_PATH", config_toml_path)
            .start()
            .await
            .expect("Failed to start sigil container");

        let host_port = container
            .get_host_port_ipv4(port)
            .await
            .expect("Failed to get host port");

        let reqwest_client = reqwest::Client::new();

        Self {
            container,
            host_port,
            reqwest_client,
        }
    }

    // make an rpc call and dump container logs if the response doesn't contain some expected value
    pub async fn rpc_with_expected(
        &self,
        method: &str,
        params: Option<&str>,
        expected: &str,
    ) -> Result<()> {
        if let Err(e) = async {
            let response = self.rpc(method, params).await?;

            if !response.contains(expected) {
                anyhow::bail!(
                    "Response does not contain expected text. \nexpected: {}\nactual: {}\n",
                    expected,
                    response
                );
            }
            Ok(())
        }
        .await
        {
            let logs = self.get_container_logs().await;
            return Err(anyhow!("Test failed: {}\nContainer logs:\n{}", e, logs));
        } else {
            Ok(())
        }
    }

    // makes the rpc call and returns just the body
    pub async fn rpc(&self, method: &str, params: Option<&str>) -> Result<String> {
        let params = params.unwrap_or("");
        let response = self
            .reqwest_client
            .post(&format!("http://localhost:{}", self.host_port))
            .json(&serde_json::json!({
                "jsonrpc": "2.0",
                "method": method,
                "params": [params],
                "id": 1
            }))
            .send()
            .await
            .context(format!(
                "Failed to send rpc request for {} with params {}",
                method, params
            ))?;

        response.text().await.context("Failed to get response body")
    }

    pub async fn get_container_logs(&self) -> String {
        match self.container.stdout_to_vec().await {
            Ok(log_bytes) => String::from_utf8_lossy(&log_bytes).into_owned(),
            Err(e) => format!("Failed to retrieve container logs: {}", e),
        }
    }
}
