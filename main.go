package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mo7rex/mongo_go_api/router"
)

func main() {
	fmt.Println("START!")
	r := router.Router()
	fmt.Println("The server running at port 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", r))

}
