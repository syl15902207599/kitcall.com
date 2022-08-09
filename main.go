package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"kitcall.com/handler"
	"kitcall.com/util"
)

var ch = make(chan os.Signal)

func main() {

	c := context.Background()
	r := handler.MakeHttpHandlers(c)
	errChan := make(chan error)
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(*util.HPort), r)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	go func() {
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-ch)
	}()
	log.Println("server start")
	e := <-errChan
	fmt.Println(e)
}
