package member

import "sakamichi/domain/group"

type MemberProps struct {
	memberProps
}

type memberProps struct {
	memberId   MemberId
	groupId    group.GroupId
	firstName  memberName
	lastName   memberName
	birthday   memberBirthday
	blood      memberBlood
	joiningDay memberJoiningDay
}

func (p memberProps) MemberId() MemberId {
	return p.memberId
}

func (p memberProps) GroupId() group.GroupId {
	return p.groupId
}

func (p memberProps) FirstName() memberName {
	return p.firstName
}

func (p memberProps) LastName() memberName {
	return p.lastName
}

func (p memberProps) Birthday() memberBirthday {
	return p.birthday
}

func (p memberProps) Blood() memberBlood {
	return p.blood
}

func (p memberProps) JoiningDay() memberJoiningDay {
	return p.joiningDay
}

func NewMemberProps(
	memberId MemberId,
	groupId group.GroupId,
	firstName memberName,
	lastName memberName,
	birthday memberBirthday,
	blood memberBlood,
	joiningDay memberJoiningDay,
) MemberProps {
	props := new(MemberProps)
	props.memberId = memberId
	props.groupId = groupId
	props.firstName = firstName
	props.lastName = lastName
	props.birthday = birthday
	props.blood = blood
	props.joiningDay = joiningDay
	return *props
}
