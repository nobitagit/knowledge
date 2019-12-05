# Networking

- [Anatomy of a browser web request](./README.MD#Anatomy of a browser web reqeust)

## DNS records

DNS (Domain name system) is the internet-wide phonebook that tells browsers or any other client how to resolve a specific domain name to an IP address of a server the can open a connection to.

Websites have 3 kinds of records today:

- A Record: returns an IPv4 address of a server
- AAAA Record: returns the IPv6 address of a server
- CNAME Record: returns another domain name the browser should look up instead.

note - there's also **TXT records** that allows associating arbitrary text with a host, such as readable info reg the server, network or data center. It can also contain machine readable info. More [here](https://en.wikipedia.org/wiki/TXT_record).

So, if we a set up is:

```md
// config for docs.example.com
CNAME => mydocs.myhost.com
// config for mydocs.myhost.com
A record => 185.28.186.242
```

A client requesting `docs.example.com` will be instructed to request `mydocs.myhost.com` instead.
When requesting mydocs.myhost.com the A record will be interrogated and it will return the IP address where the relevant data is served.

See more [here](https://www.netlify.com/blog/2016/01/12/this-weekends-ddos-attack-and-whats-in-a-cname/).

## TLS handshake

A TLS handshake takes place whenever a user navigates to a website over HTTPS and the browser first begins to query the website's origin server.

Goals:

- agree on a TLS (or SSL) version, based on the highest version supported between the 2 sides.
- agree on which cipher suites to use
- authenticate the identity of the server via its public key and the SSL cert

Steps:

**TODO**

## Anatomy of a browser web request

- the browser extracts the **domanin name** from the url
- the computer sends a DNS request to the DNS servers configured on the computer
- the DNS maps the domain to an IP address through its configured A records
