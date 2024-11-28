package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pilih kalkulator yang ingin digunakan:")
	fmt.Println("1. Celsius ke Fahrenheit")
	fmt.Println("2. Celsius ke Kelvin")
	fmt.Println("3. Fahrenheit ke Celsius")
	fmt.Println("4. Fahrenheit ke Kelvin")
	fmt.Println("5. Kelvin ke Celsius")
	fmt.Println("6. Kelvin ke Fahrenheit")
	
	var pilihan int
	fmt.Print("Masukkan pilihan Anda (1-6): ")
	for {
		fmt.Scan(&pilihan)
		if pilihan >= 1 && pilihan <= 6 {
			break
		}
		fmt.Print("Pilihan tidak valid, masukkan angka antara 1-6: ")
	}

	var suhu float64
	fmt.Print("Masukkan suhu yang ingin dikonversi: ")
	for {
		if _, err := fmt.Scan(&suhu); err == nil {
			break
		}
		fmt.Println("Input suhu tidak valid, coba lagi: ")
		fmt.Scanln() 
	}

	var suhuAkhir float64

	switch pilihan {
	case 1:
		suhuAkhir = CelsiusKeFahrenheit(suhu)
	case 2:
		suhuAkhir = CelsiusKeKelvin(suhu)
	case 3:
		suhuAkhir = FahrenheitKeCelsius(suhu)
	case 4:
		suhuAkhir = FahrenheitKeKelvin(suhu)
	case 5:
		suhuAkhir = KelvinKeCelsius(suhu)
	case 6:
		suhuAkhir = KelvinKeFahrenheit(suhu)
	}

	fmt.Printf("Suhu hasil konversi adalah: %.2f\n", suhuAkhir)
}

func CelsiusKeFahrenheit(suhuCelsius float64) float64 {
	return (9.0/5.0)*suhuCelsius + 32
}

func CelsiusKeKelvin(suhuCelsius float64) float64 {
	return suhuCelsius + 273.15
}

func FahrenheitKeCelsius(suhuFahrenheit float64) float64 {
	return (suhuFahrenheit - 32) * (5.0 / 9.0)
}

func FahrenheitKeKelvin(suhuFahrenheit float64) float64 {
	return (suhuFahrenheit + 459.67) * (5.0 / 9.0)
}

func KelvinKeCelsius(suhuKelvin float64) float64 {
	return suhuKelvin - 273.15
}

func KelvinKeFahrenheit(suhuKelvin float64) float64 {
	return (suhuKelvin * 9.0 / 5.0) - 459.67
}
