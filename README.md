# go_grcp_protobuffers

Protocol Buffers are data format which stores data in a structured format. Data in the protocol buffer format can be serialized and deserialized by multiple languages.


# Compiling ProtoBuffers

```
Generator of ProtoBuffers for GO:
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

GRPC:
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```
The compilation goes:
```
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/student.proto    
```

# GRPC

## RPC
Remote Procedure Call. Way of establish a communication between a client and a server, but hiding the implementation.

## GRPC
Framework for:
- More efficient
- High Performance
- Uses HTTP2 for transport and Protobuffers as interchange method for data.
    - Http2 allows multiplexation when sending messages, using same connection to send more messages in a faster way. And allows serializing data and send them in requests.
    - Protobuffers let us serialize and deserialize for the data exchange can be faster. 
- Allows Streaming, sending data constantly.

- Methods:
    - Unary: It's like rest, client send request to server, and server responds. 
    - Server Streaming RPCs: Client send a request to the server and gets a stream to read a sequence of messages back. gRPC guarantees message ordering withing an individual RPC call
    - Client Streaming RPCs: Client writes a sequence of messages and sends them to the server, using a provided stream. Once the client has finished, it waits for the server to read them and return its response. gRPC guarantees message ordering.
    - Bidirectional Streaming RPCs: Both sides send a sequence of messages using a read-write stream. The two streams operate independently.

# Protobuffers Vs JSON

- **JSON** 
    - Easy to read for humans.
    - Json has to be serialized to be understandable. IE. Convert a Go Struct into JSON. The same for when receive information, we need to deserialize the data. That is slow.
    - Is a standard
    - More Flexible
- **ProtoBuffer**
    - We define a file, use a compiler and generate the file. This happens only once. It's agnostic. 
    - Serialize is faster than JSON, and deserialize too.
    - Difficult to read for humans.
  
If we need speed Protobuffers are better, but for more flexibility JSON.