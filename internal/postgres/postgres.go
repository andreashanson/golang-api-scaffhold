package postgres

type PostgresDB struct{}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (pdb *PostgresDB) GetAll() string {
	return "GET ALL POSTGRES"
}
