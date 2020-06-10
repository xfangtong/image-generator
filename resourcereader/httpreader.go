package resourcereader

import (
	"io"
	"net/http"
)

const httpProtocol string = "http://"
const httpsProtocol string = "https://"

func httpOpen(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return response.Body, err
}

func init() {
	RegisterReader(httpProtocol, httpOpen)
	RegisterReader(httpsProtocol, httpOpen)
}
