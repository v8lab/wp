package base

import (
	"net/http"
)

type EntryIntf interface {
	SetId(string)
	GetId() string
	SetMethod(string)
	GetMethod() string
	SetKind(string)
	GetKind() string
	SetUrl(string)
	GetUrl() string
	SetHeader(http.Header)
	GetHeader() http.Header
	SetBody(string)
	GetBody() string
	GetData() *EntryStu
	ToJson() string
	Init(*http.Request) int
	Nop()
	Execute() int
}
