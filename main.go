package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println(time.Since(time.Date(1970,1,1,0,0,0,0, time.UTC)).Milliseconds())
}
