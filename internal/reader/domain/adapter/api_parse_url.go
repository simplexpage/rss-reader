package adapter

import (
	rssparser "github.com/simplexpage/rss-parser"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
)

type APIParseUrl interface {
	GetApiData(urls []string) ([]model.Item, error)
}

type ParseUrlPackageAdapter struct {
	ParseUrlPackage *ParseUrlPackage
}

func (p *ParseUrlPackageAdapter) GetApiData(urls []string) ([]model.Item, error) {
	rssItems, err := p.ParseUrlPackage.ParseURLs(urls)
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
