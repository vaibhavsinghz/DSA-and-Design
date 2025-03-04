package meetingSchedule

import "time"

type IMeetingRoomService interface {
	GetAvailableMeetingRooms(capacity int, startTime, endTime time.Time) ([]MeetingRoom, error)
	BookRoom(meetingRoomID int, meeting Meeting)
	AddRoom(capacity int, location string)
}
