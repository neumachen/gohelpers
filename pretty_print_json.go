package paratils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// PrettyPrintJSON ...
func PrettyPrintJSON(b []byte) error {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, b, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(prettyJSON.Bytes()))
	return nil
}
