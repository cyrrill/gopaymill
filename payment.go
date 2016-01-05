package gopaymill

import (
	"encoding/json"
)

type Payment struct {
	ID                          string      `json:"id"`
	AppId                       string      `json:"app_id"`
	CreatedAt                   int         `json:"created_at"`
	UpdatedAt                   int         `json:"updated_at"`
	PaymentType                 string      `json:"type"`
	Client                      interface{} `json:"client"`
	IsRecurring                 bool        `json:"is_recurring"`
	IsUsableForPreauthorization bool        `json:"is_usable_for_preauthorization"`
	CardType                    string      `json:"card_type"`
	Country                     string      `json:"country"`
	ExpireMonth                 string      `json:"expire_month"`
	ExpireYear                  string      `json:"expire_year"`
	CardHolder                  string      `json:"card_holder"`
	Last4                       string      `json:"last4"`
	Code                        string      `json:"code"`
	Account                     string      `json:"account"`
	Holder                      string      `json:"holder"`
	Iban                        string      `json:"iban"`
	Bic                         string      `json:"bic"`
	Token                       string      `json:"-"`
}

type PaymentResponse struct {
	Data Payment `json:"data"`
	Mode string  `json:"mode"`
}

type PaymentAllResponse struct {
	Data      []Payment `json:"data"`
	DataCount string    `json:"data_count"`
	Mode      string    `json:"mode"`
}

func (py Payment) Get(p *Paymill) Payment {
	pr := NewRequest()
	pr.SetMethod("GET")
	pr.SetUri(APIURI + ServiceUri["Payment"] + "/" + py.ID)
	response := p.ClientRequest(&pr)
	obj := PaymentResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}

func (py Payment) All(p *Paymill) []Payment {
	pr := NewRequest()
	pr.SetMethod("GET")
	pr.SetUri(APIURI + ServiceUri["Payment"])
	response := p.ClientRequest(&pr)
	obj := PaymentAllResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}

func (py Payment) New(p *Paymill) Payment {
	pr := NewRequest()
	pr.SetMethod("POST")
	pr.SetUri(APIURI + ServiceUri["Payment"])
	params := map[string]interface{}{
		"token":  py.Token,
		"client": py.Client,
	}
	pr.SetParams(params)
	response := p.ClientRequest(&pr)

	obj := PaymentResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}
