package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil { // unmarshal untuk mengkonversi data yang diterima dalam bentuk object seperti map => map[string]interface{}
			return
		}
	}
}
