package paratils

import "github.com/ParaServices/errgo"

// Errx ...
func Errx(e error) error {
	err, ok := e.(*errgo.Error)
	if ok {
		return err
	}
	return errgo.New(e)
}
