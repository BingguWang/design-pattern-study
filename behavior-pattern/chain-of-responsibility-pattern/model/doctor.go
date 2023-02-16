package model

import "fmt"

// Doctor 中间件
type Doctor struct {
    Next Handler
}

func (c *Doctor) Handle(client *Client) {
    fmt.Println("Doctor handling...")
    if client.doctorCheckUpDone {
        c.Next.Handle(client)
        return
    }
    client.doctorCheckUpDone = true
    c.Next.Handle(client)
}

func (c *Doctor) SetNext(handler Handler) {
    c.Next = handler
}
