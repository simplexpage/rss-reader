package adapter

import (
	"context"
	rssparser "github.com/simplexpage/rss-parser"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
)

//go:generate mockgen -source=api_parse_url.go -destination=mocks/mock_api_parse_url.go -package=mocks
type APIParseUrlService interface {
	GetApiData(ctx context.Context, urls []string) ([]model.Item, error)
}

type parseUrlPackageAdapter struct {
	parseUrlPackage *ParseUrlPackage
}

func NewParseUrlPackageAdapter() APIParseUrlService {
	return &parseUrlPackageAdapter{
		parseUrlPackage: &ParseUrlPackage{},
	}
}

func (p *parseUrlPackageAdapter) GetApiData(ctx context.Context, urls []string) ([]model.Item, error) {
	rssItems, err := p.parseUrlPackage.ParseURLs(urls)
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

type ParseUrlPackage struct{}

func (p *ParseUrlPackage) ParseURLs(urls []string) ([]rssparser.RssItem, error) {
	ls, err := rssparser.ParseURLs(urls)
	if err != nil {
		return nil, err
	}
	return ls, nil
}
