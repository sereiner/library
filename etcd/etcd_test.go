package etcd

import (
	"github.com/sereiner/log"
	"testing"
	"time"
)

func TestNewEtcdClient(t *testing.T) {
	server := []string{"127.0.0.1:2379"}
	et, err := NewEtcdClient(server, time.Second*2)
	if err != nil {
		t.Fatal(err)
	}

	defer et.Close()

	err = et.Connect()
	if err != nil {
		t.Fatal(err)
	}
	err = et.Update("/logagent/server", "123")
	if err != nil {
		t.Log(err)
	}

	ch, _ := et.Register("/server/db", "mysql")

	data, _ := et.WatchValue("/server/db")

	go func() {
		time.Sleep(time.Second *15)
		log.Info("删除服务")
		et.RevokeByPath("/server/db")
	}()

	for {
		select {
		case ka, ok := <-ch:
			if !ok {
				return
			} else {
				log.Info(ka.ID)
			}
		case d := <-data:
			dv , _ := d.GetValue()
			log.Debug(d.GetPath(),"  ",dv)
		}
	}

}
