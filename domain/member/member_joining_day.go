package member

import "time"

type memberJoiningDay time.Time

func (j memberJoiningDay) Value() time.Time {
	return time.Time(j)
}

func (j memberJoiningDay) Format() string {
	return j.Value().Format("2006/1/2")
}

func NewMemberJoiningDay(value time.Time) (memberJoiningDay, error) {
	// TODO validate
	return memberJoiningDay(value), nil
}
