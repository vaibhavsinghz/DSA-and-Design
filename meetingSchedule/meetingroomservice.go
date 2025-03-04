package meetingSchedule

import (
	"sync"
	"time"
)

type MeetingRoomManager struct {
	MeetingRooms  map[int]MeetingRoom //map of room id to meeting room
	RoomMeetings  map[int][]Meeting   //map of room id to list of meetings
	roomIDCounter int
	mu            sync.Mutex
}

func NewMeetingRoomManager() IMeetingRoomService {
	return &MeetingRoomManager{
		MeetingRooms: map[int]MeetingRoom{},
		RoomMeetings: map[int][]Meeting{},
	}
}

func (mrm *MeetingRoomManager) GetAvailableMeetingRooms(capacity int, startTime, endTime time.Time) ([]MeetingRoom, error) {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	possibleMeetingRoomIDs := []int{}
	for _, room := range mrm.MeetingRooms {
		if room.Capacity == capacity {
			possibleMeetingRoomIDs = append(possibleMeetingRoomIDs, room.ID)
		}
	}

	var meetingRooms []MeetingRoom
	for _, possibleMeetingRoomID := range possibleMeetingRoomIDs {
		isAvailable := true
		for _, scheduledMeeting := range mrm.RoomMeetings[possibleMeetingRoomID] {
			if (scheduledMeeting.StartTime.After(startTime) && scheduledMeeting.StartTime.Before(endTime)) || (scheduledMeeting.EndTime.After(startTime) && scheduledMeeting.EndTime.Before(endTime)) {
				isAvailable = false
				break
			}
		}
		if isAvailable {
			meetingRooms = append(meetingRooms, mrm.MeetingRooms[possibleMeetingRoomID])
		}
	}

	return meetingRooms, nil
}

func (mrm *MeetingRoomManager) BookRoom(meetingRoomID int, meeting Meeting) {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	mrm.RoomMeetings[meetingRoomID] = append(mrm.RoomMeetings[meetingRoomID], meeting)
}

func (mrm *MeetingRoomManager) AddRoom(capacity int, location string) {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	mrm.roomIDCounter++

	newMeetingRoom := MeetingRoom{
		ID:       mrm.roomIDCounter,
		Location: location,
		Capacity: capacity,
	}

	mrm.MeetingRooms[newMeetingRoom.ID] = newMeetingRoom
}
