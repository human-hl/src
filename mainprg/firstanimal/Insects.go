package firstanimal

import "fmt"

type Insects struct {
	Name       string
	Crawl      string
	Irritation string
	Canswimm   bool
	Wings      bool
}

func (i Insects) Move() {
	fmt.Println(i.Name, "can move", i.Crawl)
}
func (i Insects) Sound() {
	fmt.Println(i.Name, "can sound", i.Irritation)
}
func (i *Insects) Canswim() {
	var answer = i.Canswimm
	if answer == true {
		fmt.Println("this animal can swim")
	} else {
		fmt.Println("this animal can not swim")
	}
}
func (i *Insects) CanFly() {
	var answ = i.Wings
	if answ == true {
		fmt.Println(i.Name, "can fly")
	} else {
		fmt.Println(i.Name, "can not fly")
	}
}
