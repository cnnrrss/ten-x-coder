# CORS

[See.](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)

**Cross-Origin Resource Sharing (CORS)** is a mechanism that uses additional HTTP headers to tell a browser to let a web application running at one origin (domain) have permission to access selected resources from a server at a different origin. A web application executes a cross-origin HTTP request when it requests a resource that has a different origin (domain, protocol, and port) than its own origin.

## Simple Requests:
The only allowed methods are:
`GET` `HEAD` `POST` .... 
`DELETE` not allowed

Apart from the headers set automatically by the user agent (for example, Connection, User-Agent, or any of the other headers with names defined in the Fetch spec as a “forbidden header name”), the only headers which are allowed to be manually set are those which the Fetch spec defines as being a “CORS-safelisted request-header”, which are:

```
Accept
Accept-Language
Content-Language
Content-Type (but note the additional requirements below)
DPR
Downlink
Save-Data
Viewport-Width
Width
```

The only allowed values for the Content-Type header are:
```
application/x-www-form-urlencoded
multipart/form-data
text/plain
```