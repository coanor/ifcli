package ifcli

import (
	"fmt"
	"strings"

	"github.com/influxdata/influxdb1-client/models"
	client "github.com/influxdata/influxdb1-client/v2"
)

func ShowResp(r *client.Response) int {

	switch strings.ToUpper(CurFMT) {
	case `JSON`:
		// not IMPL
	default:
		return defaultShow(r)
	}
	return 0
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

func defaultShow(r *client.Response) int {

	nrows := 0

	for _, res := range r.Results {
		for _, s := range res.Series {

			switch len(s.Columns) {
			case 1:
				fmt.Printf("%s\n", s.Name)
				fmt.Println("--------------")
				for _, val := range s.Values {
					// measurements or dbs, we can add them as suggestions

					switch val[0].(type) {
					case string:
						AddSug(val[0].(string))
					}

					fmt.Println(val[0])
					nrows++
				}

			default:
				maxColLen := getMaxColLen(&s)
				fmtStr := "%" + fmt.Sprintf("%d", maxColLen) + "s\t%v\n"

				for _, val := range s.Values {

					nrows++
					fmt.Printf("-=-=-=-=-=-=-=-=[ %d. Row ]-=-=-=-=-=-=-=-=-\n", nrows)

					for colIdx, _ := range s.Columns {
						if DisableNil && val[colIdx] == nil {
							continue
						}

						fmt.Printf(fmtStr, s.Columns[colIdx], val[colIdx])

						AddSug(s.Columns[colIdx])
					}
				}
			}
		}
	}

	return nrows
}
