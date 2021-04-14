# AWS

## Cloudfront

CF is a CDN:

- If the content is already in the edge location with the lowest latency, CloudFront delivers it immediately.

- If the content is not in that edge location, CloudFront retrieves it from an origin that you've defined—such as an Amazon S3 bucket, a MediaPackage channel, or an HTTP server (for example, a web server) that you have identified as the source for the definitive version of your content.

How you configure CloudFront to deliver your content:

1. You specify origin servers, like an Amazon S3 bucket or your own HTTP server, from which CloudFront gets your files which will then be distributed from CloudFront edge locations all over the world.

2. You upload the files (objects) to the origin server

3. You create a CloudFront distribution, which tells CloudFront which origin servers to get your files from when users request the files through your web site or application.

4. CloudFront assigns a domain name to your new distribution that you can see in the CloudFront console,

5. CloudFront sends your distribution's configuration (but not your content) to all of its edge locations or points of presence (POPs)
