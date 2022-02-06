package group

import "time"

type groupFormationDay time.Time

func (f groupFormationDay) Value() time.Time {
	return time.Time(f)
}

func (f groupFormationDay) Format() string {
	return f.Value().Format("2006/01/02")
}

func NewGroupFormationDay(value time.Time) (groupFormationDay, error) {
	// TODO validate
	return groupFormationDay(value), nil
}
