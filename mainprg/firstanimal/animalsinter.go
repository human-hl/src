package firstanimal

type Animal interface {
	Move()
	Sound()
}

type Swimmer interface {
	Canswim()
}

type Fly interface {
	CanFly()
}

func About(a Animal) {
	a.Move()
	a.Sound()
}
func Asking(s Swimmer) {
	s.Canswim()
}
func Speshfly(f Fly) {
	f.CanFly()
}
