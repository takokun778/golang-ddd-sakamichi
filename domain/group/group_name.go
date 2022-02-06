package group

type groupName string

func (n groupName) Value() string {
	return string(n)
}

func NewGroupNama(value string) (groupName, error) {
	// TODO validate
	return groupName(value), nil
}
