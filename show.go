package ifcli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/influxdata/influxdb1-client/models"
	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	tee     io.Writer
	teeFile *os.File
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
				showFmtLine("%s\n", s.Name)
				showLine("--------------")
				for _, val := range s.Values {
					// measurements or dbs, we can add them as suggestions

					switch val[0].(type) {
					case string:
						AddSug(val[0].(string))
					}

					showLine(val[0])
					nrows++
				}

			default:
				maxColLen := getMaxColLen(&s)
				fmtStr := fmt.Sprintf("%%%ds%%s", maxColLen) + " %v\n"

				for _, val := range s.Values {

					nrows++
					showFmtLine("-=-=-=-=-=-=-=-=[ %d. Row ]-=-=-=-=-=-=-=-=-\n", nrows)

					for colIdx, _ := range s.Columns {
						if DisableNil && val[colIdx] == nil {
							continue
						}

						col := s.Columns[colIdx]
						if _, ok := s.Tags[col]; ok {
							showFmtLine(fmtStr, col, "*", val[colIdx])
						} else {
							showFmtLine(fmtStr, col, " ", val[colIdx])
						}

						AddSug(s.Columns[colIdx])
					}
				}
			}
		}
	}

	return nrows
}

func showFmtLine(fmtStr string, args ...interface{}) {
	if tee == nil {
		fmt.Printf(fmtStr, args...)
	} else {
		fmt.Fprintf(tee, fmtStr, args...)
	}
}

func showLine(args ...interface{}) {
	if tee == nil {
		fmt.Println(args...)
	} else {
		fmt.Fprintln(tee, args...)
	}
}
