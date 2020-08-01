package albums

type Album struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type AlbumList struct {
	Items []Album `json:"items"`
}

type Image struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type StoredImage struct {
	AlbumName string `json:"albumName"`
	Name      string `json:"name"`
}

type StoredImageList struct {
	Items []StoredImage `json:"items"`
}

type StoreService interface {
	CreateAlbum(album Album) error
	DeleteAlbum(albumName string) error
	ListAlbums() (*AlbumList, error)
	UploadImage(image Image, albumName string) error
	DeleteImage(imageName, albumName string) error
	ListImages(albumName string) (*StoredImageList, error)
	GetImage(imageName, albumName string) (*StoredImage, error)
}
