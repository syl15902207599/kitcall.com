package main

import (
	"context"
	"net/http"

	"kitcall.com/handler"
)

func main() {
	c := context.Background()
	r := handler.MakeHttpHandlers(c)
	http.ListenAndServe(":3001", r)
}
