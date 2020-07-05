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
