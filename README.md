# grpc-demo
This is a repository which gives an overview of basic implementation for client and server using [gRPC](https://grpc.io/)
### Folder Structure
`/calculatorpb`: Protocol buffers package which contains the gRPC message and service definitions.
`/server`: gRPC server written in GoLang which performs basic mathematical operations.
`/goclient`: gRPC client written in Golang which runs a http server and gets a request from client to forward to another server by sending a request and receiving a result of the mathematical operation as a response.
`codegen.sh` A script to generate the protocol buffers go code from the .proto file. (make sure you have thord party folder in vendors directory or you can change the path accordingly)
### Usage
Clone the repo inside your `$GOPATH`.
From the root of the project, install the required dependencies by running:
```
go get
```
#### Run your server:
```
go run server/main.go
```
 While running the above command, your server will start listening at port `:3030` for all the remote procedure calls. 
#### Run your client:
```
go run client/client.go
```
This will run a gin server which accepts an http request on port `:9002` and serves a response in http format by fetching it using gRPC from another sever. 

#### Help / Support

If you run into any issues, please email us at [ashhadsheikh@hotmail.com](mailto:ashhadsheikh@hotmail.com)
For bug reports, please [open an issue on GitHub](https://github.com/ashhadsheikh/grpc-demo/issues/new).

## Contributing
1. Fork it
2. Create your feature branch (```git checkout -b my-new-feature```).
3. Commit your changes (```git commit -am 'Added some feature'```)
4. Push to the branch (```git push origin my-new-feature```)
5. Create new Pull Request