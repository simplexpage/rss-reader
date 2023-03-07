package transport

import (
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/simplexpage/rss-reader/internal/reader/delivery/reqresp/form"
	"github.com/simplexpage/rss-reader/internal/reader/domain/adapter"
	"github.com/simplexpage/rss-reader/internal/reader/domain/service"
	"github.com/simplexpage/rss-reader/internal/reader/endpoint"
	httpUtil "github.com/simplexpage/rss-reader/pkg/transport/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

func TestParseUrlsHandler(t *testing.T) {
	testTable := []struct {
		name               string
		inputParseUrls     string
		parseUrlsForm      form.ParseUrlsForm
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name:               "Empty body",
			inputParseUrls:     ``,
			expectedResponse:   `{"code":422,"data":[],"message":"data validation failed"}`,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:               "data validation",
			inputParseUrls:     `{"urls":[]}`,
			expectedResponse:   `{"code":422,"data":{"urls":{"required":"urls is required to not be empty"}},"message":"data validation failed"}`,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:               "success",
			inputParseUrls:     `{"urls":["https://tsn.ua/rss/full.rss"]}`,
			expectedResponse:   ``,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			parseUrlPackageAdapter := adapter.NewParseUrlPackageAdapter()
			serviceReader := service.New(log.NewNopLogger(), parseUrlPackageAdapter)
			serviceEndpoints := endpoint.NewServerEndpoints(serviceReader, log.NewNopLogger())
			serviceHttpHandler := NewHTTPHandler(serviceEndpoints, log.With(log.NewNopLogger(), "component", "HTTP"))
			srv := httptest.NewServer(serviceHttpHandler)
			defer srv.Close()

			req, _ := http.NewRequest("POST", "/reader/parse", strings.NewReader(testCase.inputParseUrls))
			w := httptest.NewRecorder()
			serviceHttpHandler.ServeHTTP(w, req)
			body, _ := ioutil.ReadAll(w.Body)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponse != "" {
				assert.Equal(t, testCase.expectedResponse, strings.TrimSpace(string(body)))
			}
		})
	}
}
