package service

type ServiceGroup struct {
	AuthService
	Demo
}

var CommonServiceGroup = new(ServiceGroup)
