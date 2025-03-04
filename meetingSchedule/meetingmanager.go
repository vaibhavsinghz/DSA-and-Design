package meetingSchedule

import (
	"errors"
	"sync"
	"time"
)

type MeetingManager struct {
	MeetingRoomService IMeetingRoomService
	Meetings           map[int]*Meeting
	MeetingRooms       map[int]*MeetingRoom
	meetingIDCounter   int
	mutex              sync.Mutex
}

func NewMeetingManager(meetingRoomService IMeetingRoomService) IMeetingManager {
	return &MeetingManager{
		MeetingRoomService: meetingRoomService,
		Meetings:           make(map[int]*Meeting),
		MeetingRooms:       make(map[int]*MeetingRoom),
	}
}

func (mm *MeetingManager) BookMeeting(startTime time.Time, duration time.Duration, user []User) (*Meeting, error) {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	capacity := len(user)
	endTime := startTime.Add(duration)

	availableRooms, err := mm.MeetingRoomService.GetAvailableMeetingRooms(capacity, startTime, endTime)
	if err != nil {
		return nil, err
	}
	if len(availableRooms) == 0 {
		return nil, errors.New("no available meeting rooms")
	}

	mm.meetingIDCounter++
	newMeeting := Meeting{
		ID:        mm.meetingIDCounter,
		StartTime: startTime,
		EndTime:   endTime,
		Users:     user,
	}

	mm.MeetingRoomService.BookRoom(availableRooms[0].ID, newMeeting)

	for _, user := range user {
		user.Notify(newMeeting)
	}

	return &newMeeting, nil
}
