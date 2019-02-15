// Code generated by go generate; DO NOT EDIT.
package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *Log) GetParserName() string {
	return "log"
}

func (p *Log) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *Log) GetOne(index int) (common.ParserData, error) {
	if index < 0 || index >= len(p.data) {
		return nil, errors.FetchError
	}
	return p.data[index], nil
}

func (p *Log) Delete(index int) error {
	if index < 0 || index >= len(p.data) {
		return errors.FetchError
	}
	copy(p.data[index:], p.data[index+1:])
	p.data[len(p.data)-1] = types.Log{}
	p.data = p.data[:len(p.data)-1]
	return nil
}

func (p *Log) Insert(data common.ParserData, index int) error {
	if data == nil {
		return errors.InvalidData
	}
	switch newValue := data.(type) {
	case []types.Log:
		p.data = newValue
	case *types.Log:
		if index > -1 {
			p.data = append(p.data, types.Log{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = *newValue
		} else {
			p.data = append(p.data, *newValue)
		}
	case types.Log:
		if index > -1 {
			p.data = append(p.data, types.Log{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = newValue
		} else {
			p.data = append(p.data, newValue)
		}
	default:
		return errors.InvalidData
	}
	return nil
}

func (p *Log) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.Log:
		p.data = newValue
	case *types.Log:
		if index > -1 && index < len(p.data) {
			p.data[index] = *newValue
		} else {
			return errors.IndexOutOfRange
		}
	case types.Log:
		if index > -1 && index < len(p.data) {
			p.data[index] = newValue
		} else {
			return errors.IndexOutOfRange
		}
	default:
		return errors.InvalidData
	}
	return nil
}
