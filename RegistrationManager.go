package main

type RegistrationManager interface {
	Manage(configs RegistrationVariables)
	RegisterWithSerivceRegistry()
	SendHeartBeat(configs RegistrationVariables)
	DeRegisterFromServiceRegistry(configs RegistrationVariables)
}

type RegistrationVariables struct {
	registryType       string
	serviceRegistryURL string
}

func (rv RegistrationVariables) RegistryType() string {
	return rv.registryType
}

func (rv RegistrationVariables) ServiceRegistryURL() string {
	return rv.serviceRegistryURL
}

