package main

import (
	"fmt"
	"math"
	"strings"
)

const IMTPower = 2

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	for {
		userKg, userHeight := getUserInput()
		IMT := calculateIMT(userKg, userHeight)
		outputResult(IMT)
		if !checkRepeatCalculation() {
			break
		}
	}
}

func outputResult(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыток массы тела")
	default:
		fmt.Println("У вас степень ожирения")
	}
	fmt.Printf("Ваш индекс массы тела: %.0f\n", IMT)
}

func calculateIMT(kg float64, height float64) float64 {
	return kg * 10000 / math.Pow(height, IMTPower)
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Println("Введите ваш рост в сантиметрах:")
	fmt.Scan(&userHeight)
	fmt.Println("Введлите ваш вес в килограммах:")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Println("Вы хотите повторить расчет? (Y/N)")
	fmt.Scan(&userChoise)
	return strings.ToUpper(userChoise) == "Y"
}
