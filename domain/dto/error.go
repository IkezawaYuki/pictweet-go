package dto

type PicError struct {
}

func (e *PicError) Error() string {
	return "error"
}
