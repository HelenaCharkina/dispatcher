package types

type Agent struct {
	Id          int    `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
