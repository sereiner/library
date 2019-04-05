package etcd

import (
	"testing"
	"time"
)

func TestNewEtcdClient(t *testing.T) {
	server := []string{"127.0.0.1:2379"}
	et, err := NewEtcdClient(server, time.Second*2)
	if err != nil {
		t.Fatal(err)
	}
	err = et.Connect()
	if err != nil {
		t.Fatal(err)
	}
	err = et.Update("/logagent/server","123")
	if err != nil {
		t.Log(err)
	}
	err = et.Close()
	if err != nil {
		t.Fatal(err)
	}
}
