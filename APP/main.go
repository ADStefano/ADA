package main

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/router"
	"app/utils"
//	"encoding/hex"
	"fmt"
	"log"
	"net/http"

//	"github.com/gorilla/securecookie"
)

// Função para criar uma hash key e um block key aleatório
// func init(){
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)
// 	fmt.Println(blockKey)
// }

func main() {

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()
	
	fmt.Printf("PORTA: %d\nRodando APP...\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
