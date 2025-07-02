package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gerrod.com/di"
	"gerrod.com/hello"
	"gerrod.com/mocking"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	fmt.Println(hello.Hello("world", ""))
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
