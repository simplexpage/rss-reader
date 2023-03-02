package transport

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httpKitTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/simplexpage/rss-reader/internal/reader/endpoint"
	httpUtil "github.com/simplexpage/rss-reader/pkg/transport/http"
	httpMiddleware "github.com/simplexpage/rss-reader/pkg/transport/http/middleware"
	"net/http"
)

func NewHTTPHandler(endpoints endpoint.Set, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	// CORS
	r.Use(mux.CORSMethodMiddleware(r))
	corsMiddleware := httpMiddleware.NewCorsMiddleware()
	r.Use(corsMiddleware.Middleware)

	options := []httpKitTransport.ServerOption{
		httpKitTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httpKitTransport.ServerErrorEncoder(httpUtil.EncodeError),
	}

	// Routes Not Auth
	// POST    /reader/parse
	// GET     /health

	r.Methods(http.MethodPost).Path("/reader/parse").Handler(httpKitTransport.NewServer(
		endpoints.ParseUrlsEndpoint,
		decodeHttpParseRequest,
		httpUtil.EncodeHttpResponse,
		options...,
	))

	//Health Check
	r.HandleFunc("/health", httpUtil.HealthCheckHandler)

	return r
}

func decodeHttpParseRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return
}
