package main

import (
	"fmt"
	"time"
)

func main() {
	tt, _ := time.Parse(time.DateTime, "2023-12-31 14:24:21")
	fmt.Println(tt)
	fmt.Println(tt.Format(time.DateTime))
}
