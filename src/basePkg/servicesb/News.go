package servicesb

type NewService struct {

}

func (ns NewService) Get(id int) string  {
	return "B单个新闻内容"
}