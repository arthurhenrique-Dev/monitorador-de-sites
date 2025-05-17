package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var sites []string

func main() {

	introducao()

	var timeIndex int = 0

	InputSite(&timeIndex)

	for {

		TimerMonitoramento(timeIndex)
		InputSite(&timeIndex)

	}

}

func introducao() {
	fmt.Println("DESEJA AVERIGUAR SE O SEU SITE ESTÁ NO AR?")
	fmt.Println("------------------------------------------")
}

func Monitoramento(sites []string) {
	fmt.Println("monitorando......")
	for i, site := range sites {
		resp, _ := http.Get(site)
		var iTradutor = i + 1
		if resp.StatusCode == 200 {
			fmt.Println(time.Now().Format("02/01/2006 15:04:05"), iTradutor, "-", site, "| online")
		} else {
			fmt.Println(time.Now().Format("02/01/2006 15:04:05"), iTradutor, "-", site, "| offline")
		}
	}
}
func TimerMonitoramento(timeIndex int) {
	if timeIndex != 0 {
		for i := 0; i < timeIndex; i++ {
			time.Sleep(5 * time.Second)
			Monitoramento(sites)
		}
	} else {
		Monitoramento(sites)
	}
}

func InputSite(timeIndex *int) {
	site := textSite()
	sites = append(sites, site)

	fmt.Println("Deseja verificar mais algum site? [s/n]")
	var confirm string
	fmt.Scan(&confirm)
	if confirm == "s" {
		InputSite(timeIndex)
	} else if confirm == "n" {
		fmt.Println("Quer monitorar por um tempo? [s/n]")
		var confirmTime string
		fmt.Scan(&confirmTime)
		if confirmTime == "s" {
			timeIndex := TimeCalculation()
			TimerMonitoramento(timeIndex)
			saida()

		} else if confirmTime == "n" {
			saida()
		}
	} else {
		fmt.Println("resposta inválida, escreva literalmente como eu escrevi")
		InputSite(timeIndex)
	}
}
func textSite() string {
	fmt.Println("copie e cole o link:")

	var site string
	fmt.Scan(&site)
	return site
}
func TimeCalculation() int {
	fmt.Println("Por quanto tempo em min?")
	var TimeInput int
	fmt.Scan(&TimeInput)
	var TimeIndex = TimeInput * 12
	return TimeIndex
}
func saida() {
	fmt.Println("muito obrigado por usar este programa")
	os.Exit(0)
}
