package api

type NotifierApi interface {
	Notify(Target string, Message string) (StatusCode int, Error error)
}
