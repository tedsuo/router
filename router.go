package router

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/bmizerany/pat"
)

// Handlers map Handler keys to http.Handler objects.  The Handler key is used to
// then match the Handler to the appropriate Route in the Router.
type Handlers map[string]http.Handler

// NewRouter combines a set of Routes with their corresponding Handlers to
// produce a http request multiplexer (AKA a "router").
func NewRouter(routes Routes, handlers Handlers) (http.Handler, error) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\n\033[0;35m%s\033[0m%s:%d:%s\n", "WARNING:", file, line, " package router is deprecated, please use github.com/tedsuo/rata instead")

	p := pat.New()
	for _, route := range routes {
		handler, ok := handlers[route.Handler]
		if !ok {
			return nil, fmt.Errorf("missing handler %s", route.Handler)
		}
		switch strings.ToUpper(route.Method) {
		case "GET":
			p.Get(route.Path, handler)
		case "POST":
			p.Post(route.Path, handler)
		case "PUT":
			p.Put(route.Path, handler)
		case "DELETE":
			p.Del(route.Path, handler)
		default:
			return nil, fmt.Errorf("invalid verb: %s", route.Method)
		}
	}
	return p, nil
}
