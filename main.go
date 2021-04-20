package main

import (
	"fmt"
	"net/http"

	"github.com/slofurno/greetings/service"
	"github.com/slofurno/greetings/service/greetingspb"
)

func main() {
	svc := greetings.New()
	server := greetingspb.NewGreeterServer(svc)
	fmt.Println(http.ListenAndServe(":8083", server))
}
