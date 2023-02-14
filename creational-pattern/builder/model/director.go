package model

/**
  可以看到创建者模式有点类似工厂模式，他们最大的区别就是在这个director里
*/

type Director struct {
    builder IBuilder
}

func NewDirector(b IBuilder) *Director {
    return &Director{
        builder: b,
    }
}

func (d *Director) Build() *Music {
    music := &Music{}
    d.builder.BuildMusicStep1(music)
    d.builder.BuildMusicStep2(music)
    return music
}
