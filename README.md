# mtls-golang
A simple golang project, comprising of mtls client-server implementation.

To run this project, follow the steps:
1. Generate the required certficate and key files from script.sh file
2. Run the server using the command: go run server.go
3. Navigate to client directory and run client using the command: go run client.go
4. You can also use this curl to create a request instead of using the client: curl --cacert ca.pem --cert client.pem --key client.key --location 'https://localhost:8443'
