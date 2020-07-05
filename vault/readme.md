# Vault

Some shorter infomation about Vault and some fundamental concepts are also written [here](../security/cryptography.md#How-to-store-keys-to-encrypt-and-decrypt).

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

We can store the vault address and token in the environment, so we can start testing

```sh
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN='s.IHT9CuwqcZTPXlEw1d5hsUtD'
```

When a dev server starts it works over http and start with an unsealed state.

There 3 ways to interact with the vault:

- via cli (using the `vault` command we installed with `brew`)
- via http
- via the dedicated UI

## Vault UI

In out case, we login via the browser using the `VAULT_ADDR`, http://127.0.0.1:8200

## Inserting values in the vault

### CLI

```sh
vault kv put secret/my-test age=25
```

This line:

- selects `secret` as the Secrets Engine (note: `secret` didn't need to be created because it's there by default, but other engines can be created)
- creates a new secret named `my-test`, or selects it if present
- creates a new key/value pair `{age: 25}`

```sh
vault kv get secret/my-test

====== Metadata ======
Key              Value
---              -----
created_time     2020-07-05T11:37:08.902356Z
deletion_time    n/a
destroyed        false
version          1

===== Data =====
Key       Value
---       -----
answer    15
```

### HTTP

Vault exposes an API we can use to interact with our engines.

```sh
curl --location --request POST 'http://127.0.0.1:8200/v1/secret/data/my-test2' \
--header 'X-Vault-Token: s.IHT9CuwqcZTPXlEw1d5hsUtD' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": {
        "age": 26
    }
}'
```

We can then GET http://127.0.0.1:8200/v1/secret/data/my-test2.

```json
{
  "request_id": "078daf1f-e129-08d0-d9b1-a46f142e230a",
  "lease_id": "",
  "renewable": false,
  "lease_duration": 0,
  "data": {
    "data": {
      "age": 26
    },
    "metadata": {
      "created_time": "2020-07-05T11:47:52.435608Z",
      "deletion_time": "",
      "destroyed": false,
      "version": 1
    }
  },
  "wrap_info": null,
  "warnings": null,
  "auth": null
}
```

## Updating secrets

We use the same method as creating

```sh
vault kv put secret/my-test answer=35

Key              Value
---              -----
created_time     2020-07-05T11:58:55.15214Z
deletion_time    n/a
destroyed        false
version          2
```

Note the version bump. Vault keep by default the last 10 versions.

## Next steps

[Working with Secrets](./working-with-secrets):

- secrets lifecycles (create, read, update, delete)
- what are secrets engines?
-
