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
	lat  string
	lon  string
	date string
	url  string
}

type RetJson struct {
	tag     string
	results []RetImg
}

var TagData []Tag
