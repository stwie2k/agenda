package entity

// Meeting :

import (
 	"strings"
)
type Meeting struct {
	Sponsor string
	Participators []string
	StartDate, EndDate Date
	Title string
}
func (mMeeting Meeting) init(tSponsor string, tParticipators []string, tStartDate Date, tEndDate Date,tTitle string) {
	mMeeting.Sponsor= tSponsor
	mMeeting.SetParticipator(tParticipators)
	mMeeting.StartDate.CopyDate(tStartDate)
	mMeeting.EndDate.CopyDate(tEndDate)
	mMeeting.Title= tTitle
}

func (mMeeting Meeting) CopyMeeting (tMeeting Meeting) {
	mMeeting.Sponsor= tMeeting.Sponsor
	mMeeting.SetParticipator(tMeeting.Participators)
	mMeeting.StartDate.CopyDate(tMeeting.StartDate)
	mMeeting.EndDate.CopyDate(tMeeting.EndDate)
	mMeeting.Title= tMeeting.Title
}

func (mMeeting Meeting) GetSponsor() string {
	return mMeeting.Sponsor
}


func (mMeeting Meeting) SetSponsor(tSponsor string) {
	mMeeting.Sponsor = tSponsor
}


func (mMeeting Meeting) GetParticipator() []string {
    return mMeeting.Participators
}

func (mMeeting Meeting) SetParticipator(tParticipators []string) {
	var length= len(tParticipators)
	for i := 0; i < length; i++ {
		mMeeting.Participators[i]= tParticipators[i]
	}
}

func (mMeeting Meeting) GetStartDate() Date {
	return mMeeting.StartDate
}

func (mMeeting Meeting) SetStartDate(tStartTime Date) {
    mMeeting.StartDate.CopyDate(tStartTime)
}

func (mMeeting Meeting) GetEndDate() Date {
	return mMeeting.EndDate
}

func (mMeeting Meeting) SetEndDate(t_endTime Date) {
	mMeeting.EndDate.CopyDate(t_endTime)
}

func (mMeeting Meeting) GetTitle() string {
	return mMeeting.Title
}

func (mMeeting Meeting) SetTitle(t_title string) {
	mMeeting.Title = t_title
}

func (mMeeting Meeting) IsParticipator(t_username string) bool {
  var i int
	for i= 0; i< len(mMeeting.Participators); i++ {
		if strings.EqualFold(mMeeting.Participators[i], t_username)== true {
	    	return true
		}
	}
	return false
}

func (mMeeting *Meeting) DeleteParticipator(t_username string) {
	var i int
	tl := len(mMeeting.Participators)
	for i= 0; i< tl; i++ {
		if strings.EqualFold(mMeeting.Participators[i], t_username)== true {
	    	mMeeting.Participators = append(mMeeting.Participators[:i], mMeeting.Participators[i+1:]...)
			break
		}
	}
}

func (mMeeting *Meeting) AddParticipator(t_username string) bool {
	if strings.EqualFold(mMeeting.Sponsor, t_username) || mMeeting.IsParticipator(t_username) {
		return false
	}
	mMeeting.Participators = append(mMeeting.Participators,t_username)
	return true
}