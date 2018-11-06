package service

var Main = mainService{}

type mainService struct {
}

func (m *mainService) Stert() bool {
	return true
}

func (m *mainService) End() {
}
