package readers

import (
	"errors"
	"strings"
)

var (
	ErrWrongNumberOfParams = errors.New("wrong number of params passed")
)

type StringReader struct {
	data   string
	params []string
}

func (sr *StringReader) Read(params ...string) ([][]string, error) {
	if len(params) < 3 {
		return nil, ErrWrongNumberOfParams
	}
	sr.data = params[0]
	sr.params = params[1:]
	var rowsAndcolumns [][]string
	lines := strings.Split(params[0], params[1])
	for _, line := range lines {
		row := Split(line, params[2])
		if len(row) > 0 {
			rowsAndcolumns = append(rowsAndcolumns, row)
		}
	}
	return rowsAndcolumns, nil
}

func (sr *StringReader) Format(fns ...func(input ...interface{}) (interface{}, error)) error {
	for _, v := range fns {
		if _, err := v(); err != nil {
			return err
		}
	}
	return nil
}

// //For str,lsep,csep are the params to be passed
// func (sr *StringReader) SetParams(parms ...string) error {
// 	if len(parms) < 3 {
// 		return ErrWrongNumberOfParams
// 	}
// 	sr.data = parms[0]
// 	sr.params = parms[1:]
// 	return nil
// }

func Split(str, sep string) []string {
	var arow []string
	srow := strings.Split(str, sep)
	for _, v := range srow {
		if v != "" {
			arow = append(arow, strings.Trim(v, " "))
		}
	}
	return arow
}
