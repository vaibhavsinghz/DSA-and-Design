package meetingSchedule

import "fmt"

type User interface {
	Notify(meeting Meeting)
}

type user struct {
	Email string
}

func (ob *user) Notify(meeting Meeting) {
	fmt.Printf("Meeding %d has been scheduled from %v to %v\n", meeting.ID, meeting.StartTime, meeting.EndTime)
}
