package main

 import (
 	"fmt"
 	weather "github.com/3crabs/go-yandex-weather-api"
 )

 func main() {
	var answer int
 	yandexWeatherApiKey := "WRITE YOUR KEY"
	 w, _ := weather.GetWeather(yandexWeatherApiKey, 55.7532, 37.622504)

	 fmt.Println("Угадайте сколько сейчас градусов")
	 fmt.Scan(&answer)

	 if w.Fact.Temp == answer {
		 fmt.Printf("Правильно, на улице " + "%d" + " градусов по цельсию", answer)
	 } else if  w.Fact.Temp != answer {
				fmt.Printf("Не верно, на улице " + "%d" + " градусов по цельсию", w.Fact.Temp)
	} else {
		fmt.Printf("На улице " + "%d" + " градусов по цельсию", w.Fact.Temp)
	}
 }