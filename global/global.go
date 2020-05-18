package global

var DataInfoUrl chan *DataUrl
var AllUrl chan string

func Init() {
	DataInfoUrl = make(chan *DataUrl, 50)
	AllUrl = make(chan string, 100)
}

type DataUrl struct {
	Url      string
	Title    string
	Keywords string
	Desc     string
	ImgPath  string
}