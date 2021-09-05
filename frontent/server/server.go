package server

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	http "github.com/valyala/fasthttp"
	"google.golang.org/grpc"
)

// Server defines the http server.
type Server struct {
	S *http.Server
}

// New inits new server.
func New() *Server {
	return &Server{&http.Server{
		ReadTimeout:  60 * time.Minute,
		WriteTimeout: 60 * time.Minute,
	}}
}

// RegisterStaticFiles adds root path to the static files to be served.
func (s *Server) RegisterStaticFiles(path string) *Server {
	s.S.Handler = newHandler().router(path)
	return s
}

// Start starts the server and listens on the specified port.
func (s *Server) Start(port string) error {
	return s.S.ListenAndServe(fmt.Sprintf(":%s", port))
}

type handler struct {
	*grpcweb.WrappedGrpcServer
}

func newHandler() *handler {
	return &handler{grpcweb.WrapServer(grpc.NewServer())}
}

func isGrpcReq(ctx *http.RequestCtx) bool {
	if ctx.Request.Header.IsPost() && string(ctx.Request.Header.ContentType()) == "application/grpc-web" {
		return true
	}
	return false
}

// Router wraps static files serving handler.
func (h *handler) router(root string) http.RequestHandler {
	logger := log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Lmsgprefix|log.LUTC|log.Llongfile)
	return func(ctx *http.RequestCtx) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("unknown error")
				}
				logger.Println(err)
			}
		}()
		if isGrpcReq(ctx) {
			// h.ServeHTTP()
			logger.Println("not implemented")
		}
		http.FSHandler(root, 0)
	}
}
