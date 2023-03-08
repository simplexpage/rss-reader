package adapter

import (
	"context"
	"errors"
	rssparser "github.com/simplexpage/rss-parser"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
	"time"
)

//go:generate mockgen -source=api_parse_url.go -destination=mocks/mock_api_parse_url.go -package=mocks
type ParseUrlService interface {
	GetApiData(ctx context.Context, urls []string) ([]model.Item, error)
}

func NewParseUrlPackageService(parseService interface{}) (ParseUrlService, error) {
	switch parseService.(type) {
	case ParseUrlPackageService:
		return &parseUrlServiceAdapter{
			parseService: parseService.(ParseUrlPackageService),
		}, nil
	}
	return nil, errors.New("parseService is unknown type")
}

type parseUrlServiceAdapter struct {
	parseService ParseUrlPackageService
}

func (p *parseUrlServiceAdapter) GetApiData(ctx context.Context, urls []string) ([]model.Item, error) {
	rssItems, err := p.parseService.ParseURLs(ctx, urls)
	if err != nil {
		return nil, err
	}
	items := make([]model.Item, 0)
	for _, rssItem := range rssItems {
		items = append(items, model.Item{
			Title:       rssItem.Title,
			Link:        rssItem.Link,
			Description: rssItem.Description,
			PublishDate: rssItem.PublishDate,
			Source:      rssItem.Source,
			SourceUrl:   rssItem.SourceURL,
		})
	}
	return items, nil
}

type ParseUrlPackageService interface {
	ParseURLs(ctx context.Context, urls []string) ([]rssparser.RssItem, error)
}

type ParseTimeout struct {
	Duration time.Duration
}

type parseUrlPackage struct {
	timeout ParseTimeout
}

func NewParseUrlPackage(timeout ParseTimeout) ParseUrlPackageService {
	return &parseUrlPackage{
		timeout: timeout,
	}
}

func (p *parseUrlPackage) ParseURLs(ctx context.Context, urls []string) ([]rssparser.RssItem, error) {
	var ls []rssparser.RssItem
	var err error
	if p.timeout.Duration > 0 {
		ctxTime, cancel := context.WithTimeout(ctx, p.timeout.Duration)
		defer cancel()
		ls, err = rssparser.ParseURLs(ctxTime, urls)
	} else {
		ls, err = rssparser.ParseURLs(ctx, urls)
	}

	if err != nil {
		return nil, err
	}
	return ls, nil
}
