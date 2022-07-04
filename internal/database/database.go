package database

type DatabaseRepo interface {
	GetAll() string
}

type Database struct {
	DBRepo DatabaseRepo
}

func NewDatabase(dbRepo DatabaseRepo) *Database {
	return &Database{
		DBRepo: dbRepo,
	}
}

func (db *Database) GetAll() string {
	return db.DBRepo.GetAll()
}
