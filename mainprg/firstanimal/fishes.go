package firstanimal

import "fmt"

type Fishes struct {
	Name     string
	Swim     string
	Glugglug string
	Canswimm bool
}

func (f Fishes) Move() {
	fmt.Println(f.Name, "can move", f.Swim)
}
func (f Fishes) Sound() {
	fmt.Println(f.Name, "can sound", f.Glugglug)
}
func (f *Fishes) Canswim() {
	var answer = f.Canswimm
	if answer == true {
		fmt.Println("this animal can swim")
	} else {
		fmt.Println("this animal can not swim")
	}
}
