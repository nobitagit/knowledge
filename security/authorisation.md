# Authorisation

## Auth with JWT

### Disadvantages of cookies

- **they don't like CORS**: they can only be used by the same origin.

> The same origin policy for cookies as it is currently implemented may be too permissive in some cases. For example, some country-code top level domains have second level subdomains that act as generic, functional top level domains in their respective hierarchies. In many cases current domain name matching rules allow sharing of cookies with all sites in such domains [6, 15]. Even in domains with more “administrative affinity” some hosts may want to interact via cookies without necessarily sharing information with, or receiving information from, other peer hosts. This unwanted sharing may result in cookie leakage (cookies are sent to unauthorized Web servers) or cookie spoofing (cookies are inadvertently or maliciously set by unauthorized Web servers) [1].
>
> From these examples it should be clear that cookie sharing across domain boundaries is a desirable feature in Web applications and middleware.
>
> [Reference](https://jisajournal.springeropen.com/articles/10.1186/1869-0238-4-13)
