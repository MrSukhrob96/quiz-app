package pkg_services

type JwtService interface {
	RefreshToken()
	SignIn()
	Me()
}

type jwtService struct {
}

func NewJwtService() JwtService {
	return &jwtService{}
}

func (service *jwtService) RefreshToken() {

}

func (service *jwtService) SignIn() {

}

func (service *jwtService) Me() {

}
