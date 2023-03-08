package adapter

import (
	"context"
	"github.com/golang/mock/gomock"
	rssparser "github.com/simplexpage/rss-parser"
	"github.com/simplexpage/rss-reader/internal/reader/domain/adapter/mocks"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetApiDataSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parseUrlPackageMock := mocks.NewMockParseUrlPackageService(ctrl)
	parseUrlService, err := NewParseUrlPackageService(parseUrlPackageMock)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	mockInputUrls := []string{"https://tsn.ua/rss/full.rss"}
	mockResponse := []rssparser.RssItem{
		{
			Title:       "title",
			Source:      "source",
			SourceURL:   "sourceUrl",
			Link:        "link",
			PublishDate: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			Description: "description",
		},
	}
	expected := []model.Item{
		{
			Title:       "title",
			Source:      "source",
			SourceUrl:   "sourceUrl",
			Link:        "link",
			PublishDate: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			Description: "description",
		},
	}
	parseUrlPackageMock.EXPECT().ParseURLs(ctx, mockInputUrls).Return(mockResponse, nil).Times(1)
	items, err := parseUrlService.GetApiData(ctx, mockInputUrls)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	assert.Equal(t, expected, items)
}
