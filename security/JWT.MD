# JWT

A JSON webtoken is a standardised container format that is used to transfer information securely

## Deserialised token

When deserialised a JWT is composed of 2 JSON objects:

- a header
- a payload

## Header

The header contains the algorithm that is used to sign and encrypt the token:

```json
{
  "alg": "HS256"
}
```

A token can also be **unsecured**, ie not encrypted:

```json
{
  "alg": "none"
}
```

### Payload

Claim is a term drawn from the spec and defines the key/value pair that conveys a piece of info.
Claims can be standard or user defined. Some standard ones are:

- `iss`: issuer, that identifies the party who issued the token
- `sub`: subject, or the party that the token has info about
- `aud`: the recipients of the token
- `exp`: the expiration date
- `iat`: time the token was issued

Other claims can be any type of info the service need to transfer:

```json
{
  "name": "Andrew",
  "email": "andrew@email.com"
}
```

## Serialised token

A token is serialised in order to be transferred. Format is:

```
[ Header ].[ Payload ].[ Signature ]
```

Header and payload are always there, but signature might not, if the token is unsecured.

## Serialisation

A token is serialised with [Base64URL](https://brockallen.com/2014/10/17/base64url-encoding/) encoding, a variation of base64 that uses url safe characters.
Encoding can be represented like:

```js
function encode(h, p) {
  const header = base64UrlEncode(JSON.stringify(h));
  const payload = base64UrlEncode(JSON.stringify(p));
  return `${header}.${payload}`;
}
```

Decoding:

```js
function decode(jwt) {
  const [headerB64, payloadB64] = jwt.split(".");
  const headerStr = base64UrlDecode(headerB64);
  const payloadStr = base64UrlDecode(payloadB64);
  return {
    header: JSON.parse(headerStr),
    payload: JSON.parse(payloadStr)
  };
}
```

## Securing a token: signing vs encryption

A JWT can be signed and encrypted.

### Signature

A signature makes tampering with the token impossible.
The most common signing algo is HMAC. It combines a payload with a secret using a hash function (say, SHA-256) and produces a signature that will form the 3rd part of a serialised token (`[ Signature ]` above).
It's important to note that a signed token can still be read by anyone, the signature prevents only tampering with the data.

### Encryption

Encryption is the (optional) step to make the token unreadable to 3rd parties.
An encrypted token as 5 (rather than 3) parts to it.

More info [here](https://medium.com/ag-grid/a-plain-english-introduction-to-json-web-tokens-jwt-what-it-is-and-what-it-isnt-8076ca679843).
