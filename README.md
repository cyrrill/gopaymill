# Go Paymill

Here is how to connect to and interact with the Paymill API in Go.

Each model type is mapped to a struct, which has New() Get() All() methods.

* Client
* Payment
* Transaction
* Subscription


## Connect

Import the package into your app and 

```
package yourapp

import (
    "fmt"
    gp "github.com/cyrrill/gopaymill"
)

// Creates a client connection using a key: https://app.paymill.com/development/api-keys
p := gp.Paymill{PrivateKey: "000000000000000000"}
```

## Models
The create individual structs and call the action method on the Paymill connection &p
```
// Creates a new client 
client := gp.Client{Description: "Golang", Email: "test@golang.org"}.New(&p)

```

Reuse values directly in following calls
```
// Create a new payment object. Get the token from Paymills BridgeJS, client ID comes previous call
payment := gp.Payment{Token:"tok_000000000000", Client: client.ID}.New(&p)

// Create a recurring subscritpion with new payment method
subscription := gp.Subscription{
    Offer:   "offer_0000000000000000", 
    Client:  client.ID, 
    Payment: payment.ID,
}.New(&p)

```

## Output

The values in all the models are directly accesible after the call as:
```
fmt.Println("Subscription ID: "+subscription.ID)

fmt.Println("Client email: "+client.Email)

```

## Get one, Get all

```
// Lookup by key
payment := gp.Payment{ID:"pay_000000000000"}.Get(&p)


// Get all  []Transaction
transaction := gp.Transaction{}.All(&p)


```
