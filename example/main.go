package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-puzzles/cores"
	httpPuzzle "github.com/go-puzzles/http-puzzle"
	"github.com/go-puzzles/plog"
	"github.com/go-puzzles/plog/level"
	"github.com/gorilla/mux"
	"github.com/superwhys/venkit/lg/v2"
	"github.com/superwhys/venkit/vrouter/v2"
)

func TestWorker(ctx context.Context) error {
	t := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-t.C:
		case <-ctx.Done():
			return ctx.Err()
		}

		fmt.Println("worker running")
	}
}

func main() {
	router := mux.NewRouter()
	plog.Enable(level.LevelDebug)

	router.Methods(http.MethodGet).Path("/hello").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vrouter.WriteJSON(w, 200, map[string]string{
			"data": "hello world",
		})
	})

	srv := cores.NewPuzzleCore(
		cores.WithWorker(TestWorker),
		httpPuzzle.WithCoreHttpPuzzle("/api/v2", router),
		httpPuzzle.WithCoreHttpPuzzle("/api", router),
	)

	lg.PanicError(cores.Start(srv, 0))
}
