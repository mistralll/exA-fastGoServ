package serv

import (
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

	ans.date = d
	ans.lat = strconv.FormatFloat(float64(img.Location1), 'f', 2, 64)
	ans.lon = strconv.FormatFloat(float64(img.Location2), 'f', 2, 64)

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

	ans.url = sb.String()

	return ans, nil
}

func TagToRetTag(tag Tag)(RetTag, error) {
	ans := RetTag{}
	ans.tag = tag.Name
	ans.results = make([]RetImg, len(tag.Imgs))

	for i, row := range tag.Imgs {
		tmp, err := ImageToRetImg(row)
		if err != nil {
			return RetTag{}, err
		}
		ans.results[i] = tmp
	}

	return ans, nil
}

func RetTagToHTML(tag RetTag) string {
	var sb strings.Builder
	sb.WriteString("<html>")
	sb.WriteString(tag.tag)
	for _, row := range tag.results {
		sb.WriteString("<img src=")
		sb.WriteString(row.url)
		sb.WriteString("></img>")
		sb.WriteString(row.date)
		sb.WriteString(row.lat)
		sb.WriteString(row.lon)
	}
	sb.WriteString("</html>")
	return sb.String()
} 