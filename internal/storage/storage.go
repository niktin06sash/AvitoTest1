package storage

type Storage struct {
	Usst  *UserStorageImpl
	Tst   *TeamStorageImpl
	TxMan *TxManagerImpl
}

func NewStorage(db *DBObject) *Storage {
	return &Storage{
		Usst:  NewUserStorage(db),
		Tst:   NewTeamStorage(db),
		TxMan: NewTxManager(db),
	}
}
