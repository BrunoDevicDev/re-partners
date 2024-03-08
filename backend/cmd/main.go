package main

import "re/internal/http"

func main() {

	// in the main, we only start http server
	s := http.NewServer()
	s.StartServer()
}
