package main

import (
	"go_restfulapi/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
