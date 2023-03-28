package infrastructure

type HelloDataSource struct{}

func (ds *HelloDataSource) GetHelloMessage() (string, error) {
	return "world", nil
}

type HelloRepository struct {
	DataSource *HelloDataSource
}

func (r *HelloRepository) GetHelloMessage() (string, error) {
	return r.DataSource.GetHelloMessage()
}

func NewHelloRepository() HelloRepository {
	return HelloRepository{
		DataSource: &HelloDataSource{},
	}
}
