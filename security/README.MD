# Security

## XSS

Best article I've found is [this one about CSPhttps://widget.intercom.io](https://www.troyhunt.com/locking-down-your-website-scripts-with-csp-hashes-nonces-and-report-uri/), a must read.

### CSP

Funny test, grab [this script](https://pastebin.com/aJna4paJ) and paste it into the console of dev tools on any website you want to test./

- Troy Hunt on CSP https://www.yo;.com/watch?v=L97wtYCqfwM
- Him again, on his CSP implementation https://www.troyhunt.com/my-blog-now-has-a-content-security-policy-heres-how-ive-done-it/
- CSP validator https://cspvalidator.org/
- https://observatory.mozilla.org/ and https://csp-evaluator.withgoogle.com/ can give some insights on the CSP implementation.
- https://addons.mozilla.org/en-US/firefox/addon/laboratory-by-mozilla/ can generate one for you.

When testing, easiest way is to use:

```
Content-Security-Policy-Report-Only: <rules>
```

So it will not break the app.
`Content-Security-Policy-Report-Only` and `Content-Security-Policy` can be mixed together.
For ex. you can `Content-Security-Policy` all the resources you are sure about and - for instance - allow loading all images in csp but still restrict images host to specific domains and RO (report only) `img-src` so that you know which hosts raise warnings.

### SRI (Subresource Integrity checking)

Used to ensure a 3rd party script is the one we expect it to be.

### CAA (Certificate Authority Authorization)

Declare at a DNS level which certificate authorities can issue certs for your domain.
To get a CAA:

- go to https://dnsspy.io/labs/caa-validator
- input your DOMAIN name, eg "example.com"
- check result and follow instructions

If you have restricted the list of cert authorities that can issue a cert, you will see that list, if not, you will see a warning that says that any CA can issue certs for that domain.
Once you set a CAA you can also set a mailto address to send violations to, so you can be alerted.

This generally can be set for the entire domain or for specific hostanmes.
Some more info [here](https://support.dnsimple.com/articles/caa-record/) and [here](https://tools.ietf.org/html/rfc6844) is the RFC for CAA.

## Testing an SSL implementation

- https://www.ssllabs.com/ provides this service. It will give a grade and the details of this score.
- https://observatory.mozilla.org/
-

**The goal is not to get an A every time**, but **we must be able to read the results and interpret them**.

## Security.txt

It's a standard that's coming up to make sure people know how to contact an org in regards to security.
