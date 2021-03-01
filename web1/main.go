package main

import (
	"net/http"
	"tucker/web1/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
