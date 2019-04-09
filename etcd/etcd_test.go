package etcd

import (
	"testing"
	"time"

	"github.com/sereiner/log"
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

	ch, revision, _ := et.Register("/server/db", "mysql")

	data, _ := et.WatchValue("/server/db")

	go func() {
		time.Sleep(time.Second * 30)
		log.Info("删除服务")
		err := et.RevokeByPath("/server/db")
		if err != nil {
			log.Error(err)
		}
	}()

	for {
		select {
		case ka, ok := <-ch:
			if !ok || revision != ka.GetRevision() {
				return
			}
			log.Infof("%+v", ka.String())
		case d := <-data:
			dv, _ := d.GetValue()
			log.Debug(d.GetPath(), "  ", string(dv))
		}
	}

}
