package repo

type FileRepository struct {
	BaseRepo
}

var file *FileRepository

func init() {
	file = new(FileRepository)
}

func NewFileRepository() *FileRepository {
	return file
}
