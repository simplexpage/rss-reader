package service

import "github.com/go-kit/kit/log"

type Service interface {
}

func New(logger log.Logger) Service {
	var service Service
	{
		service = NewReaderService(logger)
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
