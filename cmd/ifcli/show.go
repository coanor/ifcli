package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/influxdata/influxdb1-client/models"
	client "github.com/influxdata/influxdb1-client/v2"
)

func showResp(r *client.Response) {

	switch strings.ToUpper(curFMT) {
	case `JSON`:
		jsonShow(r)
	default:
		defaultShow(r)
	}
}

func getMaxColLen(r *models.Row) int {
	maxColLen := 0
	for _, col := range r.Columns {
		if len(col) > maxColLen {
			maxColLen = len(col)
		}
	}

	return maxColLen
}

func defaultShow(r *client.Response) {

	for _, res := range r.Results {
		for _, s := range res.Series {

			switch len(s.Columns) {
			case 1:
				for _, val := range s.Values {
					fmt.Println(val[0])
				}
			default:
				maxColLen := getMaxColLen(&s)
				fmtStr := "%" + fmt.Sprintf("%d", maxColLen) + "s\t%v\n"

				for _, val := range s.Values {
					for colIdx, _ := range s.Columns {
						if disableNil && val[colIdx] == nil {
							continue
						}

						fmt.Printf(fmtStr, s.Columns[colIdx], val[colIdx])
					}

					fmt.Printf("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n")
				}
			}
		}
	}
}

func jsonShow(r *client.Response) {
	j, err := json.Marshal(r)
	if err == nil {
		fmt.Printf("%s", string(j))
	}
}
