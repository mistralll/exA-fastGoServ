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

func ImageToRetImg(img Image) (RetImg, error) {
	ans := RetImg{}
	d, err := uintTimeToStr(img.Date)
	if err != nil {
		return ans, err
	}

	ans.Date = d
	ans.Lat = img.Location1
	ans.Lon = img.Location2

	var sb strings.Builder

	sb.WriteString("http://farm")
	sb.WriteString(strconv.FormatUint(uint64(img.URL1), 10))
	sb.WriteString(".static.flickr.com/")
	sb.WriteString(strconv.FormatUint(uint64(img.URL2), 10))
	sb.WriteString("/")
	sb.WriteString(strconv.FormatInt(img.Id, 10))
	sb.WriteString("_")
	sb.WriteString(img.URL3)
	sb.WriteString(".jpg")

	ans.Url = sb.String()

	return ans, nil
}

func TagToRetTag(tag Tag) (RetTag, error) {
	ans := RetTag{}
	ans.Tag = tag.Name
	ans.Results = make([]RetImg, len(tag.Imgs))

	for i, row := range tag.Imgs {
		tmp, err := ImageToRetImg(row)
		if err != nil {
			return RetTag{}, err
		}
		ans.Results[i] = tmp
	}

	return ans, nil
}

func RetTagToHTML(tag RetTag) string {
	var sb strings.Builder
	sb.WriteString("<html>")
	sb.WriteString(tag.Tag)
	for _, row := range tag.Results {
		sb.WriteString("<img src=")
		sb.WriteString(row.Url)
		sb.WriteString("></img>")
		sb.WriteString(row.Date)
		sb.WriteString(strconv.FormatFloat(float64(row.Lat), 'f', 2, 10))
		sb.WriteString(strconv.FormatFloat(float64(row.Lon), 'f', 2, 10))
	}
	sb.WriteString("</html>")
	return sb.String()
}

func RetJson(tag Tag, w http.ResponseWriter) error {
	var sb strings.Builder
	sb.WriteString(`{"tag":"`)
	sb.WriteString(tag.Name)
	sb.WriteString(`", "results":[`)

	for i, row := range tag.Imgs {
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
		if i == len(tag.Imgs)-2 {
			sb.WriteString(`,`)
		} else {
			sb.WriteString(`]}`)
		}
	}

	w.Write([]byte(sb.String()))

	return nil
}
