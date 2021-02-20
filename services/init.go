package services

type service interface {
	setup()
}

func Setup(service service) {
	service.setup()
}
