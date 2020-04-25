package core

type MyController struct {
	Context *MyContext
}

func (this *MyController) Init(Context *MyContext) {
	this.Context = Context
}

type ControllerInterface interface {
	Init(Context *MyContext)
	GET()
	POST()
}