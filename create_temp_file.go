package paratils

import (
	"io/ioutil"
	"path/filepath"

	"github.com/neumachen/errorx"
)

// CreateTempFile ...
func CreateTempFile(fileContent []byte, directory, dirPrefix, fileName string) (string, string, error) {
	dir, err := ioutil.TempDir(dirPrefix, directory)
	if err != nil {
		return "", "", errorx.New(err)
	}

	tmpfn := filepath.Join(dir, fileName)
	err = ioutil.WriteFile(tmpfn, fileContent, 0666)
	if err != nil {
		return "", "", errorx.New(err)
	}
	return dir, tmpfn, nil
}
