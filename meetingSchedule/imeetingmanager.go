package meetingSchedule

import "time"

type IMeetingManager interface {
	BookMeeting(startTime time.Time, duration time.Duration, user []User) (*Meeting, error)
}
