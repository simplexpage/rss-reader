package transport

import (
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/gookit/validate"
	"github.com/simplexpage/rss-reader/internal/reader/delivery/reqresp/form"
	"github.com/simplexpage/rss-reader/internal/reader/domain/model"
	"github.com/simplexpage/rss-reader/internal/reader/domain/service/mocks"
	"github.com/simplexpage/rss-reader/internal/reader/endpoint"
	httpUtil "github.com/simplexpage/rss-reader/pkg/transport/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestParseUrlsHandlerSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputParseUrls := `{"urls":["https://tsn.ua/rss/full.rss"]}`
	inputParseUrlsFormMock := form.ParseUrlsForm{
		Urls: []string{"https://tsn.ua/rss/full.rss"},
	}
	mockResponse := []model.Item{
		{
			Title:       "title",
			Source:      "source",
			SourceUrl:   "sourceUrl",
			Link:        "link",
			PublishDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Description: "description",
		},
	}
	expectedStatusCode := http.StatusOK
	expectedResponse := `{"code":200,"data":{"items":[{"title":"title","source":"source","source_url":"sourceUrl","link":"link","publish_date":"2020-01-01 00:00:00","description":"description"}]},"message":"OK"}`

	serviceReaderMock := mocks.NewMockService(ctrl)
	serviceReaderMock.EXPECT().ParseUrls(gomock.Any(), inputParseUrlsFormMock).Return(mockResponse, nil, nil).Times(1)

	serviceEndpoints := endpoint.NewServerEndpoints(serviceReaderMock, log.NewNopLogger())
	serviceHttpHandler := NewHTTPHandler(serviceEndpoints, log.With(log.NewNopLogger(), "component", "HTTP"))
	srv := httptest.NewServer(serviceHttpHandler)
	defer srv.Close()

	req, _ := http.NewRequest("POST", "/reader/parse", strings.NewReader(inputParseUrls))
	w := httptest.NewRecorder()
	serviceHttpHandler.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, expectedStatusCode, w.Code)
	assert.Equal(t, expectedResponse, strings.TrimSpace(string(body)))
}

func TestParseUrlsHandlerValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputParseUrls := `{"urls":[]}`
	var mockResponse []model.Item
	expectedStatusCode := http.StatusUnprocessableEntity
	expectedResponse := `{"code":422,"data":{},"message":"data validation failed"}`

	serviceReaderMock := mocks.NewMockService(ctrl)
	serviceReaderMock.EXPECT().ParseUrls(gomock.Any(), gomock.Any()).Return(mockResponse, errors.New("data validation failed"), validate.Errors{}).Times(1)

	serviceEndpoints := endpoint.NewServerEndpoints(serviceReaderMock, log.NewNopLogger())
	serviceHttpHandler := NewHTTPHandler(serviceEndpoints, log.With(log.NewNopLogger(), "component", "HTTP"))
	srv := httptest.NewServer(serviceHttpHandler)
	defer srv.Close()

	req, _ := http.NewRequest("POST", "/reader/parse", strings.NewReader(inputParseUrls))
	w := httptest.NewRecorder()
	serviceHttpHandler.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, expectedStatusCode, w.Code)
	assert.Equal(t, expectedResponse, strings.TrimSpace(string(body)))
}

func TestParseUrlsHandlerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedStatusCode := http.StatusUnprocessableEntity
	expectedResponse := `{"code":422,"data":[],"message":"data validation failed"}`

	serviceReaderMock := mocks.NewMockService(ctrl)

	serviceEndpoints := endpoint.NewServerEndpoints(serviceReaderMock, log.NewNopLogger())
	serviceHttpHandler := NewHTTPHandler(serviceEndpoints, log.With(log.NewNopLogger(), "component", "HTTP"))
	srv := httptest.NewServer(serviceHttpHandler)
	defer srv.Close()

	req, _ := http.NewRequest("POST", "/reader/parse", strings.NewReader(""))
	w := httptest.NewRecorder()
	serviceHttpHandler.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, expectedStatusCode, w.Code)
	assert.Equal(t, expectedResponse, strings.TrimSpace(string(body)))
}

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpUtil.HealthCheckHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
