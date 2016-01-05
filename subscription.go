package gopaymill

import (
	"encoding/json"
)

type Subscription struct {
	ID               string      `json:"id"`
	AppId            string      `json:"app_id"`
	CreatedAt        int         `json:"created_at"`
	UpdatedAt        int         `json:"updated_at"`
	Offer            interface{} `json:"offer"`
	Livemode         bool        `json:"livemode"`
	Amount           int         `json:"amount"`
	TempAmount       int         `json:"temp_amount"`
	Currency         string      `json:"currency"`
	Name             string      `json:"name"`
	Interval         string      `json:"interval"`
	TrialStart       int         `json:"trial_start"`
	TrialEnd         int         `json:"trial_end"`
	PeriodOfValidity string      `json:"period_of_validity"`
	EndOfPeriod      int         `json:"end_of_period"`
	NextCaptureAt    int         `json:"next_capture_at"`
	CanceledAt       int         `json:"canceled_at"`
	Payment          interface{} `json:"payment"`
	IsCanceled       bool        `json:"is_canceled"`
	IsDeleted        bool        `json:"is_deleted"`
	Status           string      `json:"status"`
	MandateReference string      `json:"mandate_reference"`
	Client           interface{} `json:"client"`
}

type SubscriptionResponse struct {
	Data Subscription `json:"data"`
	Mode string       `json:"mode"`
}

type SubscriptionAllResponse struct {
	Data      []Subscription `json:"data"`
	DataCount string         `json:"data_count"`
	Mode      string         `json:"mode"`
}

func (sb Subscription) Get(p *Paymill) Subscription {
	pr := NewRequest()
	pr.SetMethod("GET")
	pr.SetUri(APIURI + ServiceUri["Subscription"] + "/" + sb.ID)
	response := p.ClientRequest(&pr)
	obj := SubscriptionResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}

func (sb Subscription) All(p *Paymill) []Subscription {
	pr := NewRequest()
	pr.SetMethod("GET")
	pr.SetUri(APIURI + ServiceUri["Subscription"])
	response := p.ClientRequest(&pr)
	obj := SubscriptionAllResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}

func (sb Subscription) New(p *Paymill) Subscription {
	pr := NewRequest()
	pr.SetMethod("POST")
	pr.SetUri(APIURI + ServiceUri["Payment"])
	params := map[string]interface{}{
		"payment": sb.Payment,
		"client":  sb.Client,
	}
	pr.SetParams(params)
	response := p.ClientRequest(&pr)

	obj := SubscriptionResponse{}
	json.Unmarshal([]byte(response), &obj)
	return obj.Data
}
