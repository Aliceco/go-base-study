package servicesb

type UserService struct {

}
func (us UserService) Get(id int) string {
	return "B单个用户信息"
}