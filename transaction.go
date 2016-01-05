package gopaymill

import (
	"encoding/json"
)

type Transaction struct {
	ID                string        `json:"id"`
	AppId             string        `json:"app_id"`
	CreatedAt         int32         `json:"created_at"`
	UpdatedAt         int32         `json:"updated_at"`
	Amount            string        `json:"amount"`
	OriginAmount      int32         `json:"origin_amount"`
	Status            string        `json:"status"`
	Description       string        `json:"description"`
	Livemode          bool          `json:"livemode"`
	Refunds           []interface{} `json:"-"`
	Client            Client        `json:"client"`
	Currency          string        `json:"currency"`
	ResponseCode      int           `json:"response_code"`
	ShortID           string        `json:"short_id"`
	Is_Fraud          bool          `json:"is_fraud"`
	Invoices          []interface{} `json:"invoices"`
	Items             []interface{} `json:"items"`
	ShippingAddress   interface{}   `json:"-"`
	BillingAddress    interface{}   `json:"-"`
	Preauthorization  interface{}   `json:"-"`
	Fees              []interface{} `json:"-"`
	Payment           Payment       `json:"payemnt"`
	MandateReference  interface{}   `json:"-"`
	IsRefundable      bool          `json:"is_redfundable"`
	IsMarkableAsFraud bool          `json:"is_markable_as_fraud"`
	Token             string        `json:"-"`
}

type TransactionResponse struct {
	Data Transaction `json:"data"`
	Mode string      `json:"mode"`
}

type TransactionAllResponse struct {
	Data      []Transaction `json:"data"`
	DataCount string        `json:"data_count"`
	Mode      string        `json:"mode"`
}

func (t Transaction) Get(p *Paymill) Transaction {
	pr := NewRequest()
	pr.SetMethod("GET")
	pr.SetUri(APIURI + ServiceUri["Transaction"] + "/" + t.ID)
	response := p.ClientRequest(&pr)
	obj := TransactionResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}

func (t Transaction) All(p *Paymill) []Transaction {
	pr := NewRequest()
	pr.SetMethod("GET")
	pr.SetUri(APIURI + ServiceUri["Transaction"])
	response := p.ClientRequest(&pr)
	obj := TransactionAllResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}

func (t Transaction) New(p *Paymill) Transaction {
	pr := NewRequest()
	pr.SetMethod("POST")
	pr.SetUri(APIURI + ServiceUri["Transaction"])
	params := map[string]interface{}{
		"amount":      t.Amount,
		"currency":    t.Currency,
		"description": t.Description,
	}
	if len(t.Client.ID) > 0 {
		params["client"] = t.Client
	}
	if len(t.Payment.ID) > 0 {
		params["payment"] = t.Payment
	}
	if len(t.Token) > 0 {
		params["token"] = t.Token
	}
	pr.SetParams(params)
	response := p.ClientRequest(&pr)

	obj := TransactionResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}
