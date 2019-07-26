package elastic

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/olivere/elastic"
)

type SpuItem struct {
	Author           string             `json:"author"`
	CategoryID       int                `json:"category_id"`
	DescriptionUnits []*DescriptionUnit `json:"description_units"`
	ID               string             `json:"id"`
	Isbn             string             `json:"isbn"`
	Status           int                `json:"status"`
	Stock            int                `json:"stock"`
	SubTitle         string             `json:"sub_title"`
	Tags             []string           `json:"tags"`
	Title            string             `json:"title"`
	Image            string             `json:"image"`
	OriginalPrice    string             `json:"original_price"`
	SellingPrice     string             `json:"selling_price"`
}

type DescriptionUnit struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func TestNew(t *testing.T) {
	es, err := New(ESConfigOption{
		Host:  []string{"http://39.98.58.40:9200/"},
		Index: "spu",
		Type:  "spu",
		Sniff: false,
	})
	if err != nil {
		t.Log(err)
	}

	keyword := "中国文学"

	boolQuery := elastic.NewBoolQuery()
	// boolQuery.Must(elastic.NewMatchQuery("status", 1))
	boolQuery.Should(elastic.NewMatchPhraseQuery("tags", keyword))
	fmt.Println(boolQuery.Source())
	datas, err := es.Conn.Search().
		Index("spu").
		Type("spu").
		Query(boolQuery).
		Do(context.Background())
	if err != nil {
		return
	}

	fmt.Println(datas.Profile)

	var typ SpuItem
	for k, item := range datas.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(SpuItem)
		fmt.Printf("key:%v,%#v\n", k, t)
	}
}
