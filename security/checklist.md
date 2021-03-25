# Security checklist

This is a list of steps to ensure a web application is secure. This is **not** exhaustive, it is meant to be a collection of some of tips that should always be part of a basic setup.

## Notable references

- [Mozilla web security cheat sheet](https://infosec.mozilla.org/guidelines/web_security) - useful guideline with also an evaluation of "cost" vs benefit of each point.
- [API keys checker](https://github.com/streaak/keyhacks)

## Checklist

### Basic

- Enable [HSTS](./https.md#HSTS)
- Use [CSP](./README#CSP)

#### API Keys

- Check they are not exposed to the clients
- If they can't be removed from the clients, check if you can restrict them. See here and example for [Google API keys](https://developers.google.com/maps/api-key-best-practices)

#### Graphql

- Disable [introspection in production](https://lab.wallarm.com/why-and-how-to-disable-introspection-query-for-graphql-apis/)

#### Sessions

- Ensure you destroy the open sessions after a password reset. To test:
   - login
   - reset password
   - open a new incognito window and log in with new creds, so you see them working
   - first window should be logged out

### More advanced

- Once sure all the domain and the subdomains support https fully, enable [hsts preloading](./https.md#HSTS-preloading).
- [HPKP](./https.md#HPKP) (Http public key pinning)
