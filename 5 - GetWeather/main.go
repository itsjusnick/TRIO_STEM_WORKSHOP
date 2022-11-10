package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var city string

	fmt.Println("Please type the name of the city")
	fmt.Scanln(&city)

	weather := getWeather(city)

	fmt.Println(weather)

}

func getWeather(city string) string {
	key := "6bc3991461917fc758882a20ac2b337d" //this key is used to help the API know who I am and allows me to use the service

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + key + "&units=imperial") // http.Get is how we make the call to the API to ask for the weather

	if err != nil {
		log.Fatalln(err)
	}

	return printResponse(resp)

}

// You can ignore this function I made. It's here to decode the response into something we can read.
func printResponse(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
