package serv

type Image struct {
	Date      string
	Id        string
	Location1 string
	Location2 string
	URL1      string
	URL2      string
	URL3      string
}

type Tag struct {
	Name string
	Imgs []Image
}

var TagData []Tag
