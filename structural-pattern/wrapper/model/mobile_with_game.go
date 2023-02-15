package model

import "fmt"

// MobileWithGame 可以玩游戏的手机, 装饰1
// 可以看到不需要修改Mobile struct 但是我们给Mobile新增了功能
type MobileWithGame struct {
    Mobile IMobile
}

func (m *MobileWithGame) Function() {
    m.Mobile.Function()
    fmt.Println("play games")
}
