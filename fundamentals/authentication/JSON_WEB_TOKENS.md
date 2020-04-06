# JWT

[Token Based Authentication](https://auth0.com/learn/token-based-authentication-made-easy/) with JWTs.

A token is a piece of data with no meaning or use on its own. Combined with the correct tokenization system, becomes a vital player in securing your app.

Token based authentication works by ensuring that each request to a server is accompanied by a signed token which the server verifies for authenticity and only then responds to the request.

###  JSON Web Token (JWT)

JWT is an open standard (RFC 7519) that defines a compact and self-contained method for securely transmitting information between parties encoded as a JSON object.

JWT has gained mass popularity due to its compact size which allows tokens to be easily transmitted via query strings, header attributes and within the body of a POST request.

## Why use Tokens?

Tokens have many benefits compared to other methods such as cookies:

- **Stateless**: Tokens are self-contained and contain all the information it needs for authentication. Frees your server from having to store session state.

- **Tokens generated anywhere**: Token generation is decoupled from token verification allowing you the option to handle the signing of tokens on a separate server or even through a different company such us Auth0.

- **Fine-grained access control**: Within the token payload you can easily specify user roles and permissions as well as resources that the user can access.


## Anayomy of a JSON Web Token

**3 Parts**:
- Header
- Payload
- Signature


## Use cases for token based authentication

- **PaaS Applications**: Exposing RESTful APIs that will be consumed by a variety of frameworks and clients
- ***Mobile Apps**: implementing native or hybrin mobile app that interact with your services.
- **Single Page Application** (SPAs): building modern applications with frameworks such as React.


#### Terms and Vocab
- JWT: Json Web Token
- JWK: JSON Web Key [link](https://tools.ietf.org/html/rfc7517)
- OTP: One-time password
- SSO: Single Sign On
- AD: Active Directory
- MSISDN: Mobile Station International Subscriber Directory Number