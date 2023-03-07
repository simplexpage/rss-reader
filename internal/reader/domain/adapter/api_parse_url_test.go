package adapter

import (
	"context"
	"github.com/golang/mock/gomock"
	mockAdapter "github.com/simplexpage/rss-reader/internal/reader/domain/adapter/mocks"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
	"testing"
)

func TestParseItems(t *testing.T) {
	type mockBehavior func(s *mockAdapter.MockAPIParseUrlService, ctx context.Context, urls []string)

	testTable := []struct {
		name         string
		parseUrls    []string
		mockBehavior mockBehavior
	}{
		{
			name:      "success",
			parseUrls: []string{"https://tsn.ua/rss/full.rss"},
			mockBehavior: func(s *mockAdapter.MockAPIParseUrlService, ctx context.Context, urls []string) {
				s.EXPECT().GetApiData(ctx, urls).Return([]model.Item{}, nil).AnyTimes()
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			parseUrlPackageAdapterMock := mockAdapter.NewMockAPIParseUrlService(ctrl)
			testCase.mockBehavior(parseUrlPackageAdapterMock, context.Background(), testCase.parseUrls)
		})
	}
}
