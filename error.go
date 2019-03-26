package paratils

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// ResponseCodeError ...
type ResponseCodeError interface {
	Error() string
	ResponseCode() int
	MarshalLogObject(zapcore.ObjectEncoder) error
}

type errExpectedResponse struct {
	expectedResponseCode int
	actualResponseCode   int
	rawURL               string
}

func (e *errExpectedResponse) Error() string {
	return fmt.Sprintf(
		"Expected response code: %v, actual response code: %v, query: %v",
		e.expectedResponseCode,
		e.actualResponseCode,
		e.rawURL,
	)
}

// ResponseCode ...
func (e *errExpectedResponse) ResponseCode() int {
	return e.actualResponseCode
}

func (e *errExpectedResponse) MarshalLogObject(kv zapcore.ObjectEncoder) error {
	kv.AddInt("expected_response_code", e.expectedResponseCode)
	kv.AddInt("actual_response_code", e.actualResponseCode)
	kv.AddString("raw_url", e.rawURL)
	return nil
}

// NewErrExpectedResponse ...
func NewErrExpectedResponse(expectedRespCode, actualRespCode int, rawURL string) ResponseCodeError {
	return &errExpectedResponse{
		expectedResponseCode: expectedRespCode,
		actualResponseCode:   actualRespCode,
		rawURL:               rawURL,
	}
}
