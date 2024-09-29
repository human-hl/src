package firstanimal

import "fmt"

type Amphibians struct {
	Id_am      int
	Name       string
	Runandswim string
	Strange    string
	Canswimm   bool
}

func (a *Amphibians) Canswim() {
	var answer = a.Canswimm
	if answer == true {
		fmt.Println("this animal can swim")
	} else {
		fmt.Println("this animal can not swim")
	}
}

func (a Amphibians) Move() {
	fmt.Println(a.Name, "can move", a.Runandswim)

}
func (a Amphibians) Sound() {
	fmt.Println(a.Name, "can sound", a.Strange)
}
