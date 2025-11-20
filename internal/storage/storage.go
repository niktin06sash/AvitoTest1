package storage

type Storage struct {
	Usst *UserStorageImpl
	Tst  *TeamStorageImpl
}

func NewStorage(db *DBObject) *Storage {
	return &Storage{
		Usst: NewUserStorage(db),
		Tst:  NewTeamStorage(db),
	}
}
