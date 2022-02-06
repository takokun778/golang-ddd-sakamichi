package main

import (
	"fmt"
	"time"

	"sakamichi/domain/group"
	"sakamichi/domain/member"
)

func main() {
	keyakizakaName, _ := group.NewGroupNama("欅坂46")
	keyakizakaFormationDay, _ := group.NewGroupFormationDay(time.Date(2015, 8, 21, 0, 0, 0, 0, time.Local))
	keyakizaka46 := group.Form(keyakizakaName, keyakizakaFormationDay)

	fmt.Println(keyakizaka46.GroupId().Value())
	fmt.Println(keyakizaka46.GroupName().Value())
	fmt.Println(keyakizaka46.FormationDay().Format())

	memberLastName, _ := member.NewMemberName("大園")
	memberFirstName, _ := member.NewMemberName("玲")
	memberBirthday, _ := member.NewMemberBirthday(time.Date(2000, 4, 18, 0, 0, 0, 0, time.Local))
	memberBlood, _ := member.NewMemberBlood("A")
	memberJoiningDay, _ := member.NewMemberJoiningDay(time.Date(2020, 2, 16, 0, 0, 0, 0, time.Local))
	rei := member.Debut(
		keyakizaka46.GroupId(),
		memberFirstName,
		memberLastName,
		memberBirthday,
		memberBlood,
		memberJoiningDay,
	)

	fmt.Println(rei.GroupId().Value())
	fmt.Println(rei.LastName().Value())
	fmt.Println(rei.FirstName().Value())
	fmt.Println(rei.Birthday().Format())
	fmt.Println(rei.JoiningDay().Format())

	sakurazakaName, _ := group.NewGroupNama("櫻坂46")
	sakurazakaFormationDay, _ := group.NewGroupFormationDay(time.Date(2020, 10, 14, 0, 0, 0, 0, time.Local))
	sakurazaka46 := group.Form(sakurazakaName, sakurazakaFormationDay)

	fmt.Println(sakurazaka46.GroupId().Value())
	fmt.Println(sakurazaka46.GroupName().Value())
	fmt.Println(sakurazaka46.FormationDay().Format())

	reichan := rei.Transfer(sakurazaka46.GroupId())

	fmt.Println(reichan.GroupId().Value())
	fmt.Println(reichan.LastName().Value())
	fmt.Println(reichan.FirstName().Value())
	fmt.Println(reichan.Birthday().Format())
	fmt.Println(reichan.JoiningDay().Format())
}
