package endpoint

import (
	"github.com/simplexpage/rss-reader/internal/reader/delivery/reqresp"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
)

func AdapterItems(items []model.Item) []reqresp.Item {
	itemsAdapter := make([]reqresp.Item, len(items))
	for _, item := range items {
		itemsAdapter = append(itemsAdapter, reqresp.Item{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Source:      item.Source,
			SourceUrl:   item.SourceUrl,
			PublishDate: item.PublishDate.Format("2006-01-02 15:04:05"),
		})
	}
	return itemsAdapter
}
