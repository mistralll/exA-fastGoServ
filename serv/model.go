package serv

type Image struct {
	Date      uint32
	Id        int64
	Location1 float32
	Location2 float32
	URL1      uint32
	URL2      uint32
	URL3      string
}

type Tag struct {
	Name string
	Imgs []Image
}

type RetImg struct {
	Lat  float32 `json:"lat,omitempty"`
	Lon  float32 `json:"lon,omitempty"`
	Date string `json:"date,omitempty"`
	Url  string `json:"url,omitempty"`
}

type RetTag struct {
	Tag     string   `json:"tag,omitempty"`
	Results []RetImg `json:"results,omitempty"`
}

var TagData []Tag
