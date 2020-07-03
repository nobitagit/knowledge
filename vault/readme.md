# Vault

Some shorter infom about Vault and some fundamental concepts are also written [here](../security/cryptography.md#How-to-store-keys-to-encrypt-and-decrypt).

```sh
brew install vault -y
```

To start the dev server

```sh
vault server -dev
```

Result:

```sh
▶ vault server -dev
==> Vault server configuration:

             Api Address: http://127.0.0.1:8200
                     Cgo: disabled
         Cluster Address: https://127.0.0.1:8201
              Listener 1: tcp (addr: "127.0.0.1:8200", cluster address: "127.0.0.1:8201", max_request_duration: "1m30s", max_request_size: "33554432", tls: "disabled")
               Log Level: info
                   Mlock: supported: false, enabled: false
           Recovery Mode: false
                 Storage: inmem
                 Version: Vault v1.4.2
             Version Sha: 18f1c494be8b06788c2fdda1a4296eb3c4b174ce+CHANGES

WARNING! dev mode is enabled! In this mode, Vault runs entirely in-memory
and starts unsealed with a single unseal key. The root token is already
authenticated to the CLI, so you can immediately begin using Vault.

You may need to set the following environment variable:

    $ export VAULT_ADDR='http://127.0.0.1:8200'

The unseal key and root token are displayed below in case you want to
seal/unseal the Vault or re-authenticate.

Unseal Key: YLDGet+Nq/06UOhwPEaQkyOADE05wEbvr7I/CGf/HLY=
Root Token: s.IHT9CuwqcZTPXlEw1d5hsUtD

Development mode should NOT be used in production installations!

==> Vault server started! Log data will stream in below:
```

```sh
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN='s.IHT9CuwqcZTPXlEw1d5hsUtD'
```
