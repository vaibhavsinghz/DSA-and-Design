package meetingSchedule

import "time"

type Meeting struct {
	ID                 int
	StartTime, EndTime time.Time
	Users              []User
}

type MeetingRoom struct {
	ID, Capacity int
	Location     string
}
