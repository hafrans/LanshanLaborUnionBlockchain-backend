package smsqueue

import (
	"bytes"
	"fmt"
	"net/url"
)

type KeyValuePair struct {
	Key   string
	Value string
}
type QueryValues []*KeyValuePair

func NewQueryValues() QueryValues {
	return make(QueryValues, 0, 2)
}

func (v *QueryValues) Add(key, value string) {
	*v = append(*v, &KeyValuePair{
		key, value,
	})
}

func (v *QueryValues) Encode() string {

	if len(*v) == 0 {
		return ""
	}
	var buf bytes.Buffer
	flag := false
	for _, v := range *v {
		if flag {
			buf.WriteString("&")
		} else {
			flag = true
		}
		buf.WriteString(fmt.Sprintf("%s=%s", url.QueryEscape(v.Key), url.QueryEscape(v.Value)))
	}
	return buf.String()
}

