# go-puzzles/cores

## 特性
简单易用、足够轻量，避免过多的外部依赖

目前实现了以下特性：
- 任务管理
- 守护进程
- 优雅终止
- 服务发现
- 服务注册

支持各种外部扩展:
- httpServer
- grpcServer
- gprcuiHandler

## 快速上手

### 安装
```shell
go get github.com/go-puzzles/cores
```

## Http服务 
```go
package main

import (
	"net/http"
	
	"github.com/go-puzzles/cores"
	httppuzzle "github.com/go-puzzles/cores/puzzles/http-puzzle"
	"github.com/go-puzzles/plog"
	"github.com/gorilla/mux"
)

func main() {
	pflags.Parse()
	router := mux.NewRouter()
	
	router.Path("/hello").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	
	core := cores.NewPuzzleCore(
		httppuzzle.WithCoreHttpPuzzle("/api", router),
	)
	
    if err := cores.Start(core, port()); err != nil {
        panic(err)
    }
}
```

## Grpc服务
```go
package main

import (
	"github.com/go-puzzles/cores"
	grpcpuzzle "github.com/go-puzzles/cores/puzzles/grpc-puzzle"
	grpcuipuzzle "github.com/go-puzzles/cores/puzzles/grpcui-puzzle"
	"github.com/go-puzzles/example/cores-with-grpc/examplepb"
	srv "github.com/go-puzzles/example/cores-with-grpc/service"
	"github.com/go-puzzles/example/cores-with-grpc/testpb"
	"google.golang.org/grpc"
)

func main() {
	example := srv.NewExampleService()
	test := srv.NewTestService()

	srv := cores.NewPuzzleCore(	
		grpcpuzzle.WithCoreGrpcPuzzle(func(srv *grpc.Server) {
			examplepb.RegisterExampleHelloServiceServer(srv, example)
			testpb.RegisterExampleHelloServiceServer(srv, test)
		}),
	)

	if err := cores.Start(srv, 0); err != nil {
		panic(err)
	}
}
```

## 开启GRPCUI
```go
srv := cores.NewPuzzleCore(	
    grpcuipuzzle.WithCoreGrpcUI(),
	grpcpuzzle.WithCoreGrpcPuzzle(func(srv *grpc.Server) {
		examplepb.RegisterExampleHelloServiceServer(srv, example)
		testpb.RegisterExampleHelloServiceServer(srv, test)
	}),
)
```

## Consul服务注册
```go
package main

import (
	"github.com/go-puzzles/cores"
	consulpuzzle "github.com/go-puzzles/cores/puzzles/consul-puzzle"
)

func main() {
	pflags.Parse(
		pflags.WithConsulEnable(),
	)

	core := cores.NewPuzzleCore(
		cores.WithService(pflags.GetServiceName()),
		consulpuzzle.WithConsulRegsiter(),
	)

	cores.Start(core, port())
}
```

## Consul服务发现
```go
package main

import (
	"github.com/go-puzzles/cores/discover"
)


func main() {
    // ....
    discover.GetServiceFinder().GetAddress("serviceName")
    discover.GetServiceFinder().GetAddressWithTag("serviceName", "v.0.0")
    // ....
}
```
