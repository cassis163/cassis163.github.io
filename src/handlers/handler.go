package handlers

type Handler interface {
	Handle()
	New()
}
