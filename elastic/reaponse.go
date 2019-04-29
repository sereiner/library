package elastic

import (
	"strings"
)

// Response 返回结果
type Response struct {
	Took int64    `json:"took"`
	Hits HitsType `json:"hits"`
}

// HitsType HitsType
type HitsType struct {
	Total int64          `json:"total"`
	Hits  []HitsDataType `json:"hits"`
}

// HitsDataType HitsDataType
type HitsDataType struct {
	Source   map[string]interface{} `json:"_source"`
	HighList map[string][]string    `json:"highlight"`
}

// Transform 转换请求结果
func (r *Response) Transform() {
	if len(r.Hits.Hits) == 0 {
		return
	}

	for _, item := range r.Hits.Hits {
		for k, v := range item.HighList {
			// item.Source[k] = v[0]
			origStr := item.Source[k].(string)
			for _, item := range v {
				repStr := item
				//repStr = strings.Replace(repStr, HighTabPre, "", -1)
				//repStr = strings.Replace(repStr, HighTabSuffix, "", -1)
				origStr = strings.Replace(origStr, repStr, item, -1)
			}
			item.Source[k] = origStr
		}
	}
}

// Analysis 解析结果
func (r *Response) Analysis() map[string]interface{} {
	r.Transform()

	logList := map[int]interface{}{}
	for i, item := range r.Hits.Hits {
		logList[i] = item.Source
	}
	total := r.Hits.Total
	took := r.Took

	return map[string]interface{}{
		"total_num": total,
		"log_list":  logList,
		"took_time": took,
	}
}
