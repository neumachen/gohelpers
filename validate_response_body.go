package paratils

import (
	"net/http"
	"strings"

	"github.com/ParaServices/gblcodes"
	uuid "github.com/satori/go.uuid"
	"github.com/xeipuuv/gojsonschema"
	"golang.org/x/net/webdav"
)

// ValidateResponseBody ...
func ValidateResponseBody(snapshot, schema []byte) error {
	result, err := gojsonschema.Validate(
		gojsonschema.NewBytesLoader(schema),
		gojsonschema.NewBytesLoader(snapshot),
	)
	if err != nil {
		return Errx(err)
	}

	if !result.Valid() {
		var ee validationErrors
		for _, resErr := range result.Errors() {
			dc := gblcodes.HTTPJSONSchemaValidationFailed
			e := validationError{
				ID:     uuid.NewV4().String(),
				Status: http.StatusText(webdav.StatusUnprocessableEntity),
				Code:   dc.Code(),
				Title:  dc.Message(),
				Detail: resErr.Description(),
				Source: &errorSource{
					Pointer: strings.Replace((resErr.Context().String() + "/" + resErr.Field()), ".", "/", -1),
				},
			}
			ee = append(ee, e)
		}
		return NewErrValidationErrors(ee)
	}

	return nil
}
