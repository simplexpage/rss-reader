package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/validate"
	"github.com/simplexpage/rss-reader/internal/reader/delivery/reqresp/form"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (l loggingMiddleware) ParseUrls(ctx context.Context, parseUrlsForm form.ParseUrlsForm) (items []model.Item, err error, errV validate.Errors) {
	defer func() {
		level.Info(l.logger).Log("method", "ParseUrls", "urls", arrutil.AnyToString(parseUrlsForm.Urls), "err", err, "validation", errV)

	}()
	return l.next.ParseUrls(ctx, parseUrlsForm)
}
