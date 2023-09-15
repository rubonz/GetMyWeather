package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	fmt.Print("Enter your town: ")
	var city string
	_, err := fmt.Scanln(&city)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial&lang=en&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var weatherData map[string]interface{}
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	mainData, ok := weatherData["main"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: main data not found in response")
		return
	}

	temperature, ok := mainData["temp"].(float64)
	if !ok {
		fmt.Println("Error: temperature data not found in response")
		return
	}

	temperatureFeels, ok := mainData["feels_like"].(float64)
	if !ok {
		fmt.Println("Error: feels like temperature data not found in response")
		return
	}

	fmt.Printf("Now in the town %s %.0f°F\n", city, temperature)
	fmt.Printf("Feels like %.0f°F\n", temperatureFeels)
}
