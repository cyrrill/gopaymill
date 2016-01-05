package gopaymill

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type PaymillRequest struct {
	Method string
	Uri    string
	Body   url.Values
}

func NewRequest() PaymillRequest {
	r := PaymillRequest{}
	r.Body = make(url.Values)
	return r
}

func (p *Paymill) ClientRequest(pr *PaymillRequest) string {
	var request *http.Request
	var err error
	if pr.Method == "POST" {
		data := pr.Body
		reader := bytes.NewBufferString(data.Encode())
		request, err = http.NewRequest(pr.Method, pr.Uri, reader)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	} else {
		request, err = http.NewRequest(pr.Method, pr.Uri, bytes.NewBufferString(""))
	}
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(p.PrivateKey, p.PrivateKey)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func (pr *PaymillRequest) SetMethod(method string) {
	pr.Method = method
}

func (pr *PaymillRequest) SetUri(uri string) {
	pr.Uri = uri
}

func (pr *PaymillRequest) SetParams(params map[string]interface{}) {
	for k, v := range params {
		switch v.(type) {
		case string:
			if len(v.(string)) > 0 {
				pr.Body.Set(k, v.(string))
			}
		case int32, int64:
			if v.(int) > 0 {
				pr.Body.Set(k, v.(string))
			}
		default:
		}
	}
}

type count struct {
	Active   int `json:active`
	Inactive int `json:inactive`
}
