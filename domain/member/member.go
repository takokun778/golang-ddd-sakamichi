package member

import "sakamichi/domain/group"

type member struct {
	memberProps
}

func (m member) Props() MemberProps {
	return MemberProps(m)
}

func constructor(
	memberId MemberId,
	groupId group.GroupId,
	firstName memberName,
	lastName memberName,
	birthday memberBirthday,
	blood memberBlood,
	joiningDay memberJoiningDay,
) member {
	member := new(member)
	member.memberId = memberId
	member.groupId = groupId
	member.firstName = firstName
	member.lastName = lastName
	member.birthday = birthday
	member.blood = blood
	member.joiningDay = joiningDay
	return *member
}

func Reconstruct(props MemberProps) member {
	return constructor(
		props.memberId,
		props.groupId,
		props.firstName,
		props.lastName,
		props.birthday,
		props.blood,
		props.joiningDay,
	)
}

func Debut(
	groupId group.GroupId,
	firstName memberName,
	lastName memberName,
	birthday memberBirthday,
	blood memberBlood,
	joiningDay memberJoiningDay,
) member {
	return constructor(
		GenerateMemberId(),
		groupId,
		firstName,
		lastName,
		birthday,
		blood,
		joiningDay,
	)
}

func (m member) Transfer(groupId group.GroupId) member {
	return constructor(
		m.memberId,
		groupId,
		m.firstName,
		m.lastName,
		m.birthday,
		m.blood,
		m.joiningDay,
	)
}
