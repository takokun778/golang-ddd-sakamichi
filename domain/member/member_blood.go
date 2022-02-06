package member

import "fmt"

type memberBlood string

const (
	A       memberBlood = "A"
	B       memberBlood = "B"
	O       memberBlood = "O"
	AB      memberBlood = "AB"
	UNKNOWN memberBlood = "UNKNOWN"
)

func (b memberBlood) Value() string {
	return string(b)
}

func NewMemberBlood(value string) (memberBlood, error) {
	switch value {
	case "A", "B", "O", "AB", "UNKNOWN":
		return memberBlood(value), nil
	default:
		return memberBlood(""), fmt.Errorf("error")
	}
}
