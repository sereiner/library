package elastic

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// Request 请求数据
type Request struct {
	Query    QueryType              `json:"query"`
	Collapse map[string]interface{} `json:"collapse"`
	From     int                    `json:"from,omitempty"`
	Size     int                    `json:"size,omitempty"`
	Sort     []map[string]string    `json:"sort"`
}

// QueryType 请求类型
type QueryType struct {
	Bool BoolType `json:"bool"`
}

// BoolType BoolType类型
type BoolType struct {
	Must    []map[string]interface{} `json:"must,omitempty"`
	MustNot []map[string]interface{} `json:"must_not,omitempty"`
}

// NewRequestModel 构建请求数据
func NewRequestModel(from, size int) *Request {
	return &Request{
		Query: QueryType{
			Bool: BoolType{
				Must: []map[string]interface{}{},
			},
		},
		Collapse: map[string]interface{}{},
		From:     from,
		Size:     size,
		Sort:     []map[string]string{},
	}
}

func (l *Request) AddCollapse(field string) {
	if _, ok := l.Collapse["field"]; ok {
		return
	}
	l.Collapse["field"] = field
}

// AddCondition 添加条件
func (l *Request) AddCondition(fields []string, query string) {
	// 处理字符串
	query = strings.Trim(query, " ")
	reg, _ := regexp.Compile(" +")
	query = reg.ReplaceAllString(query, " ")

	// 筛选must和must_not
	str := strings.Split(query, " ")
	var mustStr []string
	var mustNotStr []string
	for _, item := range str {
		if strings.HasPrefix(item, "!") && len(item) > 1 {
			mustNotStr = append(mustNotStr, item[1:])
		} else {
			mustStr = append(mustStr, item)
		}
	}

	// 添加筛选条件
	if len(mustStr) > 0 {
		l.addMustCondition(fields, mustStr)
	}

	if len(mustNotStr) > 0 {
		l.addMustNotCondition(fields, mustNotStr)
	}
}

func (l *Request) addMustCondition(fields []string, query []string) {
	for _, item := range query {
		l.Query.Bool.Must = append(l.Query.Bool.Must, map[string]interface{}{
			"query_string": map[string]interface{}{
				"fields": fields,
				"query":  l.transQuery(item),
			},
		})
	}
}

func (l *Request) addMustNotCondition(fields []string, query []string) {
	for _, item := range query {
		l.Query.Bool.MustNot = append(l.Query.Bool.MustNot, map[string]interface{}{
			"query_string": map[string]interface{}{
				"fields": fields,
				"query":  l.transQuery(item),
			},
		})
	}
}

// AddExactMatch 添加精准匹配
func (l *Request) AddExactMatch(fields string, query string) {
	l.Query.Bool.Must = append(l.Query.Bool.Must, map[string]interface{}{
		"match_phrase": map[string]interface{}{
			fields: map[string]interface{}{
				"query": fmt.Sprintf(`"%s"`, query),
			},
		},
	})
}

// AddSortCondition 添加排序
func (l *Request) AddSortCondition(field string, isReverse bool) {
	sort := "asc"
	if isReverse {
		sort = "desc"
	}
	l.Sort = append(l.Sort, map[string]string{
		field: sort,
	})
}

// AddRange 添加范围
func (l *Request) AddRange(field string, rangeData map[string]interface{}) {
	l.Query.Bool.Must = append(l.Query.Bool.Must, map[string]interface{}{
		"range": map[string]interface{}{
			field: rangeData,
		},
	})
}

func (l *Request) transQuery(query string) string {
	//+ - = && || > < ! ( ) { } [ ] ^ " ~ * ? : \ /
	query = strings.Replace(query, `+`, `\+`, -1)
	query = strings.Replace(query, `-`, `\-`, -1)
	query = strings.Replace(query, `=`, `\=`, -1)
	query = strings.Replace(query, `&`, `\&`, -1)
	query = strings.Replace(query, `|`, `\|`, -1)
	query = strings.Replace(query, `>`, `\>`, -1)
	query = strings.Replace(query, `<`, `\<`, -1)
	query = strings.Replace(query, `(`, `\(`, -1)
	query = strings.Replace(query, `)`, `\)`, -1)
	query = strings.Replace(query, `{`, `\{`, -1)
	query = strings.Replace(query, `}`, `\}`, -1)
	query = strings.Replace(query, `[`, `\[`, -1)
	query = strings.Replace(query, `]`, `\]`, -1)
	query = strings.Replace(query, `^`, `\^`, -1)
	query = strings.Replace(query, `"`, `\"`, -1)
	query = strings.Replace(query, `~`, `\~`, -1)
	query = strings.Replace(query, `*`, `\*`, -1)
	query = strings.Replace(query, `?`, `\?`, -1)
	query = strings.Replace(query, `:`, `\:`, -1)
	query = strings.Replace(query, `\`, `\\`, -1)
	query = strings.Replace(query, `/`, `\/`, -1)
	return fmt.Sprintf(`"%s"`, query)
}

// TransfromJSON 序列化结果
func (l *Request) TransfromJSON() (requestStr string, err error) {
	data, err := json.Marshal(l)
	if err != nil {
		err = fmt.Errorf("序列化Request内容失败:err:%+v", err)
		return
	}
	requestStr = string(data)
	return
}
