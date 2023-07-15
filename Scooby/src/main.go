package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var SCOOBYCALLSCOUNT int

const serverPort = 3333

func main() {

	log.SetPrefix("Scooby: ")
	log.Println("searching SCOOBYCALLSCOUNT in system variables")
	systemVar := os.Getenv("SCOOBYCALLSCOUNT")
	if systemVar != "" {
		log.Println("Found systemVar: ", systemVar)
		callsCount, err := strconv.Atoi(systemVar)
		if err != nil {
			log.Println(err)
		} else {
			SCOOBYCALLSCOUNT = callsCount
		}

	} else {
		log.Println("No SCOOBYCALLSCOUNT found in system variable, searching local config file...")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error reading config file: %s\n", err)
		}

		log.Printf("Found %s yaml file", "config.yaml")
		//myVariable := viper.GetString("SCOOBYCALLSCOUNT")
		myVariable := viper.GetString("SCOOBYCALLSCOUNT")
		if myVariable == "" {
			log.Fatal("no SCOOBYCALLSCOUNT in local config file")
		}
		log.Println("SCOOBYCALLSCOUNT: ", myVariable)
		configValue, err := strconv.Atoi(myVariable)
		if err != nil {
			log.Fatal(err)
		} else {
			SCOOBYCALLSCOUNT = configValue
		}

	}

	log.Printf("Calling velma service %d times", SCOOBYCALLSCOUNT)
	callVelma(SCOOBYCALLSCOUNT)

}

func callVelma(callsCount int) {
	for i := 0; i < callsCount; i++ {
		//resp, err := http.Get("http://172.17.0.2:10000/sayHello/Alex")
		resp, err := http.Get("http://velma:10000/sayHello/Alex")

		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sb := string(body)
		log.Printf(sb)
	}

}
