package serv

type Image struct {
	Date      uint32
	Id        int64
	Location1 float32
	Location2 float32
	URL1      uint8
	URL2      uint16
	URL3      string
}

type Tag struct {
	Name string
	Imgs []Image
}

var TagData []Tag
