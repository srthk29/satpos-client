# GO GRPC EXAMPLE

## Reference
1. https://github.com/grpc/grpc-go/tree/master/examples/helloworld
2. https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go
3. https://github.com/grpc/grpc-go/blob/master/examples/route_guide/client/client.go
4. https://www.jetbrains.com/guide/go/tutorials/grpc_part_one/grpc_in_go/
5. https://grpc.io/docs/languages/go/generated-code/
6. https://github.com/jeanbza/jeanbza.github.io/blob/master/_posts/2020-10-09-stubbing-grpc.md
7. https://www.youtube.com/watch?v=Vbw8h0RCn2E



## Proto
```console
protoc --go_out=. --go-grpc_out=. proto/sat.proto
```
