package main

import (
	"fmt"
	"github.com/sereiner/log"
	"strconv"
	"time"
)

func main() {
	v  := []byte{55,46,49,50,51}
	log.Info(string(v))
	log.Info([]byte("."))
	value, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 32)
	if err != nil {
		log.Error(err)
		log.Info(value)
	}

	time.Sleep(time.Second*1)
}
