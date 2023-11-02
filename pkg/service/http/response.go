package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	resp *http.Response
}

func (r *Response) UnmarshalJSON(v any) error {
	b, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (r *Response) Close() {
	_ = r.resp.Body.Close()
}
