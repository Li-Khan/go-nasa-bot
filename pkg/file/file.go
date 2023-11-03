package file

import (
	"io/ioutil"
	"os"
)

func OpenAndOverwriteFile(name string, text string) (string, error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return "", err
	}
	_, err = f.Write([]byte(text))
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
