package firstanimal

import "fmt"

type Birds struct {
	Id_b     int
	Name     string
	Fly      string
	Song     string
	Canswimm bool
	Wings    bool
}

func (b Birds) Move() {
	fmt.Println(b.Name, "can move", b.Fly)

}
func (b Birds) Sound() {
	fmt.Println(b.Name, "can sound", b.Song)
}
func (b *Birds) Canswim() {
	var answer = b.Canswimm
	if answer == true {
		fmt.Println("this animal can swim")
	} else {
		fmt.Println("this animal can not swim")
	}
}
func (b *Birds) CanFly() {
	var answ = b.Wings
	if answ == true {
		fmt.Println(b.Name, "can fly")
	} else {
		fmt.Println(b.Name, "can not fly")
	}
}
