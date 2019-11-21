# Networking

## DNS records

DNS (Domain name system) is the internet-wide phonebook that tells browsers or any other client how to resolve a specific domain name to an IP address of a server the can open a connection to.

Websites have 3 kinds of records today:

- A Record: returns an IPv4 address of a server
- AAAA Record: returns the IPv6 address of a server
- CNAME Record: returns another domain name the browser should look up instead.

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
