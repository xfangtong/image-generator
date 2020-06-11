package resourcereader

import (
	"io"
	"os"
	"strings"
)

const localProtocol string = "local://"

func localOpen(url string) (io.ReadCloser, error) {
	path := strings.TrimPrefix(url, localProtocol)
	return os.Open(path)
}

func init() {
	RegisterReader(localProtocol, localOpen)
}
