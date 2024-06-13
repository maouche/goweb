package main

import (
	"fmt"
	"net/http"

	"github.com/maouche/goweb/internal/handlers"
)

const port = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)

	fmt.Println("(http://localhost:"+port+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
