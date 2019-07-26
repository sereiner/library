package elastic

import (
	"context"

	"github.com/olivere/elastic"
)

// Search es组件
type Search struct {
	Host  []string
	Conn  *elastic.Client
	Index string
	Type  string
}

//ESConfigOption 配置文件
type ESConfigOption struct {
	Host  []string `json:"hosts"`
	Index string
	Type  string
	Sniff bool
	Log   elastic.Logger
}

// New 创建elastic 实例
func New(conf ESConfigOption) (es *Search, err error) {
	es = &Search{}
	es.Host = conf.Host
	es.Index = conf.Index
	es.Type = conf.Type
	es.Conn, err = elastic.NewClient(elastic.SetErrorLog(conf.Log), elastic.SetURL(es.Host...), elastic.SetSniff(conf.Sniff))
	if err != nil {
		return nil, err
	}
	_, _, err = es.Conn.Ping(es.Host[0]).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return
}
