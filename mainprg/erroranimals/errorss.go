package erroranimals

import "errors"

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
