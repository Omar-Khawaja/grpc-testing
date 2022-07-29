# gRPC Testing

To start the gRPC server, run the following command from the root of the repository:

```shell
go run cmd/server/server.go 
```

To execute the gRPC client, run the following command from the root of the repository:

```shell
go run cmd/client/client.go
```

## Alternative Ways to Test gRPC Server

Running the client is not necessary to test the gRPC service. The gRPC server implements [server reflection](https://pkg.go.dev/google.golang.org/grpc/reflection) (you can read more about it [here](https://github.com/grpc/grpc/blob/master/doc/server-reflection.md)) so you may use tools like [grpcurl](https://github.com/fullstorydev/grpcurl) to test the service.

Example:

```shell
grpcurl -plaintext -d '{"name":"Omar"}' localhost:8080 hello.Greeter/SayHello
```

The previous command should give you the following output from the gRPC server:

```shell
{
  "message": "Hello there Omar"
}
```
