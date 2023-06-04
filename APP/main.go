package main

import (
	"app/src/router"
	"app/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {


	utils.CarregarTemplates()
	r := router.Gerar()
	
	fmt.Println("Rodando APP na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
