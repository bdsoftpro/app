package main

import (
	"fmt"
	"net/http"
	"github.com/bdsoftpro/app/server"
	"os"
)	
	
func main() {
	go func() {
		port := os.Getenv("PORT")
		fmt.Println(fmt.Sprintf(":%s", port))
		http.ListenAndServe(fmt.Sprintf(":%s", port), http.HandlerFunc(routers.Serve))
	}()
	for {}
}
