package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Config()

	fmt.Println("PORTA:",config.Porta)
	fmt.Println("Rodando API...")

	r := router.GerarRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",config.Porta), r))

}
