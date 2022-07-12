package csvUtilLegacy

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func SplitAll(infileName string, outfilePath string) error {
	// readfile
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	r := csv.NewReader(infile)

	// infile to vec: data
	data := [][]string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		data = append(data, record)

		if len(data)%100000 == 0 {
			fmt.Println("reading: " + strconv.Itoa(len(data)))
		}
	}

	// split by tag
	ans := [][][]string{}
	cnt := -1
	tag := ""

	for i, record := range data {
		if record[1] != tag {
			cnt++
		}

		if len(ans) <= cnt {
			ans = append(ans, [][]string{record})
		} else {
			ans[cnt] = append(ans[cnt], record)
		}

		if i%10000 == 0 {
			fmt.Println("splitting: " + strconv.Itoa(i) + ", " + strconv.Itoa(cnt))
		}
	}

	// write to csv
	for i := 0; i < len(ans); i++ {
		outfile, err := os.Create(outfilePath + ans[i][0][0] + ".csv")
		if err != nil {
			return err
		}
		defer outfile.Close()

		w := csv.NewWriter(outfile)

		for j, rec := range ans[i] {
			if j > 100 {
				break
			}
			if err := w.Write(rec); err != nil {
				return err
			}
		}

		if i%1000 == 0 {
			fmt.Println("writing: " + strconv.Itoa(i))
		}
	}

	return nil
}
