package main

import "fmt"

const (
	fromEmail 		 = "matatimata01@gmail.com"
	fromName 			 = "Kauan"
	toEmail 			 = "kauanssousa@gmail.com"
	toName 				 = "Kauan"
	subject 			 = "Weather from today"
	thresholdValue = 28
	city 					 = "New Delhi"
	EmailAPIKey 	 = ""
	WeatherAPIkey  = ""
)

func getCityTemperature(city string) (int, error) {
	url := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s", WeatherAPIkey)
}