package balancer

import (
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	ip := "127.0.0.1"
	port := "9999"

	err := Register("http://127.0.0.1:2379", "receipt","receipt", ip, port, time.Second*10, 15)
	if err != nil {
		t.Error(err)
	}

	time.Sleep(60*time.Second)

}

func TestUnRegister(t *testing.T) {
	UnRegister()
}