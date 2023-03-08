// Code generated by MockGen. DO NOT EDIT.
// Source: api_parse_url.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	rss_parser "github.com/simplexpage/rss-parser"
	model "github.com/simplexpage/rss-reader/internal/reader/domain/model"
)

// MockParseUrlService is a mock of ParseUrlService interface.
type MockParseUrlService struct {
	ctrl     *gomock.Controller
	recorder *MockParseUrlServiceMockRecorder
}

// MockParseUrlServiceMockRecorder is the mock recorder for MockParseUrlService.
type MockParseUrlServiceMockRecorder struct {
	mock *MockParseUrlService
}

// NewMockParseUrlService creates a new mock instance.
func NewMockParseUrlService(ctrl *gomock.Controller) *MockParseUrlService {
	mock := &MockParseUrlService{ctrl: ctrl}
	mock.recorder = &MockParseUrlServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParseUrlService) EXPECT() *MockParseUrlServiceMockRecorder {
	return m.recorder
}

// GetApiData mocks base method.
func (m *MockParseUrlService) GetApiData(ctx context.Context, urls []string) ([]model.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiData", ctx, urls)
	ret0, _ := ret[0].([]model.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApiData indicates an expected call of GetApiData.
func (mr *MockParseUrlServiceMockRecorder) GetApiData(ctx, urls interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiData", reflect.TypeOf((*MockParseUrlService)(nil).GetApiData), ctx, urls)
}

// MockParseUrlPackageService is a mock of ParseUrlPackageService interface.
type MockParseUrlPackageService struct {
	ctrl     *gomock.Controller
	recorder *MockParseUrlPackageServiceMockRecorder
}

// MockParseUrlPackageServiceMockRecorder is the mock recorder for MockParseUrlPackageService.
type MockParseUrlPackageServiceMockRecorder struct {
	mock *MockParseUrlPackageService
}

// NewMockParseUrlPackageService creates a new mock instance.
func NewMockParseUrlPackageService(ctrl *gomock.Controller) *MockParseUrlPackageService {
	mock := &MockParseUrlPackageService{ctrl: ctrl}
	mock.recorder = &MockParseUrlPackageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParseUrlPackageService) EXPECT() *MockParseUrlPackageServiceMockRecorder {
	return m.recorder
}

// ParseURLs mocks base method.
func (m *MockParseUrlPackageService) ParseURLs(ctx context.Context, urls []string) ([]rss_parser.RssItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseURLs", ctx, urls)
	ret0, _ := ret[0].([]rss_parser.RssItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseURLs indicates an expected call of ParseURLs.
func (mr *MockParseUrlPackageServiceMockRecorder) ParseURLs(ctx, urls interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseURLs", reflect.TypeOf((*MockParseUrlPackageService)(nil).ParseURLs), ctx, urls)
}
