package main

import (
	"fmt"
)


func metersToFeet(meters float64) float64 {
	return meters * 3.28084
}

func feetToMeters(feet float64) float64 {
	return feet / 3.28084
}

func kilogramsToPounds(kg float64) float64 {
	return kg * 2.20462
}

func poundsToKilograms(pounds float64) float64 {
	return pounds / 2.20462
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var choice int
	var value float64

	fmt.Println("Select the type of conversion:")
	fmt.Println("1. Meters to Feet")
	fmt.Println("2. Feet to Meters")
	fmt.Println("3. Kilograms to Pounds")
	fmt.Println("4. Pounds to Kilograms")
	fmt.Println("5. Celsius to Fahrenheit")
	fmt.Println("6. Fahrenheit to Celsius")
	fmt.Print("Enter your choice (1-6): ")
	fmt.Scan(&choice)

	fmt.Print("Enter the value to convert: ")
	fmt.Scan(&value)

	switch choice {
	case 1:
		fmt.Printf("%.2f meters = %.2f feet\n", value, metersToFeet(value))
	case 2:
		fmt.Printf("%.2f feet = %.2f meters\n", value, feetToMeters(value))
	case 3:
		fmt.Printf("%.2f kilograms = %.2f pounds\n", value, kilogramsToPounds(value))
	case 4:
		fmt.Printf("%.2f pounds = %.2f kilograms\n", value, poundsToKilograms(value))
	case 5:
		fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", value, celsiusToFahrenheit(value))
	case 6:
		fmt.Printf("%.2f Fahrenheit = %.2f Celsius\n", value, fahrenheitToCelsius(value))
	default:
		fmt.Println("Invalid choice")
	}
}
