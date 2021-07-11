# mtlsexample
An example of using mTLS with self signed certificates


## Rational

I wanted to create a set of microservices which must share sensitive data. The connections must be encrypted and client and server need to authenticate themselves. The microservices must spin up and spin down on demand without disturbing any other running services. This setup eleminates using a proxy like Nginx to handle TLS as changing the nginx config requires restarting Nginx and working with the Nginx config files adds a layer of complexity.

An application could, in theory, use an already exisiting the certificate on the server such as one obtained through letsencrypt. However, the certs are owned by root so this would require changing the ownership of the certs or running the application as root. These are not good options from a security standpoint.

The solution is to use mTLS with self-signed certificates. The connection is encrypted and secure. The validating CA is only accessable to the application and both the client and server validate each other.

Furthermore, this implementation takes care of encryptian and authentican at the same time making it simpler than handeling them seperatly by using one method for setting up the TLS such as an Nginx proxy and another method like OAuth for authentication. 


## Creating the cert

As of go version 1.14, certificats must have a SAN and be signed by a CA. The easist tool for creating these kinds of certs is [certstrap](https://github.com/square/certstrap/releases).

Instructions for creating the required certificate using certstrap can be found on Rich Youngkin's [blog](https://youngkin.github.io/post/gohttpsclientserver/).

In a production environment, you would want to keep your certs somewhere safely hidden away. They are only in this repo for demonstration purposes.


## Benmarking

It took 2961ms to send and recieve 10,000 requests. It apears there is little overhead to this implementation. 


## Further reading

- [Create Secure Clients and Servers in Golang Using HTTPS](https://youngkin.github.io/post/gohttpsclientserver/)
- [A step by step guide to mTLS in Go](https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go)