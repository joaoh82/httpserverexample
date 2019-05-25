# httpserverexample
Simple example of a HTTP Server and Client with HTTP/1 and HTTP/2 with Go Language

Generating New SSL Certificates: https://www.digitalocean.com/community/tutorials/openssl-essentials-working-with-ssl-certificates-private-keys-and-csrs

```
openssl req \
       -newkey rsa:2048 -nodes -keyout server.key \
       -x509 -days 365 -out server.crt
       