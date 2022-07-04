package mongo

type MongoDB struct {
}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (mdb *MongoDB) GetAll() string {
	return "GET ALL MONGO"
}
