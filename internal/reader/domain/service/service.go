package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/gookit/validate"
	"github.com/simplexpage/rss-reader/internal/reader/delivery/reqresp/form"
	"github.com/simplexpage/rss-reader/internal/reader/domain/adapter"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
	"github.com/simplexpage/rss-reader/pkg/errors"
)

type Service interface {
	ParseUrls(ctx context.Context, parseUrlsForm form.ParseUrlsForm) ([]model.Item, error, validate.Errors)
}

func New(logger log.Logger) Service {
	var service Service
	{
		service = NewReaderService(logger)
		service = LoggingMiddleware(logger)(service)
	}
	return service
}

func NewReaderService(logger log.Logger) Service {
	return &readerService{
		logger: logger,
	}
}

type readerService struct {
	logger log.Logger
}

func (r *readerService) ParseUrls(ctx context.Context, parseUrlsForm form.ParseUrlsForm) ([]model.Item, error, validate.Errors) {
	validationReq := validate.Struct(parseUrlsForm)
	if !validationReq.Validate() {
		return nil, errors.ErrDataValidation, validationReq.Errors
	}

	parseUrlPackageAdapter := adapter.ParseUrlPackageAdapter{ParseUrlPackage: &adapter.ParseUrlPackage{}}
	items, err := r.getApiDataUrls(&parseUrlPackageAdapter, parseUrlsForm.Urls)
	if err != nil {
		return nil, err, nil
	}
	return items, nil, nil
}

func (r *readerService) getApiDataUrls(api adapter.APIParseUrl, urls []string) ([]model.Item, error) {
	return api.GetApiData(urls)
}
