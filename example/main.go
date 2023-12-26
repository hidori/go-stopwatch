package main

import (
	"fmt"
	"time"

	"github.com/hidori/go-stopwatch"
)

func main() {
	sw := stopwatch.NewStopwatch()
	fmt.Println(time.Now())

	sw.Start()
	time.Sleep(1 * time.Second)
	fmt.Println(sw.Duration())
	time.Sleep(1 * time.Second)
	fmt.Println(sw.Duration())

	sw.Stop()
	time.Sleep(1 * time.Second)
	fmt.Println(sw.Duration())
	time.Sleep(1 * time.Second)
	fmt.Println(sw.Duration())

	sw.Start()
	time.Sleep(1 * time.Second)
	fmt.Println(sw.Duration())
	time.Sleep(1 * time.Second)
	fmt.Println(sw.Duration())
}
