package model

import "fmt"

// Drugstore 中间件
type Drugstore struct {
    Next Handler
}

func (c *Drugstore) Handle(client *Client) {
    fmt.Println("Drugstore handling...")
    if client.getMedicineDone {
        c.Next.Handle(client)
        return
    }
    client.getMedicineDone = true
    c.Next.Handle(client)
}

func (c *Drugstore) SetNext(handler Handler) {
    c.Next = handler
}
