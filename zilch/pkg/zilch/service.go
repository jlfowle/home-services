package zilch

type ZilchService interface {
	Get() error
}

func NewZilchService() ZilchService {
	return &zilchService{}
}

type zilchService struct{}

func (zilchService) Get() error {
	return nil
}
