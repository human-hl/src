package firstanimal

import "fmt"

type Animals struct {
	Name          string
	Run           string
	Frighteningly string
	Canswimm      bool
}

func (a Animals) Move() {
	fmt.Println(a.Name, "can move", a.Run)
}
func (a Animals) Sound() {
	fmt.Println(a.Name, "can sound", a.Frighteningly)
}

func (a *Animals) Canswim() {
	var answer = a.Canswimm
	if answer == true {
		fmt.Println("this animal can swim")
	} else {
		fmt.Println("this animal can not swim")
	}
}
