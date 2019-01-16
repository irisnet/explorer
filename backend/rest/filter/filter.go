package filter

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
)

const GlobalFilterPath = "*"

type Filter interface {
	Do(request *model.IrisReq) (bool, interface{}, types.BizCode)
}

type Type int

const (
	Pre  Type = 0
	Post Type = 1
)

type RouteFilter map[string][]Filter

var preFilters = make(RouteFilter, 0)
var postFilters = make(RouteFilter, 0)

func RegisterFilters(path string, typ Type, fs []Filter) {
	var filers RouteFilter
	switch typ {
	case Pre:
		filers = preFilters
		break
	case Post:
		filers = postFilters
	default:
		logger.Panic("not existed filter type", logger.Any("type", typ))

	}
	if _, ok := filers[path]; ok {
		logger.Panic("duplicate registration filter", logger.String("path", path))
	}
	filers[path] = fs
}

func DoFilters(req *model.IrisReq, typ Type) (bool, interface{}, types.BizCode) {
	var filers RouteFilter
	switch typ {
	case Pre:
		filers = preFilters
		break
	case Post:
		filers = postFilters
	default:
		logger.Panic("not existed filter type", logger.Any("type", typ))

	}
	return doFilter(req, filers)
}

func doFilter(req *model.IrisReq, filters RouteFilter) (bool, interface{}, types.BizCode) {
	//check global filters
	globalFilters, ok := filters[GlobalFilterPath]
	if ok {
		for _, f := range globalFilters {
			ok, data, err := f.Do(req)
			if !ok {
				return false, data, err
			}
		}
	}

	//check custom filters
	customFilters, ok := filters[req.RequestURI]
	if ok {
		for _, f := range customFilters {
			ok, data, err := f.Do(req)
			if !ok {
				return false, data, err
			}
		}
	}
	return true, nil, types.CodeSuccess
}
