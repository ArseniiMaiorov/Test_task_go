package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// структура арабских цифр в римские
	arabToRome := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	// структура римских цифр в арабские
	romeToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	// ввод числа
	fmt.Println("Введите строку для рассчета:")
	reader := bufio.NewReader(os.Stdin)
	mathString, _ := reader.ReadString('\n')
	mathString = strings.TrimSpace(mathString)
	// разделение строки на составные части
	values := strings.Split(mathString, " ")

	// проверка на корректность ввода выражения
	if len(values) != 3 {
		fmt.Println("Неверный формат ввода")
		return
	}

	// определение системы счисления
	firstVal, firstErr := strconv.Atoi(values[0])
	secondVal, secondErr := strconv.Atoi(values[2])

	// проверка на римские цифры
	firstRoman, firstRomanExists := romeToArab[values[0]]
	secondRoman, secondRomanExists := romeToArab[values[2]]

	if firstErr == nil && secondErr == nil && (0 <= firstVal && firstVal <= 10) && (0 <= secondVal && secondVal <= 10) {
		// оба числа арабские
		var result int

		// выполнение операций
		switch values[1] {
		case "+":
			result = firstVal + secondVal
		case "-":
			result = firstVal - secondVal
		case "*":
			result = firstVal * secondVal
		case "/":
			result = firstVal / secondVal
		default:
			fmt.Println("Неверный оператор")
			return
		}
		fmt.Println(printArabicResult(firstVal, secondVal, result, values[1]))
	} else if firstRomanExists && secondRomanExists {
		// оба числа римские
		var result int

		// выполнение операций
		switch values[1] {
		case "+":
			result = firstRoman + secondRoman
		case "-":
			result = firstRoman - secondRoman
		case "*":
			result = firstRoman * secondRoman
		case "/":
			result = firstRoman / secondRoman
		default:
			fmt.Println("Неверный оператор")
			return
		}

		// проверка римских цифр на корректность ввода
		if result <= 0 {
			fmt.Println("Результат меньше или равен нулю, что недопустимо для римских цифр")
		} else if result > 10 {
			fmt.Println("Результат превышает допустимый диапазон римских цифр")
		} else {
			fmt.Println(printRomeResult(values[0], values[2], arabToRome[result], values[1]))
		}
	} else {
		fmt.Println("Неверный формат ввода либо же системы счисления не совпадают или значения должны быть в диапазоне от 0 до 10")
	}
}

// функция для отрисовки вывода арабских чисел
func printArabicResult(a, b, c int, o string) string {
	return fmt.Sprintf("%d %v %d = %d", a, o, b, c)
}

// функция для отрисовки вывода римских чисел
func printRomeResult(a, b, c, o string) string {
	return fmt.Sprintf("%v %v %v = %v", a, o, b, c)
}
