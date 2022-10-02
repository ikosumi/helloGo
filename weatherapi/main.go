package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type WeatherRes struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	Elevation            float64 `json:"elevation"`
	GenerationTimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Hourly               struct {
		Time          []string  `json:"time"`
		Temperature2M []float64 `json:"temperature_2m"`
	} `json:"hourly"`
	HourlyUnits struct {
		Temperature2M string `json:"temperature_2m"`
	} `json:"hourly_units"`
	CurrentWeather struct {
		Time          string  `json:"time"`
		Temperature   float64 `json:"temperature"`
		WeatherCode   int     `json:"weathercode"`
		WindSpeed     float64 `json:"windspeed"`
		WindDirection int     `json:"winddirection"`
	} `json:"current_weather"`
}

func main() {

	jsonFile, _ := os.Open("weather.json")
	defer jsonFile.Close()
	data, _ := io.ReadAll(jsonFile)

	var weatherRes WeatherRes
	err := json.Unmarshal(data, &weatherRes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(weatherRes.Latitude)
	fmt.Println(weatherRes.Longitude)
	fmt.Println(weatherRes.Timezone)
	fmt.Println(weatherRes.Hourly.Temperature2M)
}
