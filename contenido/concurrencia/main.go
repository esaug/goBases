package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
	}

	ch := make(chan string)

	fmt.Print("First")
	for _, api := range apis {
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapse := time.Since(start)
	fmt.Printf("Listo Tomo %v segundos! \n ", elapse.Seconds())
}

func checkApi(api string, ch chan string) {

	_, err := http.Get(api)

	if err != nil {
		ch <- fmt.Sprintf("ERRO: Api %s caida!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCESS: Api %s en funcionamiento\n", api)
}
