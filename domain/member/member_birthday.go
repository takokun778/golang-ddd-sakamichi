package member

import "time"

type memberBirthday time.Time

func (b memberBirthday) Value() time.Time {
	return time.Time(b)
}

func (b memberBirthday) Format() string {
	return b.Value().Format("2006/1/2")
}

func NewMemberBirthday(value time.Time) (memberBirthday, error) {
	// TODO validate
	return memberBirthday(value), nil
}
