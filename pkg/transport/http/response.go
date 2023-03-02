package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/gookit/validate"
	"github.com/simplexpage/rss-reader/pkg/errors"
	"net/http"
)

func EncodeHttpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(FailedAndValidation); ok && f.Validation() != nil {
		encodeValidation(ctx, f.Validation(), w)
		return nil
	}
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		EncodeError(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{"code": http.StatusOK, "message": "OK", "data": response}
	return json.NewEncoder(w).Encode(data)
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errorToCodeHttpAnswer(err))
	data := map[string]interface{}{"code": errorToCodeHttpAnswer(err), "message": err.Error(), "data": []string{}}
	json.NewEncoder(w).Encode(data)
}

func encodeValidation(_ context.Context, v validate.Errors, w http.ResponseWriter) {
	if v == nil {
		panic("encodeValidation with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errorToCodeHttpAnswer(errors.ErrDataValidation))
	data := map[string]interface{}{"code": errorToCodeHttpAnswer(errors.ErrDataValidation), "message": errors.ErrDataValidation.Error(), "data": v}
	json.NewEncoder(w).Encode(data)
}

func errorToCodeHttpAnswer(err error) int {
	switch err {
	case errors.ErrNotFound:
		return http.StatusNotFound
	case errors.ErrAlreadyExists, errors.ErrInconsistentIDs:
		return http.StatusBadRequest
	case errors.ErrDataValidation:
		return http.StatusUnprocessableEntity
	case ErrUnauthorized, ErrInvalidAuthToken, ErrWrongAuthToken, ErrMissingAuthToken, ErrValidAuthToken:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

type FailedAndValidation interface {
	Failed() error
	Validation() validate.Errors
}
