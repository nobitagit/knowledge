# Node: Crypto

## Example with protecting passwords

Use the [crypto module](https://nodejs.org/en/knowledge/cryptography/how-to-use-crypto-module/).
**Crypto** is a wrapper around Open SSL. Open SSL is an open source library that contains the implementation of SSL and TLS. In fact:

> The core library, written in the C programming language, implements basic cryptographic functions and provides various utility functions. **Wrappers allowing the use of the OpenSSL library in a variety of computer languages are available**.
>
> -- [source](https://en.wikipedia.org/wiki/OpenSSL)

The hash functions (or digests) that are available depend on the system the program is being executed on. To see you available hashing algorithms you can type on you CLI:

```sh
openssl list-message-digest-algorithms
```

> One of the most common hash algorithms is SHA-256. Older popular types like SHA-1 or MD5 **are not secure any more and should not be used**.

Here's how crypto works:

```js
const crypto = require("crypto");

var hash = crypto.createHash("sha256"); // or md5 or any other we need

// we then push data to the hash (it will be stored, but not yet passed throught he function)
var data = hash.update("my-password");
// we finally invoke the function and specify the output format (binary/hex/base64)
data.digest("hex");
// => '6fa2288c361becce3e30ba4c41be7d8ba01e3580566f7acc76a7f99994474c46'
```

Digest can be invoked/chained so data can be pushed incrementally (like via a buffer).

## HMAC

[What is HMAC?](../security/cryptography.md#HMAC).

```js
require("crypto")
  .createHmac("sha256", "my-secret-key")
  .update("my-password")
  .digest("hex");
// '6448578a3b74d91ca1e0f4cc84db14685ad6586d717cacae2a6ec8b6aad59030'
```

## Tools for symmetric encryption

Node exposes `createCipheriv` and `CreateDecipheriv`.

```js
const crypto = require("crypto");

const algo = "aes-256-cbc";
const password = "my-safe-password";
const salt = crypto.randomBytes(32);
console.log(`Salt is : ${salt.toString("hex")}`);

const key = crypto.scryptSync(password, salt, 32);

// initialisation vector
const iv = crypto.randomBytes(16);
const cipher = crypto.createCipheriv(algo, key, iv);

const dataToEncrypt = "My data";

let encrypted = cipher.update(dataToEncrypt, "utf8", "hex");
encrypted += cipher.final("hex");

console.log(`In the DB you'll store: ${encrypted}`);

// DECRYPT
const decipher = crypto.createDecipheriv(algo, key, iv);

let decrypted = decipher.update(encrypted, "hex", "utf8");
decrypted += decipher.final("utf8");

console.log(decrypted === dataToEncrypt); // true
```
