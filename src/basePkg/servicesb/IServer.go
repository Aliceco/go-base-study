package servicesb

type IServer interface {
	Get(id int) string
}

type ServiceFactory struct {

}

func (sf ServiceFactory) Create(name string) IServer  {
	switch name {
	case "news":
		return NewService{}
	case "user":
		return UserService{}
	default:
		return nil
	}
}