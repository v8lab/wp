package base

import (
	"net/http"
)

type EntryIntf interface {
	Init(string)
	SetKind(string)
	GetKind() string
	SetUrl(string)
	GetUrl() string
	SetBody(string)
	GetBody() string
	SetMethod(string)
	GetMethod() string
	SetHeader(http.Header)
	GetHeader() http.Header
	GetThis() *EntryStu
	Execute() (*EntryStu, int)
	Construct(interface{}) int
	ToJson() string
}
