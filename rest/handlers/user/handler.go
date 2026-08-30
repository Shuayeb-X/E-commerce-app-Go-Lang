package user

type Handler struct {
	FirstName string
}

func newHandler() *Handler{
  return &Handler{}
}