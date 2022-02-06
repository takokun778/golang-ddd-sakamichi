package member

import "github.com/google/uuid"

type MemberId uuid.UUID

func (i MemberId) Value() uuid.UUID {
	return uuid.UUID(i)
}

func GenerateMemberId() MemberId {
	return MemberId(uuid.New())
}

func NewMemberId(id string) (MemberId, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return MemberId{}, err
	}
	return MemberId(uuid), nil
}
