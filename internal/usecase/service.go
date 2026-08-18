package usecase

import "github.com/t4vr1s/go-service-template/internal/domain"

type HealthCheck struct{}

func (HealthCheck) Execute() domain.HealthStatus {
	return domain.HealthStatus{Status: "ok"}
}

type Greeter struct{}

func (Greeter) Execute(name string) domain.Greeting {
	if name == "" {
		name = "world"
	}

	return domain.Greeting{Message: "hello, " + name}
}
