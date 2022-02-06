package group

import "github.com/google/uuid"

type GroupId uuid.UUID

func (i GroupId) Value() uuid.UUID {
	return uuid.UUID(i)
}

func GenerateGroupId() GroupId {
	return GroupId(uuid.New())
}

func NewGroupId(id string) (GroupId, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return GroupId{}, err
	}
	return GroupId(uuid), nil
}
