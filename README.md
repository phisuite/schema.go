# Phi Suite Schema.go

| **Homepage** | [https://phisuite.com][0]        |
| ------------ | -------------------------------- | 
| **GitHub**   | [https://github.com/phisuite][1] |

## Overview

This project contains the Go module to create the **Schema API** server & client.  
For more details, see [Phi Suite Schema][2].

## Installation

```bash
go get github.com/phisuite/schema.go
```

## Creating the server

```go
package main
import "github.com/phisuite/schema.go"

type eventAPIServer struct { ... }

func (s *eventAPIServer) Get(ctx context.Context, req *schema.Options) (*schema.Event, error) {
    ... 
}

lis, err := net.Listen("tcp", ":50051")
grpcServer := grpc.NewServer()
schema.RegisterEventAPIServer(grpcServer, &eventAPIServer{})
grpcServer.Serve(lis)
```
For more details, see [gRPC Basics - Go: Creating the server][10].

## Creating the client

```go
package main
import "github.com/phisuite/schema.go"

conn, err := grpc.Dial(*serverAddr)
defer conn.Close()

client := schema.NewEventAPIClient(conn)
event := client.Get(context.Background(), &schema.Options{ ... })
```
For more details, see [gRPC Basics - Go: Creating the client][11].

[0]: https://phisuite.com
[1]: https://github.com/phisuite
[2]: https://github.com/phisuite/schema
[10]: https://www.grpc.io/docs/tutorials/basic/go/#server
[11]: https://www.grpc.io/docs/tutorials/basic/go/#client
