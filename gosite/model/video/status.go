package status

type VideoStatusValue byte

const (
	CREATED    = 1
	PROCESSING = 2
	READY      = 3
	DELETED    = 4
	ERROR      = 5
)