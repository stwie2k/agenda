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

/**
* @brief set the sponsor of a meeting
* @param  the new sponsor string
*/
func (mMeeting Meeting) SetSponsor(tSponsor string) {
	mMeeting.Sponsor = tSponsor
}

/**
* @brief  get the participator of a meeting
* @return return a string indicate participator
*/
func (mMeeting Meeting) GetParticipator() []string {
    return mMeeting.Participators
}

/**
*   @brief set the new participator of a meeting
*   @param the new participator string
*/

func (mMeeting Meeting) SetParticipator(tParticipators []string) {
	var length= len(tParticipators)
	for i := 0; i < length; i++ {
		mMeeting.Participators[i]= tParticipators[i]
	}
}

/**
* @brief get the startDate of a meeting
* @return return a string indicate startDate
*/
func (mMeeting Meeting) GetStartDate() Date {
	return mMeeting.StartDate
}

/**
* @brief  set the startDate of a meeting
* @param  the new startdate of a meeting
*/
func (mMeeting Meeting) SetStartDate(tStartTime Date) {
    mMeeting.StartDate.CopyDate(tStartTime)
}

/**
* @brief get the endDate of a meeting
* @return a date indicate the endDate
*/
func (mMeeting Meeting) GetEndDate() Date {
	return mMeeting.EndDate
}

/**
* @brief  set the endDate of a meeting
* @param  the new enddate of a meeting
*/
func (mMeeting Meeting) SetEndDate(t_endTime Date) {
	mMeeting.EndDate.CopyDate(t_endTime)
}

/**
* @brief get the title of a meeting
* @return a date title the endDate
*/
func (mMeeting Meeting) GetTitle() string {
	return mMeeting.Title
}

/**
* @brief  set the title of a meeting
* @param  the new title of a meeting
*/
func (mMeeting Meeting) SetTitle(t_title string) {
	mMeeting.Title = t_title
}

/**
* @brief check if the user take part in this meeting
* @param t_username the source username
* @return if the user take part in this meeting
*/
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