package model

import "fmt"

// Reception 中间件
type Reception struct {
    Next Handler
}

func (c *Reception) Handle(client *Client) {
    fmt.Println("Reception handling...")
    if client.registrationDone {
        c.Next.Handle(client)
        return
    }
    client.registrationDone = true
    c.Next.Handle(client)
}

func (c *Reception) SetNext(handler Handler) {
    c.Next = handler
}
