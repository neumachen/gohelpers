package paratils

import "github.com/ParaServices/errgo"

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
