package types

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	switch s {
	case Todo:
		return "Pending"
	case InProgress:
		return "In progress"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}
