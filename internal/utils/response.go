package utils

import (
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Resp http.ResponseWriter
}

func (r *Response) Text(code int, body string, response_type string) {
	r.Resp.Header().Set("Content-Type", response_type)
	r.Resp.WriteHeader(code)

	io.WriteString(r.Resp, fmt.Sprintf("%s\n", body))

}
