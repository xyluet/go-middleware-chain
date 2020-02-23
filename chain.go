package chain

import "net/http"

type Chain struct {
	middlewares []func(http.Handler) http.Handler
}

func (c *Chain) Handler(h http.Handler) http.Handler {
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		h = c.middlewares[i](h)
	}
	return h
}

func (c *Chain) HandlerFunc(h http.HandlerFunc) http.Handler {
	return c.Handler(h)
}

func Middleware(middlewares ...func(http.Handler) http.Handler) *Chain {
	return &Chain{
		middlewares: middlewares,
	}
}

// chained := chain.Middleware(mid1, mid2)
// chained2 := chain.Middleware(chained.Handler, mid3)
// chained3 := chained2.Handle
