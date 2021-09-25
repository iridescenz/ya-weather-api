package main

import (
	"fmt"
	weather "github.com/3crabs/go-yandex-weather-api"
)

func main() {
	yandexWeatherApiKey := "WRITE YOUR KEY"
	w, _ := weather.GetWeather(yandexWeatherApiKey, 56.7532, 37.622504)
	cond := weatherCondition(w.Fact.Condition)
	temp := w.Fact.Temp
	fellsLike := w.Fact.FeelsLike
	wind := w.Fact.WindSpeed
	windGust := w.Fact.WindGust

	fmt.Printf("Сегодня %s, температура %d °С, ощущается как %d °С. \nВетер %d метров в секунду,"+
		" с порывами до %d метров в секунду. \n"+
		"Давление %d мм рт. ст.", cond, temp, fellsLike, wind, windGust, w.Fact.PressureMm)
}

func weatherCondition(str string) string {
	conditions := map[string]string{
		"partly-cloudy":          "малооблачно",
		"overcast":               "пасмурно",
		"drizzle":                "морось",
		"cloudy":                 "облачно с прояснениями",
		"clear":                  "ясно",
		"light-rain":             "небольшой дождь",
		"rain":                   "дождь",
		"moderate-rain":          "умеренно сильный дождь",
		"heavy-rain":             "сильный дождь",
		"continuous-heavy-rain":  "длительный сильный дождь",
		"showers":                "ливень",
		"wet-snow":               "дождь со снегом",
		"light-snow":             "небольшой снег",
		"snow":                   "снег",
		"snow-showers":           "снегопад",
		"hail":                   "град",
		"thunderstorm":           "гроза",
		"thunderstorm-with-rain": "дождь с грозой",
		"thunderstorm-with-hail": "гроза с градом",
	}

	return conditions[str]
}
