package erroranimals

import "errors"
import "strconv"
import "fmt"

func ErrTypeAnimal(answer string) (string, error) {
	if answer != "насекомое" && answer != "птица" && answer != "дикое" && answer != "амфибия" && answer != "рыба" {
		return "false", errors.New("вы ввели неправильный тип животного")
	}
	return answer, nil
}
func Errtypeboolansw(answer bool) (bool, error) {
	if answer != true && answer != false {
		return false, errors.New("вы ввели неправильный ответ на вопрос")
	}
	return answer, nil
}

func ErrorWeight(answer string) (string, error) {
	if answer != "больших" && answer != "маленьких" {
		return "указан не верный размер животного", errors.New("вы ввели неправильный ответ на вопрос")
	}
	return answer, nil
}
func ErrorLiving(answer string) (string, error) {
	if answer != "ocean" && answer != "desert" && answer != "forest" && answer != "rainforest" && answer != "savannah" && answer != "jungle" && answer != "domestic" {
		return "указан не верный ореал обитания", errors.New("вы ввели неправильный ответ на вопрос")
	}
	return answer, nil
}

func ErrorAge(answer string) (string, error) {
	floatNum, err := strconv.ParseFloat(answer, 64)
	if err == nil {
		fmt.Println("неправильный тип данных", floatNum)
	} else {
		num, err := strconv.Atoi(answer)
		if err != nil {
			return "вы ввели не число", errors.New("вы ввели неправильный ответ на вопрос")
		} else {
			fmt.Println("выбранный возраст:", num)
		}
	}
	return answer, nil
}
