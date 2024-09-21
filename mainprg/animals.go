package main

import (
	"fmt"
	"nadonado/mainprg/erroranimals"
	"nadonado/mainprg/firstanimal"
	"os"
)

func main() {
	var cat = firstanimal.Animals{
		Name:          "motis",
		Run:           "slow",
		Frighteningly: "miow",
		Canswimm:      false,
	}
	cat2 := &cat
	firstanimal.Asking(cat2)
	firstanimal.About(cat)
	var myxa = firstanimal.Insects{
		Name:       "myxa",
		Crawl:      "fast",
		Irritation: "bzzz",
		Canswimm:   false,
		Wings:      true,
	}
	myxa2 := &myxa
	firstanimal.About(myxa)
	firstanimal.Asking(myxa2)
	firstanimal.Speshfly(myxa2)
	var answer string
	fmt.Print("какой тип животног(насекомое, птица, дикое, амфибия, рыба)")
	fmt.Fscan(os.Stdin, &answer)
	if _, err := erroranimals.ErrTypeAnimal(answer); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid animal type entered")
	}
	switch answer {
	case "насекомое":
		var Insect = firstanimal.Insects{}
		var name string
		var crawl string
		var irritation string
		var canswim bool
		var wings bool

		fmt.Print("what is it? ")
		fmt.Fscan(os.Stdin, &name)
		Insect.Name = name
		fmt.Print("How is it crowl? ")
		fmt.Fscan(os.Stdin, &crawl)
		Insect.Crawl = crawl
		fmt.Print("How is it irritation? ")
		fmt.Fscan(os.Stdin, &irritation)
		Insect.Irritation = irritation
		fmt.Print("it can swim? true of false ")
		fmt.Fscan(os.Stdin, &canswim)
		if _, err := erroranimals.Errtypeboolansw(canswim); err != nil {
			fmt.Println(err)
		} else {
			Insect.Canswimm = canswim
		}
		fmt.Print("it have wings? true of false ")
		fmt.Fscan(os.Stdin, &wings)
		if _, err := erroranimals.Errtypeboolansw(wings); err != nil {
			fmt.Println(err)
		} else {

			Insect.Wings = wings
		}
		firstanimal.About(Insect)
		firstanimal.Asking(&Insect)
		firstanimal.Speshfly(&Insect)
		fmt.Println(Insect)
	case "птица":
		var Bird = firstanimal.Birds{}
		var name string
		var fly string
		var song string
		var canswim bool
		var wings bool

		fmt.Print("what is it? ")
		fmt.Fscan(os.Stdin, &name)
		Bird.Name = name
		fmt.Print("How is it fly? ")
		fmt.Fscan(os.Stdin, &fly)
		Bird.Fly = fly
		fmt.Print("How is it song? ")
		fmt.Fscan(os.Stdin, &song)
		Bird.Song = song
		fmt.Print("it can swim? true of false ")
		fmt.Fscan(os.Stdin, &canswim)
		if _, err := erroranimals.Errtypeboolansw(canswim); err != nil {
			fmt.Println(err)
		} else {
			Bird.Canswimm = canswim
		}
		fmt.Print("it have wings? true of false ")
		fmt.Fscan(os.Stdin, &wings)
		if _, err := erroranimals.Errtypeboolansw(wings); err != nil {
			fmt.Println(err)
		} else {

			Bird.Wings = wings
		}
		firstanimal.About(Bird)
		firstanimal.Asking(&Bird)
		firstanimal.Speshfly(&Bird)
		fmt.Println(Bird)
	case "дикое":
		var Animal = firstanimal.Animals{}
		var name string
		var run string
		var frighteningly string
		var canswim bool

		fmt.Print("what is it? ")
		fmt.Fscan(os.Stdin, &name)
		Animal.Name = name
		fmt.Print("How is it run? ")
		fmt.Fscan(os.Stdin, &run)
		Animal.Run = run
		fmt.Print("How is it frighteningly? ")
		fmt.Fscan(os.Stdin, &frighteningly)
		Animal.Frighteningly = frighteningly
		fmt.Print("it can swim? true of false ")
		fmt.Fscan(os.Stdin, &canswim)
		if _, err := erroranimals.Errtypeboolansw(canswim); err != nil {
			fmt.Println(err)
		} else {
			Animal.Canswimm = canswim
		}
		firstanimal.About(Animal)
		firstanimal.Asking(&Animal)
		fmt.Println(Animal)
	case "амфибия":
		var Amphibian = firstanimal.Amphibians{}
		var name string
		var runandswim string
		var strange string
		var canswim bool

		fmt.Print("what is it? ")
		fmt.Fscan(os.Stdin, &name)
		Amphibian.Name = name
		fmt.Print("How is it run and swim? ")
		fmt.Fscan(os.Stdin, &runandswim)
		Amphibian.Runandswim = runandswim
		fmt.Print("How is it strange sound? ")
		fmt.Fscan(os.Stdin, &strange)
		Amphibian.Strange = strange
		fmt.Print("it can swim? true of false ")
		fmt.Fscan(os.Stdin, &canswim)
		if _, err := erroranimals.Errtypeboolansw(canswim); err != nil {
			fmt.Println(err)
		} else {
			Amphibian.Canswimm = canswim
		}
		firstanimal.About(Amphibian)
		firstanimal.Asking(&Amphibian)
		fmt.Println(Amphibian)
	case "рыба":
		var Fish = firstanimal.Fishes{}
		var name string
		var swim string
		var glugglug string
		var canswim bool

		fmt.Print("what is it? ")
		fmt.Fscan(os.Stdin, &name)
		Fish.Name = name
		fmt.Print("How is it swim? ")
		fmt.Fscan(os.Stdin, &swim)
		Fish.Swim = swim
		fmt.Print("How is it sound? ")
		fmt.Fscan(os.Stdin, &glugglug)
		Fish.Glugglug = glugglug
		fmt.Print("it can swim? true of false ")
		fmt.Fscan(os.Stdin, &canswim)
		if _, err := erroranimals.Errtypeboolansw(canswim); err != nil {
			fmt.Println(err)
		} else {
			Fish.Canswimm = canswim
		}
		firstanimal.About(Fish)
		firstanimal.Asking(&Fish)
		fmt.Println(Fish)
	}

}
