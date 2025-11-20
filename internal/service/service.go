package service

type Service struct {
	UserService *UserServiceImpl
	TeamService *TeamServiceImpl
}

func NewService(usSt UserStorage, txman TxManagerStorage, tts TeamServiceTeamStorage, tsu TeamServiceUserStorage) *Service {
	return &Service{
		UserService: NewUserService(usSt),
		TeamService: NewTeamService(tts, tsu, txman),
	}
}

type TxManagerStorage interface {
}
