package reader

import (
	"errors"
	"strings"
)

var (
	ErrColNotFound        = errors.New("column not found")
	ErrNotTable           = errors.New("not a table")
	ErrNilTableReader     = errors.New("nil table reader")
	ErrNilOrEmptyRow      = errors.New("nil or empty table row")
	ErrNilOrEmptyHeader   = errors.New("nil or empty table header")
	ErrOutOfBoundsRowCell = errors.New("row cell index out of bounds")
	ErrOutOfBoundsColCell = errors.New("column cell index out of bounds")
)

type Reader interface {
	Read(str ...string) ([][]string, error) // Convert converts data to table format
	//Read1(str ...string) ([][]string, error) // Convert converts data to table format
	Format(fns ...func(input ...interface{}) (interface{}, error)) error
	//Convert(params ...string) ([][]string, error)
	//SetParams(str ...string) error // SetParams are to set the parameters
}

type TableReader struct {
	table     [][]string
	hasHeader bool
	Reader    Reader
	isTable   bool // is all rows have same number of columns then can consider it as table form
}

func New(reader Reader, str ...string) (*TableReader, error) {
	//reader.SetParams(str[0], str[1], str[2])
	table, err := reader.Read(str...)
	if err != nil {
		return nil, err
	}
	tr := &TableReader{table: table, hasHeader: true}
	return tr, nil
}

func NewFromStr(str, lsep, csep string) (*TableReader, error) {
	table, err := GetFromStr(str, lsep, csep)
	if err != nil {
		return nil, err
	}
	tr := &TableReader{table: table, hasHeader: true}
	return tr, nil
}

func (tr *TableReader) SetHasHeader(hasHeader bool) error {
	if tr == nil {
		return ErrNilTableReader
	}
	if tr.table == nil {
		return ErrNilTableReader
	}
	tr.hasHeader = hasHeader
	return nil
}

func (tr *TableReader) SetHeader(header []string) error {
	if tr == nil {
		return ErrNilTableReader
	}
	if tr.table == nil {
		return ErrNilTableReader
	}
	if len(header) == 0 {
		return ErrNilOrEmptyHeader
	}

	tr.table = append(tr.table[:1], tr.table[0:]...)
	tr.table[0] = header
	tr.hasHeader = true // Set the header to true irrespective of the status
	return nil
}

func (tr *TableReader) AddRow(row []string) error {
	if tr == nil {
		return ErrNilTableReader
	}
	if tr.table == nil {
		return ErrNilTableReader
	}
	if len(row) == 0 {
		return ErrNilOrEmptyRow
	}
	tr.table = append(tr.table, row)

	return nil
}
func (tr *TableReader) GetCell(row, col uint) (string, error) {
	if tr == nil {
		return "", ErrNilTableReader
	}
	if tr.table == nil {
		return "", ErrNilTableReader
	}
	if !(int(row) < len(tr.table)) {
		return "", ErrOutOfBoundsRowCell
	}
	if !(int(col) < len(tr.table[row])) {
		return "", ErrOutOfBoundsColCell
	}
	return tr.table[row][col], nil
}

func (tr *TableReader) GetCellsByColumn(col string) ([]string, error) {
	index, err := tr.GetColIndex(col)
	if err != nil {
		return nil, err
	}
	if index >= 0 {
		var arr []string
		for i := 1; i < len(tr.table); i++ {
			if len(tr.table[i]) >= index {
				arr = append(arr, tr.table[i][index])
			}
		}
		return arr, nil
	}
	return nil, nil
}

func (tr *TableReader) GetColIndex(col string) (int, error) {
	if tr.hasHeader {
		if tr == nil {
			return -1, ErrNilTableReader
		}
		if tr.table == nil {
			return -1, ErrNilTableReader
		}
		for i, c := range tr.table[0] {
			if strings.EqualFold(c, col) {
				return i, nil
			}
		}
		return -1, ErrColNotFound
	}
	return -1, ErrNilOrEmptyHeader
}

func (tr *TableReader) GetTable() ([][]string, error) {
	if tr == nil {
		return nil, ErrNilTableReader
	}
	if tr.table == nil {
		return nil, ErrNilTableReader
	}
	return tr.table, nil
}

func (tr *TableReader) Length() (rowLength, colLength int) {
	if tr == nil {
		return -1, -1
	}
	if tr.table == nil {
		return -1, -1
	}
	return len(tr.table), len(tr.table[0])
}

func GetFromStr(str, lsep, csep string) ([][]string, error) {
	if !isTable(str) {
		return nil, ErrNotTable
	}
	var rowsAndcolumns [][]string
	lines := strings.Split(str, lsep)
	for _, line := range lines {
		row := Split(line, csep)
		if len(row) > 0 {
			rowsAndcolumns = append(rowsAndcolumns, row)
		}
	}
	return rowsAndcolumns, nil
}

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

// isTable is to check with the given data , can table be formed or not
func isTable(str string) bool {
	return strings.Trim(str, " ") != ""
}
