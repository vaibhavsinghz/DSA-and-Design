package main

import (
	"Self/JIRA/models"
	"fmt"
)

func main() {
	jira := models.InitJIRA()
	devopsBoardID, err := jira.AddBoard("DEVOPS", true, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(devopsBoardID)
}
