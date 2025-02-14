package main

import (
	"fmt"
	"time"
)

type Task interface {
	Process()
}

type MeetingScheduler struct {
	Name, Room, Email string
}

func NewMeetingScheduler(name, room, email string) Task {
	return &MeetingScheduler{name, room, email}
}

func (m MeetingScheduler) Process() {
	fmt.Printf("Meeting %s is scheduled on %s via %s\n", m.Name, m.Room, m.Email)
	time.Sleep(2 * time.Second)
}

type Salary struct {
	Name   string
	Amount int
}

func NewSalaryProcessor(name string, amount int) Task {
	return &Salary{name, amount}
}

func (s Salary) Process() {
	fmt.Printf("Salary %d is credited to %s account\n", s.Amount, s.Name)
	time.Sleep(2 * time.Second)
}
