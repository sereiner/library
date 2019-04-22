package elastic

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"github.com/sereiner/lib/types"
	"reflect"
)

// ElasticSearch es组件
type ElasticSearch struct {
	host  []string
	conn  *elastic.Client
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
func New(conf ESConfigOption) (es *ElasticSearch, err error) {
	es = &ElasticSearch{}
	es.host = conf.Host
	es.Index = conf.Index
	es.Type = conf.Type
	es.conn, err = elastic.NewClient(elastic.SetErrorLog(conf.log), elastic.SetURL(es.host...), elastic.SetSniff(conf.Sniff))
	if err != nil {
		return nil, err
	}
	_, _, err = es.conn.Ping(es.host[0]).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return
}

// Create 创建一条记录,没有索引会同时创建索引
func (es *ElasticSearch) Create(id string,data map[string]interface{}) (err error) {

	_, err = es.conn.Index().
		Index(es.Index).
		Type(es.Type).
		Id(id).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		return
	}
	return
}

// Delete 根据id删除一条记录
func (es *ElasticSearch) Delete(id string) (err error) {

	_, err = es.conn.Delete().
		Index(es.Index).
		Type(es.Type).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}

	return
}

// Update 修改一条记录,字段可以不完整
func (es *ElasticSearch) Update(id string, doc map[string]interface{}) (err error) {

	_, err = es.conn.Update().
		Index(es.Index).
		Type(es.Type).
		Id(id).
		Doc(doc).
		Do(context.Background())
	if err != nil {
		return
	}

	return
}

// Gets 通过id查找记录,没有记录返回错误
func (es *ElasticSearch) Gets(id string) (res []byte, err error) {

	get, err := es.conn.Get().
		Index(es.Index).
		Type(es.Type).
		Id(id).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	if get.Found {
		return get.Source.MarshalJSON()
	}

	return nil, errors.New("not found")
}

// List 分页获取内容
// tag 字段名称
// query 要查的内容
// size 每页显示的条数
// page 页码
func (es *ElasticSearch) List(tag, query string, size, page int) (data []interface{},total int64, err error) {

	if size < 0 || page < 1 {
		return nil,0, errors.New("param error")
	}

	q := elastic.NewMatchPhraseQuery(tag, query)

	res, err := es.conn.Search(es.Index).
		Type(es.Type).
		Query(q).
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	if err != nil {
		return nil,0, err
	}

	typ := map[string]interface{}{}
	return res.Each(reflect.TypeOf(typ)), res.TotalHits(), nil
}

// Bulk 批量导入,修改记录.字段必须完整,否则记录可能丢失字段
func (es *ElasticSearch) Bulk(array []map[string]interface{}) (res []*elastic.BulkResponseItem, err error) {

	bulkRequest := es.conn.Bulk()
	for _, v := range array {
		if types.GetString(v["id"]) == "" {
			return nil, errors.New("缺少id")
		}
		req := elastic.NewBulkIndexRequest().
			Index(es.Index).
			Type(es.Type).
			Id(types.GetString(v["id"])).
			Doc(v)
		bulkRequest.Add(req)
	}

	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		return nil, err
	}

	return bulkResponse.Indexed(), nil
}
