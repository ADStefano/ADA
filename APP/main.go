package main

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/router"
	"app/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()
	
	fmt.Printf("Rodando APP na porta: %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
