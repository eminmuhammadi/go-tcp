mkdir -p certs

rm -f certs/*.pem

openssl req -config cert.conf -new -nodes -x509 -newkey rsa:4096 -sha512 -keyout certs/client.key -out certs/client.pem -days 365
openssl req -config cert.conf -new -nodes -x509 -newkey rsa:4096 -sha512 -keyout certs/server.key -out certs/server.pem -days 365