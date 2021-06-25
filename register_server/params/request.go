package params

type RequestRegister struct {
	Env      string
	AppID    string
	HostName string
	Addrs    []string
	Version  string
	Status   uint32
}
