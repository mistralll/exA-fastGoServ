package serv

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

func strTimeToUint(instr string) (uint32, error) {
	t1, err := time.Parse("2006-01-02 15:04:05", instr)
	if err != nil {
		return 0, err
	}

	t2 := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)

	ans := t1.Sub(t2).Seconds()
	return uint32(ans), nil
}

func uintTimeToStr(inuint uint32) (string, error) {
	t2 := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
	ans := t2.Add(time.Duration(inuint) * time.Second)
	return ans.Format("2006-01-02 15:04:05"), nil
}

func RetJson(index int, w http.ResponseWriter) error {
	var sb strings.Builder
	sb.WriteString(`{"tag":"`)
	sb.WriteString(TagData[index].Name)
	sb.WriteString(`", "results":[`)

	for i, row := range TagData[index].Imgs {
		sb.WriteString(`{"lat":`)
		sb.WriteString(strconv.FormatFloat(float64(row.Location1), 'f', 2, 64))
		sb.WriteString(`,"lon":`)
		sb.WriteString(strconv.FormatFloat(float64(row.Location2), 'f', 2, 64))
		sb.WriteString(`,"date":"`)
		d, err := uintTimeToStr(row.Date)
		if err != nil {
			return err
		}
		sb.WriteString(d)
		sb.WriteString(`","url":":http://farm`)
		sb.WriteString(strconv.FormatUint(uint64(row.URL1), 10))
		sb.WriteString(`.static.flickr.com/`)
		sb.WriteString(strconv.FormatUint(uint64(row.URL2), 10))
		sb.WriteString("/")
		sb.WriteString(strconv.FormatInt(row.Id, 10))
		sb.WriteString("_")
		sb.WriteString(row.URL3)
		sb.WriteString(`.jpg"}`)
		if i == len(TagData[index].Imgs)-2 {
			sb.WriteString(`,`)
		} else {
			sb.WriteString(`]}`)
		}
	}

	w.Write([]byte(sb.String()))

	return nil
}
