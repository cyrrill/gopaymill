package gopaymill

const APIURI = "https://api.paymill.com/v2.1"

var ServiceUri = map[string]string{
	"Client":       "/clients",
	"Payment":      "/payments",
	"Subscription": "/subscriptions",
	"Transaction":  "/transactions",
}

type Paymill struct {
	PrivateKey string
}
