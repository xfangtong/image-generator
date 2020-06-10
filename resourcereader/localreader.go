package resourcereader

import (
	"io"
	"os"
	"strings"
)

const localProtocol string = "local://"

func localOpen(url string) (io.ReadCloser, error) {
	url = strings.TrimPrefix(url, localProtocol)
	return os.Open(url)
}

func init() {
	RegisterReader(localProtocol, localOpen)
}
