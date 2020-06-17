package response

type (
	appError interface {
		Error() string
		Code() uint32
		Message() string
	}
)
