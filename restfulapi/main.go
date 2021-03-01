package main

import (
	"net/http"
	"tucker/restfulapi/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
