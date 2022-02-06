package group

type group struct {
	groupProps
}

func (g group) Props() GroupProps {
	return GroupProps(g)
}

func construct(
	groupId GroupId,
	name groupName,
	formationDay groupFormationDay,
) group {
	group := new(group)
	group.groupId = groupId
	group.groupName = name
	group.groupFormationDay = formationDay
	return *group
}

func Reconstruct(props GroupProps) group {
	return construct(
		props.groupId,
		props.groupName,
		props.groupFormationDay,
	)
}

func Form(
	name groupName,
	formatinDay groupFormationDay,
) group {
	return construct(
		GenerateGroupId(),
		name,
		formatinDay,
	)
}
