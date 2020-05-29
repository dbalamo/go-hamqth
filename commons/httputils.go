package commons

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func SetHttpTransportParams() {
	defaultTransportPointer, ok := (http.DefaultTransport).(*http.Transport)
	if !ok {
		InnerPrintln("defaultRoundTripper not an *http.Transport")
	}
	defaultTransportPointer.MaxIdleConns = 100
	defaultTransportPointer.MaxIdleConnsPerHost = 100

}

// newRequest creates an HTTP request.
func newRequest(method, path string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, path, body)
}

func DoGet(url string, headerParams map[string]string) ([]byte, error) {
	var req *http.Request
	var resp *http.Response
	var err error

	if req, err = newRequest("GET", url, nil); err != nil {
		InnerPrintln("ERROR in http request creation", err.Error())
		return nil, errors.New("ERROR in req creation")
	}

	for k, v := range headerParams {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	if resp, err = client.Do(req); err != nil {
		InnerPrintln("ERROR in httpClient.Do", err.Error())
		return nil, errors.New("ERROR in req.Do")
	}
	buf := make([]byte, 65535)
	if buf, err = ioutil.ReadAll(resp.Body); err != nil {
		InnerPrintln("ERROR in http response body read", err.Error())
		return nil, errors.New("ERROR in http response body read")
	}

	return buf, nil
}
