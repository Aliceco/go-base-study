package services

type NewService struct {

}

func (ns NewService) Get(id int) string  {
	return "单个新闻内容"
}