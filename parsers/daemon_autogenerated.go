// Code generated by go generate; DO NOT EDIT.
package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *Daemon) Init() {
    p.data = nil
}

func (p *Daemon) GetParserName() string {
	return "daemon"
}

func (p *Daemon) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.Enabled{}
			return p.data, nil
		}
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *Daemon) GetOne(index int) (common.ParserData, error) {
	if index > 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *Daemon) Delete(index int) error {
	p.Init()
	return nil
}

func (p *Daemon) Insert(data common.ParserData, index int) error {
	return p.Set(data, index)
}

func (p *Daemon) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.Enabled:
		p.data = newValue
	case types.Enabled:
		p.data = &newValue
	default:
		return errors.InvalidData
	}
	return nil
}
