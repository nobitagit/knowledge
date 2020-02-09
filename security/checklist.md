# Security checklist

This is a list of steps to ensure a web application is secure. This is **not** exhaustive, it is meant to be a collection of some of tips that should always be part of a basic setup.

## Notable references

- [Mozilla web security cheat sheet](https://infosec.mozilla.org/guidelines/web_security) - useful guideline with also an evaluation of "cost" vs benefit of each point.

## Checklist

### Basic

- Enable [HSTS](./https.md#HSTS)
- use [CSP](./README#CSP)

### More advanced

- Once sure all the domain and the subdomains support https fully, enable [hsts preloading](./https.md#HSTS-preloading).
- [HPKP](./https.md#HPKP) (Http public key pinning)
