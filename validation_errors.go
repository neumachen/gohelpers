package paratils

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap/zapcore"
)

// validationError represents the error object that is specified on the
// http://jsonapi.org
type validationError struct {
	ID     string       `json:"error_id"`
	Links  *errorLinks  `json:"links,omitempty"`
	Status string       `json:"status,omitempty"`
	Code   string       `json:"code,omitempty"`
	Title  string       `json:"title,omitempty"`
	Detail string       `json:"detail,omitempty"`
	Source *errorSource `json:"source,omitempty"`
	Meta   interface{}  `json:"meta,omitempty"`
}

// MarshalLogObject ...
func (e validationError) MarshalLogObject(kv zapcore.ObjectEncoder) error {
	kv.AddString("error_id", e.ID)
	if e.Links != nil {
		kv.AddObject("links", e.Links)
	}
	if e.Status != "" {
		kv.AddString("status", e.Status)
	}
	if e.Code != "" {
		kv.AddString("code", e.Code)
	}
	if e.Title != "" {
		kv.AddString("title", e.Title)
	}
	if e.Detail != "" {
		kv.AddString("detail", e.Detail)
	}
	if e.Source != nil {
		kv.AddObject("source", e.Source)
	}

	return nil
}

// NewError ...
func NewError() (*validationError, error) {
	e := &validationError{
		ID: uuid.NewV4().String(),
	}
	return e, nil
}

// errorLinks is used to provide an About URL that leads to
// further details about the particular occurrence of the problem.
// for more information see http://jsonapi.org/format/#error-objects
type errorLinks struct {
	About string `json:"about,omitempty"`
}

// MarshalLogObject ...
func (e errorLinks) MarshalLogObject(kv zapcore.ObjectEncoder) error {
	kv.AddString("about", e.About)
	return nil
}

// errorSource is used to provide references to the source of an error.
// The Pointer is a JSON Pointer to the associated entity in the request
// document.
// The Parameter is a string indicating which query parameter caused the error.
// for more information see http://jsonapi.org/format/#error-objects
type errorSource struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
	Value     string `json:"value,omitempty"`
}

// MarshalLogObject ...
func (e errorSource) MarshalLogObject(kv zapcore.ObjectEncoder) error {
	if e.Pointer != "" {
		kv.AddString("pointer", e.Pointer)
	}
	if e.Parameter != "" {
		kv.AddString("parameter", e.Parameter)
	}
	if e.Value != "" {
		kv.AddString("value", e.Value)
	}
	return nil
}

// validationErrors ...
type validationErrors []validationError

// Marshal ...
func (v validationErrors) Marshal() ([]byte, error) {
	return json.Marshal(v)
}

// MarshalLogArray ...
func (v validationErrors) MarshalLogArray(kv zapcore.ArrayEncoder) error {
	for _, ve := range v {
		err := kv.AppendObject(ve)
		if err != nil {
			return err
		}
	}
	return nil
}

// ValidationError ...
type ValidationError interface {
	Error() string
	ValidationErrors() validationErrors
	MarshalLogObject(zapcore.ObjectEncoder) error
}

type errValidation struct {
	validationErrors validationErrors
}

func (e *errValidation) Error() string {
	b, err := e.validationErrors.Marshal()
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Validation Error: %s", b)
}

func (e *errValidation) ValidationErrors() validationErrors {
	return e.validationErrors
}

func (e *errValidation) MarshalLogObject(kv zapcore.ObjectEncoder) error {
	return kv.AddArray("error", e.validationErrors)
}

// NewErrValidationErrors ...
func NewErrValidationErrors(valErrors validationErrors) ValidationError {
	return &errValidation{
		validationErrors: valErrors,
	}
}
