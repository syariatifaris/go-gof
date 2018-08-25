package main

import (
	"fmt"
	"errors"
)

type Leave struct{
	numOfDay int
}

func (l *Leave) GetNumOfDay() int{
	return l.numOfDay
}

type IApprover interface {
	SetNextApprover(approver IApprover)
	PerformApproval(Leave *Leave) error
}

type TeamLead struct{
	nextApprover IApprover
}

func (t *TeamLead) SetNextApprover(approver IApprover) {
	t.nextApprover = approver
}

func (t *TeamLead) PerformApproval(Leave *Leave) error {
	fmt.Println("approved by the teamlead")
	if t.nextApprover != nil{
		return t.nextApprover.PerformApproval(Leave)
	}
	return errors.New("no more approval needed")
}

type Manager struct{
	nextApprover IApprover
}

func (m *Manager) SetNextApprover(approver IApprover) {
	m.nextApprover = approver
}

func (m *Manager) PerformApproval(Leave *Leave) error {
	fmt.Println("approved by the manager")
	if m.nextApprover != nil{
		return m.nextApprover.PerformApproval(Leave)
	}
	return errors.New("no more approval needed")
}

type VPEngineering struct{}

func (*VPEngineering) SetNextApprover(approver IApprover) {
	return
}

func (*VPEngineering) PerformApproval(Leave *Leave) error {
	fmt.Println("approved by vp of engineering")
	return nil
}

func main(){
	teamLead := new(TeamLead)
	manager := new(Manager)
	vp := new(VPEngineering)

	teamLead.SetNextApprover(manager)
	manager.SetNextApprover(vp)

	teamLead.PerformApproval(&Leave{numOfDay:3})
}




