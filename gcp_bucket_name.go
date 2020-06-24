package paratils

import (
	"strings"

	"github.com/ParaServices/paraenv"
)

// Para ...
const Para = "para"

// GCPBucketName ...
func GCPBucketName(appEnv paraenv.AppEnv, folders []string) string {
	ss := []string{Para, appEnv.String()}
	ss = append(ss, folders...)

	return strings.Join(ss, "-")
}
