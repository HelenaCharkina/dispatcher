package types

type Agent struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Ip          string `json:"ip" db:"ip"`
	Port        string `json:"port" db:"port"`
	Description string `json:"description" db:"description"`
	Schedule    string `json:"schedule" db:"schedule"`
	State       State  `json:"state" db:"state"`
}

type State byte

const (
	STARTED State = 2
	STOPPED State = 1
	ERROR   State = 3
)
