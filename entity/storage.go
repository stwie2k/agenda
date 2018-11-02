
package entity

import (
	

)
var userlist []User
var meetinglist []Meeting

type uFilter func (*User) bool
type uSwitcher func (*User) 
type mFilter func (*Meeting) bool
type mSwitcher func (*Meeting) 

func CreateUser(_user User) {
	userlist = append(userlist,_user)
	
}
func QueryUser(filter uFilter) []User {
	var dy []User;
	for _, u := range userlist {
		if filter(&u) {
			dy = append(dy, u)
		}
	}
	return dy
}
func UpdateUser(filter uFilter, switcher uSwitcher) int {
	n := 0
	for _, u := range userlist {
		if filter(&u) {
			switcher(&u)
			n++
		}
	}
	return n
}
func CreateMeeting(_meeting Meeting){
   meetinglist=append(meetinglist,_meeting)

}

