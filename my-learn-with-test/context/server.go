package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	// Cancel(ctx context.Context)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)

		// ctx := r.Context()

		// data := make(chan string, 1)

		// go func() {
		// 	data, err <- store.Fetch(ctx)
		// }()

		// select {
		// case d := <-data:
		// 	fmt.Fprintf(w, d)
		// case <-ctx.Done():
		// 	store.Cancel(ctx)
		// }
	}
}
