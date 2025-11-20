package storage

type Storage struct {
	Usst  *UserStorageImpl
	Tst   *TeamStorageImpl
	PRst  *PullRequestStorageImpl
	TxMan *TxManagerImpl
}

func NewStorage(db *DBObject) *Storage {
	return &Storage{
		Usst:  NewUserStorage(db),
		Tst:   NewTeamStorage(db),
		PRst:  NewPullRequestStorage(db),
		TxMan: NewTxManager(db),
	}
}
