package paratils

import (
	"errors"

	"github.com/ParaServices/errgo"
)

// Errx wraps the current e into an errgo.Error if it's not of that type. If
// if is, it ignores and returns it as is. If the e is nil, it returns nil
// (this is handled in the errgo package)
func Errx(e error) error {
	err, ok := e.(*errgo.Error)
	if ok {
		return err
	}
	return errgo.New(e)
}

// NewErrx helper methods for initializing a new error.
func NewErrx(s string) error {
	var err error

	if s != "" {
		err = errors.New(s)
	}

	return errgo.New(err)
}
