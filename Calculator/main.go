package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Маппинг римских чисел в арабские
var romanToArabic = map[string]int{
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

// Маппинг арабских чисел в римские
var arabicToRoman = map[int]string{
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

// convertToArabic преобразует строку с числом в арабское число
func convertToArabic(input string) (int, error) {
	if value, ok := romanToArabic[input]; ok {
		return value, nil
	}
	return strconv.Atoi(input)
}

// convertToRoman преобразует арабское число в римское число
func convertToRoman(num int) (string, error) {
	if num < 1 {
		return "", fmt.Errorf("римские числа не могут быть меньше I")
	}
	if value, ok := arabicToRoman[num]; ok {
		return value, nil
	}

	// Построение римских чисел больших, чем 10
	roman := ""
	for num >= 10 {
		roman += "X"
		num -= 10
	}
	if num > 0 {
		roman += arabicToRoman[num]
	}
	return roman, nil
}

// calculate выполняет арифметическую операцию над двумя числами
func calculate(operand1, operand2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("деление на ноль невозможно")
		}
		return operand1 / operand2, nil
	default:
		return 0, fmt.Errorf("неизвестный оператор: %s", operator)
	}
}

// isRoman проверяет, является ли входная строка римским числом
func isRoman(input string) bool {
	_, isRoman := romanToArabic[input]
	return isRoman
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Калькулятор")
	fmt.Println("Введите операцию в формате: число1 оператор число2")
	fmt.Println("Пример: 5 + 3 или VII * II")
	fmt.Println("Для выхода введите 'exit'.")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic("Ошибка чтения ввода")
		}

		// Удаляем символы перевода строки и пробелы
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Выход из программы.")
			break
		}

		// Разбиваем строку на части
		parts := strings.Fields(input)
		if len(parts) != 3 {
			panic("Неправильный формат ввода. Ожидается: число1 оператор число2")
		}

		num1Str, operator, num2Str := parts[0], parts[1], parts[2]

		// Определяем, какие числа у нас - арабские или римские
		isNum1Roman, isNum2Roman := isRoman(num1Str), isRoman(num2Str)

		var num1, num2 int

		if isNum1Roman {
			// Парсим римские числа
			num1, err = convertToArabic(num1Str)
			if err != nil || num1 < 1 || num1 > 10 {
				panic("Некорректное римское число. Ожидается число от I до X")
			}

			num2, err = convertToArabic(num2Str)
			if err != nil || num2 < 1 || num2 > 10 {
				panic("Некорректное римское число. Ожидается число от I до X")
			}
		} else {
			// Парсим арабские числа
			num1, err = strconv.Atoi(num1Str)
			if err != nil || num1 < 1 || num1 > 10 {
				panic("Некорректное арабское число. Ожидается число от 1 до 10")
			}

			num2, err = strconv.Atoi(num2Str)
			if err != nil || num2 < 1 || num2 > 10 {
				panic("Некорректное арабское число. Ожидается число от 1 до 10")
			}
		}

		if isNum1Roman != isNum2Roman {
			panic("Калькулятор не может работать с римскими и арабскими числами одновременно")
		}

		// Выполняем расчет
		result, err := calculate(num1, num2, operator)
		if err != nil {
			panic(err)
		}

		// Выводим результат
		if isNum1Roman {
			romanResult, err := convertToRoman(result)
			if err != nil {
				panic(err)
			}
			fmt.Println("Результат:", romanResult)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}
