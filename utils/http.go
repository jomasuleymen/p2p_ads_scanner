package utils

import "net/http"

func SafeCloseRequestResponseBody(req *http.Request, resp *http.Response) {
	if req != nil && req.Body != nil {
		req.Body.Close()
	}

	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
}
