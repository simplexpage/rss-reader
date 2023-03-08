package service

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/simplexpage/rss-reader/internal/reader/delivery/reqresp/form"
	"github.com/simplexpage/rss-reader/internal/reader/domain/adapter/mocks"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseUrlsSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parseUrlService := mocks.NewMockParseUrlService(ctrl)
	service := New(log.NewNopLogger(), parseUrlService)

	ctx := context.Background()

	parseUrlsForm := form.ParseUrlsForm{Urls: []string{"https://tsn.ua/rss/full.rss"}}
	mockResp := []model.Item{
		{
			Title:       "title",
			Source:      "source",
			SourceUrl:   "sourceUrl",
			Link:        "link",
			PublishDate: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			Description: "description",
		},
	}
	parseUrlService.EXPECT().GetApiData(ctx, parseUrlsForm.Urls).Return(mockResp, nil).Times(1)
	items, err, _ := service.ParseUrls(ctx, parseUrlsForm)
	require.NoError(t, err)
	require.NotEmpty(t, items)
}

func TestParseUrlsErrorValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parseUrlService := mocks.NewMockParseUrlService(ctrl)
	service := New(log.NewNopLogger(), parseUrlService)

	ctx := context.Background()

	parseUrlsForm := form.ParseUrlsForm{
		Urls: []string{},
	}

	items, err, errV := service.ParseUrls(ctx, parseUrlsForm)
	require.Error(t, err, errors.New("data validation failed"))
	require.Error(t, errV, errors.New("urls is required to not be empty"))
	require.Empty(t, items)
}
