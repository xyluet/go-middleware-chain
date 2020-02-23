package chain_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	chain "github.com/xyluet/go-middleware-chain"
)

var middlewares []func(http.Handler) http.Handler

func init() {
	for i := 0; i < 100; i++ {
		middlewares = append(middlewares, func(http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
		})
	}
}

func Test_(t *testing.T) {
	// w, r := httptest.NewRecorder(), httptest.NewRequest("", "/", nil)
	// chain.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "asd")
	// })).ServeHTTP(w, r)
	// t.Fatal(w.Body.String())
}

func Benchmark(b *testing.B) {
	// mux := http.NewServeMux()
	// mux.HandleFunc()
	chained := chain.Middleware(middlewares...)
	chained2 := chain.Middleware(chained.Handler)
	w, r := httptest.NewRecorder(), httptest.NewRequest("", "/", nil)
	for i := 0; i < b.N; i++ {
		chained2.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}).ServeHTTP(w, r)
		chained2.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, r)
	}
}

func Benchmark2(b *testing.B) {
	chained := chi.Chain(middlewares...)
	chained2 := chi.Chain(chained.Handler)
	w, r := httptest.NewRecorder(), httptest.NewRequest("", "/", nil)
	for i := 0; i < b.N; i++ {
		chained2.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}).ServeHTTP(w, r)
		chained2.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, r)
	}
}
