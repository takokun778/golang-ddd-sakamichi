package member

type memberName string

func (m memberName) Value() string {
	return string(m)
}

func NewMemberName(name string) (memberName, error) {
	// TODO validate
	return memberName(name), nil
}
