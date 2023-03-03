package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/simplexpage/rss-reader/internal/reader/delivery/reqresp"
	"github.com/simplexpage/rss-reader/internal/reader/domain/service"
)

type Set struct {
	ParseUrlsEndpoint endpoint.Endpoint
}

func NewServerEndpoints(s service.Service, logger log.Logger) Set {
	var parseUrlsEndpoint endpoint.Endpoint
	{
		parseUrlsEndpoint = MakeParseUrlsEndpoint(s)
		parseUrlsEndpoint = LoggingMiddleware(log.With(logger, "method", "createPartner"))(parseUrlsEndpoint)
	}
	return Set{
		ParseUrlsEndpoint: parseUrlsEndpoint,
	}
}

func MakeParseUrlsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(reqresp.ParseUrlsRequest)
		items, err, errV := s.ParseUrls(ctx, req.ParseUrlsForm)
		return reqresp.ParseUrlsResponse{Items: AdapterItems(items), Err: err, ValidationErr: errV}, nil
	}
}
