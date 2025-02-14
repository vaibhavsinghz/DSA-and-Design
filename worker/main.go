package main

import "fmt"

func main() {
	wp := NewWorkerPool(2)
	wp.Start()

	wp.AddTask(NewMeetingScheduler("Interview", "MAKALU", "interview@gmail.com"))
	wp.AddTask(NewMeetingScheduler("Loyalty V2 Flow", "K2", "vaibhav@dotpe.in"))
	wp.AddTask(NewSalaryProcessor("Kamlesh", 95000))
	wp.AddTask(NewMeetingScheduler("Postpaid Flow", "Olympus", "deepak@dotpe.in"))
	wp.AddTask(NewSalaryProcessor("Nandan", 145000))
	wp.AddTask(NewMeetingScheduler("Bifrost Flow Revamp", "Kafka", "nandan@dotpe.i"))
	wp.AddTask(NewMeetingScheduler("Merchant Panel Revamp", "Kailash", "discuss@dotpe.in"))
	wp.AddTask(NewSalaryProcessor("Vaibhav", 45000))
	wp.AddTask(NewSalaryProcessor("Tikesh", 115000))
	wp.AddTask(NewSalaryProcessor("Deepak", 245000))

	wp.Stop()
	fmt.Println("All tasks has been finished")
}
