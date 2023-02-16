package model

import "fmt"

// Cashier 收银中间件
type Cashier struct {
    Next Handler
}

func (c *Cashier) Handle(client *Client) {
    fmt.Println("Cashier handling...")
    if client.paymentDone {
        c.Next.Handle(client)
        return
    }
    client.paymentDone = true
    c.Next.Handle(client)
}

func (c *Cashier) SetNext(handler Handler) {
    c.Next = handler
}
