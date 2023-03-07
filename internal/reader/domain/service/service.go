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

//go:generate mockgen -source=service.go -destination=mocks/mock_service.go -package=mocks
type Service interface {
	ParseUrls(ctx context.Context, parseUrlsForm form.ParseUrlsForm) ([]model.Item, error, validate.Errors)
}

func New(logger log.Logger, parseUrlService adapter.APIParseUrlService) Service {
	var service Service
	{
		service = NewReaderService(logger, parseUrlService)
		service = LoggingMiddleware(logger)(service)
	}
	return service
}

func NewReaderService(logger log.Logger, parseUrlService adapter.APIParseUrlService) Service {
	return &readerService{
		logger:          logger,
		parseUrlService: parseUrlService,
	}
}

type readerService struct {
	logger          log.Logger
	parseUrlService adapter.APIParseUrlService
}

func (r *readerService) ParseUrls(ctx context.Context, parseUrlsForm form.ParseUrlsForm) ([]model.Item, error, validate.Errors) {
	validationReq := validate.Struct(parseUrlsForm)
	if !validationReq.Validate() {
		return nil, errors.ErrDataValidation, validationReq.Errors
	}

	items, err := r.parseUrlService.GetApiData(ctx, parseUrlsForm.Urls)
	if err != nil {
		return nil, err, nil
	}
	return items, nil, nil
}
