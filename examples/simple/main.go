package main

import (
	"errors"
	"fmt"

	"github.com/JitenPalaparthi/tablereader/test/pkg/cmdhelper"

	readers "github.com/JitenPalaparthi/tablereader/pkg/readers"

	tablereader "github.com/JitenPalaparthi/tablereader/reader"
)

var (
	ErrNotTable = errors.New("not a table")
)

func main() {
	c, _ := cmdhelper.New(nil, nil)                  // use command helper
	str, _ := c.CliRunner("docker", nil, "ps", "-a") // pass the command assuming that it will return data in the table form
	fmt.Println(str)                                 //
	sr := &readers.StringReader{}                    // create a string reader
	tr, err := tablereader.New(sr, str, "\n", "  ")  // mass parameters 1st string reader , 2nd the actual data, 3rd line separator, 4th column separator
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tr.Length())      // gives the length of the table row and columns
	fmt.Println(tr.GetCell(6, 9)) // get the data from 6th row , 9th column
	////tr.SetHeader([]string{"IMAGE ID"})  to set headers if not present. It sets in the oth index
	fmt.Println(tr.GetTable())                  // get the table in []][]string array
	i, _ := tr.GetColIndex("created")           // get the index of the column "created". Case insensitive
	fmt.Println(i)                              // print the index
	fmt.Println(tr.GetCellsByColumn("created")) // get all values of created column

}
