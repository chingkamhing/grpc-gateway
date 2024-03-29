# gRPC Gateway Evaluation Project

This is a project to evaluate grpc-gateway.

## Use Cases

* authentication
    + in gateway proto file, config endpoints with google.api.http
    + config all endpoints with or without authentication
    + in gateway's out-bound interceptor, check if there is any "authorization" in the metadata, perform corresponding authentication if so
    + upon runtime.NewServeMux(), forward necessary request info to the metadata to fullfill the authentication needs
* cookies or session handling
    + how to integrate with session library (e.g. scs)?
    + ???
* open API settings
    + e.g. "x-nullable", "x-go-custom-tag", "required", etc.
    + ???

## Knowledge Base

* [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/)
* [Use gRPC interceptor for authorization with JWT](https://dev.to/techschoolguru/use-grpc-interceptor-for-authorization-with-jwt-1c5h)
* [grpc-auth-example](https://github.com/johanbrandhorst/grpc-auth-example)
* [Creating OpenSSL x509 certificates](https://adfinis.com/en/blog/openssl-x509-certificates/)
* [How to create an HTTPS certificate for localhost domains](https://gist.github.com/cecilemuller/9492b848eb8fe46d462abeb26656c4f8)
* [mTLS example](https://github.com/islishude/grpc-mtls-example)
* [How to secure gRPC connection with SSL/TLS in Go](https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph)