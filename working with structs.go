//The goal is to get data in format:

/*
{
        "kind": {
                "Kind": true
        },
        "text": "",
        "rows": [
                {
                        "column": {
                                "content": "1",
                                "tag": "1"
                        }
                },
                {
                        "column": {
                                "content": "12",
                                "tag": "12"
                        }
                },
                {
                        "column": {
                                "content": "30",
                                "tag": "30"
                        }
                },
                {
                        "column": {
                                "content": "4",
                                "tag": "4"
                        }
                },
                {
                        "column": {
                                "content": "5",
                                "tag": "5"
                        }
                }
        ]
}
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ColumnStructure struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}
type RowStructure struct {
	Column ColumnStructure `json:"column"`
}

type format struct {
	Kind map[string]bool `json:"kind"`
	Text string          `json:"text"`
	Rows []RowStructure  `json:"rows"`
}

func serialize(s []int) string {
	var formatted format
	formatted.Kind = map[string]bool{"Kind": true}

	for i := 0; i < len(s); i++ {
		var col_i = ColumnStructure{
			Content: strconv.Itoa(s[i]),
			Tag:     strconv.Itoa(s[i]),
		}
		var row = RowStructure{
			Column: col_i,
		}
		formatted.Rows = append(formatted.Rows, row)
	}
	formatted_output, _ := json.MarshalIndent(formatted, "", "\t")
	return string(formatted_output)
}

func main() {
	var s = []int{1, 12, 30, 4, 5}
	formatted_data := serialize(s)

	fmt.Print(string(formatted_data))

}
