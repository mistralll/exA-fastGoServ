package serv

type Image struct {
	id      string
	date    string
	locate1 string
	locate2 string
	URL     string
}

type Tag struct {
	name string
	id   string
}

type ImageList struct {
	images []Image
}
