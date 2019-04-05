package main

import (
	"github.com/sereiner/lib/etcd"
	"github.com/sereiner/log"
	"time"
)

func main() {
	log.Info("开始连接....")
	server := []string{"127.0.0.1:2379"}
	et, err := etcd.NewEtcdClient(server, time.Second*5)
	defer et.Close()
	if err != nil {
		log.Error(err)
	}
	err = et.Connect()
	if err != nil {
		log.Error(err)
	}
	log.Info("连接成功")
	log.Info("开始监听")
	data, err := et.WatchChildren("/logagent")
	if err != nil {
		log.Error(err)
	}

	dataV, err := et.WatchValue("/logagent/a")
	if err != nil {
		log.Error(err)
	}

	for {
		select {
		case d := <-data:
			val, version := d.GetValue()
			log.Info(string(val), version, d.GetPath(), d.GetError())
		case dv := <-dataV:
			val, version := dv.GetValue()
			log.Info(string(val), version, dv.GetPath(), dv.GetError())
		}
	}

}
