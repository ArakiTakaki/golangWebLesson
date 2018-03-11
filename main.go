package main

import "github.com/ArakiTakaki/golangWebLesson/router"

func main() {

	r := router.GetRouter()
	// ルーターの設定は一度しか読み込まれないっぽい？
	r.Run(":8000")

}

/*
メソッドの作り方
type Perple struct{ height, weight float32 }

func New(h, w float32) Perple {
	p := Perple{h, w}
	return p
}

func (p Perple) BMI() float32 {
	return p.weight / (p.height * p.height)
}
*/
