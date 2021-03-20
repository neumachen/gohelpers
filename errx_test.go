package paratils

import (
	"errors"
	"testing"

	"github.com/ParaServices/errgo"
	"github.com/stretchr/testify/require"
)

func TestErrx(t *testing.T) {
	t.Run("is *errgo.Error", func(t *testing.T) {
		e := errgo.New(nil)
		errx := Errx(e)

		switch errx.(type) {
		case *errgo.Error:
			require.True(t, true)
		default:
			require.FailNow(t, "is not errgo.Error")
		}
	})

	t.Run("is not *errgo.Error", func(t *testing.T) {
		errx := Errx(errors.New("test"))

		switch errx.(type) {
		case *errgo.Error:
			require.True(t, true)
		default:
			require.FailNow(t, "is not errgo.Error")
		}
	})

	t.Run("error is nil", func(t *testing.T) {
		errx := Errx(nil)
		require.Nil(t, errx)

		var err error
		errx = Errx(err)
		require.Nil(t, errx)
	})
}
