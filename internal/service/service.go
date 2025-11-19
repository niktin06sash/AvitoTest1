package service

type Service struct {
	UserService *UserServiceImpl
}

func NewService(usSt UserStorage) *Service {
	return &Service{
		UserService: NewUserService(usSt),
	}
}
