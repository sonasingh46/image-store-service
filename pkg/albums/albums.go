package albums

type Album struct {
	Name string `json:"name"`
	Location string `json:"location"`
}

type Image struct {
	Path string `json:"path"`
}

type StoreService interface {
	CreateAlbum(Album) error
	DeleteAlbum(string) error
	UploadImage(Image,string) error
	DeleteImage(string,string) error
}

