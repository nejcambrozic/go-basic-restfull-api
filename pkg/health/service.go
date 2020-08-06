package health

type Service interface {
	GetHealth() (ApiUserMessage, int)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetHealth() (ApiUserMessage, int) {
	healthMessage := ApiUserMessage{Status: "ok"}
	status := 200
	return healthMessage, status
}

type ApiUserMessage struct {
	Status string
}
