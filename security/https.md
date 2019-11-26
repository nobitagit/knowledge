# HTTPS

## MitM (Man in the middle)

Typical insertion points include:

- the connection between client and router (1)
- the router itself
- the connection between router and ISP (2)
- the ISP itself
- the connection between ISP and server, eg governments can intercept traffic (3)

```
  ----------   1    ----------   2   ----------   3   ----------
  | client | ---->  | router | ----> |   ISP  | ----> | server |
  ----------        ----------       ----------       ----------
```

Basically every node and every transition of data can be a potential insertion point for a MitM attack.

## HTTPS value proposition

- **Confidentiality**: we want to be sure that info that we send from client to server remains secret (we read an email, we send a pw etc.).
- **Integrity**: we want to be sure that info is not tampered with. so responses arrive as they are sent by the server, and for example malware is not injected into the page.

For instance:

> Malicious code injected into Tunisian versions of Facebook, Gmail, and Yahoo! stole login credentials of users critical of the North African nation's authoritarian government, according to security experts and news reports.
>
> [...]
>
> After more than ten days of intensive investigation and study, Facebook's security team realized something very, very bad was going on,” The Atlantic article stated. “The country's internet service providers were running a malicious piece of code that was recording users' login information when they went to sites like Facebook.
> -- [source](https://www.theregister.co.uk/2011/01/25/tunisia_facebook_password_slurping/)

- **Authenticity** so we are actually confident we are really talking to the server we think we're talking to. HTTPS avoids attacks such as DNS hijacking (by poisoning the DNS, forcing them to resolve to different IP addresses that correspond to rogue sites).

## CA (Certificate authorities)

When the browser receives the cert from a server, it checks whether the CA is trusted, by checking against a list of CAs.

Every browser follows a different list:

- Firefox: https://www.mozilla.org/en-US/about/governance/policies/security-group/certs/included/
- Chrome: it uses the certificates [included with the OS][1].
- Opera: it install the most uses CA within installing the application, you can find the rest in the online root repository: https://certs.opera.com/
- iOs: https://support.apple.com/kb/ht5012

([source](https://security.stackexchange.com/a/49008/155108))

> The Mozilla CA Certificate Program's list of included root certificates is stored in a file called [certdata.txt](https://hg.mozilla.org/releases/mozilla-beta/file/tip/security/nss/lib/ckfw/builtins/certdata.txt) in the Mozilla source code management system.
>
> -- [source](https://wiki.mozilla.org/CA/Included_Certificates)

> Google Chrome attempts to use the root certificate store of the underlying operating system to determine whether an SSL certificate presented by a site is indeed trustworthy, with a few exceptions.
>
> -- [source](http://www.chromium.org/Home/chromium-security/root-ca-policy)

New CAs are added and removed with time.

A staple of CAs is that **they can only issue certs to owners of the domain they are certifying**.

## SSL

- Secure Socket Layer
- built by Netscape in the 90s

## TLS

- Transport Layer Security
- built in 1999
- a **direct** upgrade of SSL 3.0

## SSL and TLS

TSL and SSL terms are used interchangeably, since TLS is a direct upgrade of SSL, but basically SSL reached EOL in 2014. When people say SSL they pretty much every time mean TLS.
**TLS is the current standard for implementing HTTPS**.

A server and a browser communicating might support different versions of TLS.
For instance:

- server A supports TLS v.1.3 (and lower)
- client B supports TLS v.1.2 (and lower)

The standard is to fallback to the **highest commonly supported versions** (in this case 1.2).
