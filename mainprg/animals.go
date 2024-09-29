package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"nadonado/mainprg/erroranimals"
	"nadonado/mainprg/firstanimal"
	"os"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@/animalsdb")
	if err != nil {
		panic(err)
	} else {
		fmt.Print("good")
	}
	defer db.Close()

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

		result, err := db.Exec("insert into animalsdb.insects (Name, Crawl, Irratation, Canswimm, Wings) values (?, ?, ?, ?, ?)",
			name, crawl, irritation, canswim, wings)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
		rows, err := db.Query("select * from animalsdb.insects")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		Insecttwo := []firstanimal.Insects{}
		for rows.Next() {
			insec := firstanimal.Insects{}
			err := rows.Scan(&insec.Id_in, &insec.Name, &insec.Crawl, &insec.Irritation, &insec.Canswimm, &insec.Wings)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Insecttwo = append(Insecttwo, insec)
		}
		for _, ins := range Insecttwo {
			fmt.Println(ins.Name, ins.Crawl, ins.Irritation, ins.Canswimm, ins.Wings)
		}

		var answer2 string
		var answerid string
		fmt.Print("хотите удалить данные? y/n")
		fmt.Fscan(os.Stdin, &answer2)
		if answer2 == "y" {
			fmt.Print("какой id? ")
			fmt.Fscan(os.Stdin, &answerid)
			result, err = db.Exec("DELETE FROM animalsdb.insects WHERE id = ?", answerid)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.LastInsertId())
			fmt.Println(result.RowsAffected())
		}

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

		result, err := db.Exec("insert into animalsdb.Birds (Name, Fly, Song, Canswimm, Wings) values (?, ?, ?, ?, ?)",
			name, fly, song, canswim, wings)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
		rows, err := db.Query("select * from animalsdb.Birds")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		Birdtwo := []firstanimal.Birds{}
		for rows.Next() {
			bird := firstanimal.Birds{}
			err := rows.Scan(&bird.Id_b, &bird.Name, &bird.Fly, &bird.Song, &bird.Canswimm, &bird.Wings)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Birdtwo = append(Birdtwo, bird)
		}
		for _, ins := range Birdtwo {
			fmt.Println(ins.Name, ins.Fly, ins.Song, ins.Canswimm, ins.Wings)
		}

		var answer2 string
		var answerid string
		fmt.Print("хотите удалить данные? y/n")
		fmt.Fscan(os.Stdin, &answer2)
		if answer2 == "y" {
			fmt.Print("какой id? ")
			fmt.Fscan(os.Stdin, &answerid)
			result, err = db.Exec("DELETE FROM animalsdb.Birds WHERE id = ?", answerid)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.LastInsertId())
			fmt.Println(result.RowsAffected())
		}
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
		result, err := db.Exec("insert into animalsdb.wild_aninals (Name, Run, Frighteningly, Canswimm) values ( ?, ?, ?, ?)",
			name, run, frighteningly, canswim)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
		rows, err := db.Query("select * from animalsdb.wild_aninals")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		Animaltwo := []firstanimal.Animals{}
		for rows.Next() {
			an := firstanimal.Animals{}
			err := rows.Scan(&an.Id_an, &an.Name, &an.Run, &an.Frighteningly, &an.Canswimm)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Animaltwo = append(Animaltwo, an)
		}
		for _, ins := range Animaltwo {
			fmt.Println(ins.Name, ins.Run, ins.Frighteningly, ins.Canswimm)
		}
		var answer2 string
		var answerid string
		fmt.Print("хотите удалить данные? y/n")
		fmt.Fscan(os.Stdin, &answer2)
		if answer2 == "y" {
			fmt.Print("какой id? ")
			fmt.Fscan(os.Stdin, &answerid)
			result, err = db.Exec("DELETE FROM animalsdb.wild_aninals WHERE id = ?", answerid)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.LastInsertId())
			fmt.Println(result.RowsAffected())
		}
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

		result, err := db.Exec("insert into animalsdb.Amphibians (Name, Strange, Runandswim, Canswimm) values ( ?, ?, ?, ?)",
			name, strange, runandswim, canswim)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
		rows, err := db.Query("select * from animalsdb.Amphibians ")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		Amphibiantwo := []firstanimal.Amphibians{}
		for rows.Next() {
			am := firstanimal.Amphibians{}
			err := rows.Scan(&am.Id_am, &am.Name, &am.Strange, &am.Runandswim, &am.Canswimm)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Amphibiantwo = append(Amphibiantwo, am)
		}
		for _, ins := range Amphibiantwo {
			fmt.Println(ins.Name, ins.Strange, ins.Runandswim, ins.Canswimm)
		}
		var answer2 string
		var answerid string
		fmt.Print("хотите удалить данные? y/n")
		fmt.Fscan(os.Stdin, &answer2)
		if answer2 == "y" {
			fmt.Print("какой id? ")
			fmt.Fscan(os.Stdin, &answerid)
			result, err = db.Exec("DELETE FROM animalsdb.Amphibians WHERE id = ?", answerid)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.LastInsertId())
			fmt.Println(result.RowsAffected())
		}
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

		result, err := db.Exec("insert into animalsdb.fishes (Name, Swim , Glugglug , Canswimm) values ( ?, ?, ?, ?)",
			name, swim, glugglug, canswim)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
		rows, err := db.Query("select * from animalsdb.fishes ")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		Fishtwo := []firstanimal.Fishes{}
		for rows.Next() {
			fsh := firstanimal.Fishes{}
			err := rows.Scan(&fsh.Id_f, &fsh.Name, &fsh.Swim, &fsh.Glugglug, &fsh.Canswimm)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Fishtwo = append(Fishtwo, fsh)
		}
		for _, ins := range Fishtwo {
			fmt.Println(ins.Name, ins.Swim, ins.Glugglug, ins.Canswimm)
		}
		var answer2 string
		var answerid string
		fmt.Print("хотите удалить данные? y/n")
		fmt.Fscan(os.Stdin, &answer2)
		if answer2 == "y" {
			fmt.Print("какой id? ")
			fmt.Fscan(os.Stdin, &answerid)
			result, err = db.Exec("DELETE FROM animalsdb.fishes WHERE id = ?", answerid)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.LastInsertId())
			fmt.Println(result.RowsAffected())
		}
	}

}
