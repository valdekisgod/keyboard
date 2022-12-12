//Package keyboard Пакет для считывания ввода пользователя с клавиатуры (ну типа)
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//GetFloat Считывает число с плавающей точкой с клавиатуры
//Возвращает считываемое число или любую возникшую ошибку
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}
