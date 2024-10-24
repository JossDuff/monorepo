# Infra

This folder is for scripts and configuration for our hosted sigil cluster.

The cluster is hosted on digital ocean and consists of 4 pods: `op-geth`, `op-node`, `op-batcher`, and `op-proposer`.  The `op-geth` binary is from our [ethereum-optimism/op-geth fork](https://github.com/unattended-backpack/op-geth/) and the rest of the binaries `op-node`, `op-batcher`, and `op-proposer` are from our [ethereum-optimism/optimism fork](https://github.com/unattended-backpack/optimism).

This cluster is primarily for development work at the moment.

## deployments.yml

This is the cluster deploy script.  It can be run with `kubectl apply -f deployments.yml`.
