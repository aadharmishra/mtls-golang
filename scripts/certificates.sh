openssl genrsa -out ca.key 2048

openssl req -x509 -new -key ca.key -out ca.crt -subj "/C=US/ST=California/L=SanFrancisco/O=MyCompany/CN=MyRootCA" -days 365

openssl genrsa -out server.key 2048

openssl req -new -key server.key -out server.csr -subj "/C=US/ST=California/L=SanFrancisco/O=MyCompany/CN=localhost"

openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365

openssl genrsa -out client.key 2048

openssl req -new -key client.key -out client.csr -subj "/C=US/ST=California/L=SanFrancisco/O=MyCompany/CN=YourClientName"

openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365