# Security: Cryptography

## Hash algorithms or digests

One way functions that given an input produce a hash as output.

## What is a hash?

A hash is:

- a **fixed-length** string
- deterministic
- collision resistant
- unidirectional

This info is based on [Node's Crypto documentation](https://nodejs.org/en/knowledge/cryptography/how-to-use-crypto-module/).

### Fixed length

For example, SHA-256 hashes are **always** 256 bits long whether the input data is a few bits or a few gigabytes.

### Deterministic

Same input === same output. This makes hashes useful for [checksums](https://en.wikipedia.org/wiki/Checksum).

### Collision resistant

Hash algos are designed to make it unlikely that two different inputs yield the same hash.

### Unidirectional

It should be very easy to go from input to hash, but very hard to go the opposite direction.

## Hashing algorithms

You can create a hash and then try to see if [Hash Toolkit](https://hashtoolkit.com/reverse-hash) can decrypt it. The most common passwords will be guessable by that tool.
You can try:

- "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f" - "password" through SHA-256, will be guessed
- "643c5b36edb46869d5471e49fb12c23a6285b46e765105db14ce8d84dd7d3503" - "aziT7662" again SHA-256, will not be found

(Hash toolkit is using rainbow tables).

### MD5

- Message Digest 5
- Cryptographic hash function created in 1992
- Many vulnerabilities have been discovered so far, so it's no longer safe for protecting passwords

## HMAC

- Hash-based Message Authentication Code
- defines the process of applying a hash to both the data and a secret key
- the result is one single hash

So 2 inputs (data and key), 1 output (hashed output).

It checks for **integrity** AND **authenticity** (the data and the key are validated).
