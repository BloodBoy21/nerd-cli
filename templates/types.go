package templates

type FlagOption interface {
	int | string | bool
}
type Configs struct {
	Token string `json:"token"`
}