# Vault: working with secrets

## Lifecycles: getting and setting

```sh
# put a new value
vault kv put secret/my-test2 age=44
# update it and also add a new k/v pair
vault kv put secret/my-test2 age=22 country=uk
# get a key as JSON
vault kv get -format=json secret/my-test2 age

$ {
  "request_id": "7e57ae49-8074-b019-4644-9eb61400c0fa",
  "lease_id": "",
  "lease_duration": 0,
  "renewable": false,
  "data": {
    "data": {
      "age": "22",
      "country": "uk"
    },
    "metadata": {
      "created_time": "2020-07-05T12:59:02.970821Z",
      "deletion_time": "",
      "destroyed": false,
      "version": 3
    }
  },
  "warnings": null
}
```

## What are secrets engines?

SE are the mechanism that Vault uses to store secrets. There are a number of different types of engines depending on the types of secrets that we want to store.

> Secrets engines are components which store, generate, or encrypt data. Secrets engines are incredibly flexible, so it is easiest to think about them in terms of their function. Secrets engines operate by receiving API calls that meet this interface. The calls receive data from the caller, take some action on that data, and they return a result.
>
> Some secrets engines simply store and read data - like encrypted Redis/Memcached. Other secrets engines connect to other services and generate dynamic credentials on demand. Other secrets engines provide encryption as a service, totp generation, certificates, and much more.

Since SEs map to a path in the vault

> To the user, secrets engines behave similar to a virtual filesystem, supporting operations like read, write, and delete.

Paths enable introspection through `path-help`.

To see all engines:

```sh
vault secrets list [--format=json]
```
