package group

type GroupProps struct {
	groupProps
}

type groupProps struct {
	groupId GroupId
	groupName
	groupFormationDay
}

func (p groupProps) GroupId() GroupId {
	return p.groupId
}

func (p groupProps) GroupName() groupName {
	return p.groupName
}

func (p groupProps) FormationDay() groupFormationDay {
	return p.groupFormationDay
}
