package storage

type Storage struct {
	Usst *UserStorageImpl
}

func NewStorage(db *DBObject) *Storage {
	return &Storage{
		Usst: NewUserStorage(db),
	}
}
