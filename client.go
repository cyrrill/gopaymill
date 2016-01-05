package gopaymill

import (
	"encoding/json"
)

type Client struct {
	ID           string         `json:"id"`
	AppId        string         `json:"app_id"`
	CreatedAt    int            `json:"created_at"`
	UpdatedAt    int            `json:"updated_at"`
	Email        string         `json:"email"`
	Description  string         `json:"description"`
	Payment      []Payment      `json:"payment"`
	Subscription []Subscription `json:"subscription"`
}

type ClientResponse struct {
	Data Client `json:data`
	Mode string `json:mode`
}

type ClientAllResponse struct {
	Data []Client `json:"data"`
	Mode string   `json:"mode"`
}

func (c Client) Get(p *Paymill) Client {
	pr := NewRequest()
	pr.SetMethod("GET")
	pr.SetUri(APIURI + ServiceUri["Client"] + "/" + c.ID)
	response := p.ClientRequest(&pr)
	obj := ClientResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}

func (c Client) All(p *Paymill) []Client {
	pr := NewRequest()
	pr.SetMethod("GET")
	pr.SetUri(APIURI + ServiceUri["Client"])
	response := p.ClientRequest(&pr)
	obj := ClientAllResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}

func (c Client) New(p *Paymill) Client {
	pr := NewRequest()
	pr.SetMethod("POST")
	pr.SetUri(APIURI + ServiceUri["Client"])
	params := map[string]interface{}{
		"email":       c.Email,
		"description": c.Description,
	}
	pr.SetParams(params)
	response := p.ClientRequest(&pr)

	obj := ClientResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}
