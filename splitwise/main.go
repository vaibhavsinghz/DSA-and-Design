package main

import (
	"Self/splitwise/models/splitStrategy"
	"Self/splitwise/service"
)

func main() {
	gs := service.NewGroupService()
	us := service.NewUserService()

	vID := us.RegisterUser("Vaibhav")
	tID := us.RegisterUser("Tikesh")
	kID := us.RegisterUser("Kamlesh")

	h768 := gs.CreateGroup("H768", []int{vID, tID, kID})

	splitMap := make(map[int]float64)
	splitMap[vID] = 50
	splitMap[tID] = 25
	splitMap[kID] = 25

	gs.AddExpense(h768, "Electricity", 4999, kID, splitStrategy.PERCENT, splitStrategy.NewSplitList(splitMap))

	gs.AddExpense(h768, "Rent", 33000.0, vID, splitStrategy.EQUAL, nil) //user have to pass with whome we need to split equally

	gs.GetUserBalanceRecord(vID)
	gs.GetUserBalanceRecord(kID)
	//gs.CreateGroup("H768")
}
