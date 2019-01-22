package filter

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"strings"
)

const (
	Pre  Type = 0
	Post Type = 1

	GlobalFilterPath = "*"
)

var globalCondition = Condition{path: GlobalFilterPath}
var filterChain = make(FChain, 0)

type Filter interface {
	Do(request *model.IrisReq, data interface{}) (interface{}, types.BizCode)
	Match() []Condition
	Type() Type
}

type Condition struct {
	path   string
	method string
}

func (c Condition) isGlobal() bool {
	return c.path == GlobalFilterPath
}

type FChain []Filter
type Type int

func Register(filter Filter) {
	filterChain = append(filterChain, filter)
}

func DoFilters(req *model.IrisReq, data interface{}, typ Type) (interface{}, types.BizCode) {
	var reqUrl = strings.TrimPrefix(req.RequestURI, types.UrlRoot)
	for _, f := range filterChain {
		if typ != f.Type() {
			continue
		}
		conditions := f.Match()
		if conditions[0].isGlobal() {
			data, err := f.Do(req, data)
			if !err.Success() {
				return data, err
			}
		} else {
			for _, condition := range conditions {
				p := removeUrlParam(condition.path)
				if strings.HasPrefix(reqUrl, p) && equal(condition.method, req.Method) {
					data, err := f.Do(req, data)
					if !err.Success() {
						return data, err
					}
				}
			}
		}

	}
	return nil, types.CodeSuccess
}

func removeUrlParam(url string) string {
	var starting = "{"
	firstIndex := strings.Index(url, starting)
	if firstIndex < 0 {
		return url
	}
	return url[:firstIndex]
}

func equal(src, dest string) bool {
	src = strings.ToLower(src)
	dest = strings.ToLower(dest)
	return src == dest
}
