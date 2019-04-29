package elastic

import (
	"testing"
)

func TestNew(t *testing.T) {
	conf := ESConfigOption{
		Host:  []string{"http://localhost:9200/"},
		Index: "tag_test",
		Type:  "book",
		Sniff: false,
	}
	es, err := New(conf)
	if err != nil {
		t.Error(err)
	}
	res, err := es.Gets("1")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
}

func TestElasticSearch_Gets(t *testing.T) {
	conf := ESConfigOption{
		Host:  []string{"http://localhost:9200/"},
		Index: "tag_test",
		Type:  "book",
		Sniff: false,
	}
	es, err := New(conf)
	if err != nil {
		t.Error(err)
	}
	res, err := es.Gets("1037240")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
}

func TestElasticSearch_List(t *testing.T) {
	conf := ESConfigOption{
		Host:  []string{"http://localhost:9200/"},
		Index: "tag_test",
		Type:  "book",
		Sniff: false,
	}
	es, err := New(conf)
	if err != nil {
		t.Error(err)
	}
	res, total, err := es.List("spu_id", "9787111432326", 5, 1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", res)
	t.Log("total: ", total)
}

func TestElasticSearch_Update(t *testing.T) {
	conf := ESConfigOption{
		Host:  []string{"http://calorie-qa.manyoujing.net:9200/"},
		Index: "tag_test",
		Type:  "book",
		Sniff: false,
	}
	es, err := New(conf)
	if err != nil {
		t.Error(err)
	}
	err = es.Update("2", map[string]interface{}{
		"tag_name": "送你一匹马",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestElasticSearch_Create(t *testing.T) {
	conf := ESConfigOption{
		Host:  []string{"http://calorie-qa.manyoujing.net:9200/"},
		Index: "tag_test",
		Type:  "book",
		Sniff: false,
	}
	es, err := New(conf)
	if err != nil {
		t.Error(err)
	}
	err = es.Create("4", map[string]interface{}{
		"id":       4,
		"spu_id":   "11223345",
		"tag_id":   1234,
		"tag_name": "雨季不再来4",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestElasticSearch_Delete(t *testing.T) {
	conf := ESConfigOption{
		Host:  []string{"http://39.98.58.40:9200/"},
		Index: "tag",
		Type:  "book",
		Sniff: false,
	}
	es, err := New(conf)
	if err != nil {
		t.Error(err)
	}
	err = es.Delete("2091727")
	if err != nil {
		t.Error(err)
	}
}

func TestElasticSearch_Bulk(t *testing.T) {
	conf := ESConfigOption{
		Host:  []string{"http://localhost:9200/"},
		Index: "tag_test",
		Type:  "book",
		Sniff: false,
	}
	es, err := New(conf)
	if err != nil {
		t.Error(err)
		return
	}
	arr := []map[string]interface{}{}

	arr = append(arr, map[string]interface{}{
		"id":       1,
		"spu_id":   "11223345",
		"tag_id":   1234,
		"tag_name": "雨季不再来4",
	}, map[string]interface{}{
		"id":       2,
		"spu_id":   "11223345",
		"tag_id":   1234,
		"tag_name": "雨季不再来4",
	})

	res, err := es.Bulk(arr)
	if err != nil {
		t.Error(err)
	}
	for _, v := range res {
		t.Logf("%+v", v)
	}
}

func TestElasticSearch_Group(t *testing.T) {

	conf := ESConfigOption{
		Host:  []string{"http://39.98.58.40:9200/"},
		Index: "tag",
		Type:  "book",
		Sniff: false,
	}
	es, err := New(conf)
	if err != nil {
		t.Error(err)
	}
	res, total, err := es.Group("tag_name", "漫画", "tag_id", 20, 1)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", res)
	t.Log(total)

}
