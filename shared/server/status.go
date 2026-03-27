package server

type Status int

const (
	OK Status = iota
	Shutdown
)

func (s Status) String() string {
	switch s {
	case Shutdown:
		return "shutdown"
	default:
		return "ok"
	}
}

func (g *GracefulShutdown) GetServerStatus() Status {
	if g.getShutdownProcess() {
		return Shutdown
	}
	return OK
}
