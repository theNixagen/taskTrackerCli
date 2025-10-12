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
		return "Todo"
	case InProgress:
		return "In progress"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}

func (s *Status) Decode(strStatus string) {
	switch strStatus {
	case "Todo":
		*s = Todo
	case "In progress":
		*s = InProgress
	case "Done":
		*s = Done
	}
}
