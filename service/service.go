package service

import (
	"agenda/entity"
	"agenda/loghelper"
	"log"
)

var userInfoFile = "/src/agenda/data/curuser.txt"
var Log *log.Logger

type User entity.User
type Meeting entity.Meeting

func init() {
	Log = loghelper.Error
}

func StartAgenda() bool {
	return true
}

func QuitAgenda() bool {
	return false
}

//@param Username, password, email address, phone number
func UserRegister(_name string, password string, email string, phone string) (bool, error) {
	user := entity.QueryUser(func(u *entity.User) bool {
		return u.Name == _name
	})
	if len(user) == 1 {
		Log.Println("User Register: Already exist username")
		return false, nil
	}
	entity.CreateUser(&entity.User{_name, password, email, phone})

	return true, nil
}

//@param username
func DeleteUser(_name string) bool {
	entity.DeleteUser(func(u *entity.User) bool {
		return u.Name == _name
	})
	entity.UpdateMeeting(
		func(m *entity.Meeting) bool {
			return m.IsParticipator(_name)
		},
		func(m *entity.Meeting) {
			m.DeleteParticipator(_name)
		})
	entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == _name || len(m.GetParticipator()) == 0
	})
	if err := entity.Sync(); err != nil {
		Log.Println("Delete user failed: Syncing went wrong")
		return false
	}
	return UserLogout()
}

//@param Username and password
func UserLogin(_name string, _pw string) bool {
	userList := entity.QueryUser(func(_user *entity.User) bool {
		if _user.Name == _name && _user.Password == _pw {
			return true
		}
		return false
	})
	if len(userList) == 0 {
		Log.Println("Login failed: User not found")
		return false
	}
	if len(userList) > 1 {
		Log.Println("Login failed:Something went wrong, redundant users with same name")
		Log.Println("			  Please contact us!")
		return false
	}
	entity.SetCurUser(&userList[0])
	if err := entity.Sync(); err != nil {
		Log.Println("Login failed: Syncing went wrong")
		return false
	}
	return true
}

func UserLogout() bool {
	if err := entity.Logout(); err != nil {
		Log.Println("Logout failed: Something went wrong")
		return false
	}
	return true
}

//@return current user
func GetCurUser() (entity.User, bool) {
	if cu, err := entity.GetCurUser(); err != nil {
		//err is error
		return cu, false
	} else {
		return cu, true
	}
}

func ListAllUser() []entity.User {
	return entity.QueryUser(func(u *entity.User) bool {
		return true
	})
}

//@param Sponsor username, meeting title, start date, end date, participator names
func CreateMeeting(_name string, title string, startDate string, endDate string, participator []string) bool {
	for _, curPart := range participator {
		if _name == curPart {
			Log.Println("Create meeting failed: Sponsor can't be a participator")
			return false
		}
		l := entity.QueryUser(func(u *entity.User) bool {
			return u.Name == curPart
		})
		if len(l) == 0 {
			Log.Println("Create meeting failed: no such user: ", curPart)
			return false
		}
		dc := 0
		for _, j := range participator {
			if j == curPart {
				dc++
				if dc == 2 {
					Log.Println("Create meeting failed: Duplicate participator.Participators should be unique.")
					return false
				}
			}
		}
	}

	nSponsor := entity.QueryUser(func(u *entity.User) bool {
		return u.Name == _name
	})
	if len(nSponsor) == 0 {
		Log.Println("Create meeting failed: Sponsor ", _name, " not exist")
		return false
	}

	startTime, err := entity.StringToDate(startDate)
	if err != nil {
		Log.Println("Create meeting failed: Wrong with start date.")
		return false
	}
	endTime, err := entity.StringToDate(endDate)
	if err != nil {
		Log.Println("Create meeting failed: Wrong with end date.")
		return false
	}
	if endTime.LessThan(startTime) == true {
		Log.Println("Create meeting failed: Start time should be earlier than end time")
		return false
	}
	for _, p := range participator {
		l := entity.QueryMeeting(func(m *entity.Meeting) bool {
			if m.Sponsor == p || m.IsParticipator(p) {
				if m.StartDate.LessOrEqual(startTime) && m.EndDate.MoreThan(startTime) {
					return true
				}
				if m.StartDate.LessThan(endTime) && m.EndDate.GreateOrEqual(endTime) {
					return true
				}
				if m.StartDate.GreateOrEqual(startTime) && m.EndDate.LessOrEqual(endTime) {
					return true
				}
			}
			return false
		})
		if len(l) > 0 {
			Log.Println("Create meeting failed: ", p, " time conflict")
			return false
		}
	}

	//Compare meeting time
	//Including sponsor's meeting and participator's
	l := entity.QueryMeeting(func(m *entity.Meeting) bool {
		if m.Sponsor == _name || m.IsParticipator(_name) {
			if m.StartDate.LessOrEqual(startTime) && m.EndDate.MoreThan(startTime) {
				return true
			}
			if m.StartDate.LessThan(endTime) && m.EndDate.GreateOrEqual(endTime) {
				return true
			}
			if m.StartDate.GreateOrEqual(startTime) && m.EndDate.LessOrEqual(endTime) {
				return true
			}
		}
		return false
	})

	if len(l) > 0 {
		Log.Println("Create meeting failed: ", _name, " time conflict")
		return false
	}
	entity.CreateMeeting(&entity.Meeting{_name, participator, startTime, endTime, title})
	if err := entity.Sync(); err != nil {
		Log.Println("Create meeting failed: Syncing went wrong")
		return false
	}
	return true
}

//@param username and meeting title
func DeleteMeeting(_name string, title string) int {
	if err := entity.Sync(); err != nil {
		return false
	}
	return entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == _name && m.Title == title
	})
}

//@param Sponsor username and meeting start date,end date
func QueryMeeting(_name string, startDate string, endDate string) ([]entity.Meeting, bool) {
	var meetings []entity.Meeting

	startTime, err := entity.StringToDate(startDate)
	if err != nil {
		Log.Println("Meeting query failed: Wrong with start date")
		return meetings, false
	}

	endTime, err := entity.StringToDate(endDate)
	if err != nil {
		Log.Println("Meeting query failed: Wrong with end date")
		return meetings, false
	}

	if endTime.LessThan(startTime) == true {
		Log.Println("Meeting query failed: Start time should be earlier than end time")
		return meetings, false
	}

	m := entity.QueryMeeting(func(temp *entity.Meeting) bool {
		if temp.Sponsor == _name || temp.IsParticipator(_name) {
			if temp.StartDate.LessOrEqual(startTime) && temp.EndDate.MoreThan(startTime) {
				return true
			}
			if temp.StartDate.LessOrEqual(endTime) && temp.EndDate.GreateOrEqual(endTime) {
				return true
			}
			if temp.StartDate.GreateOrEqual(startTime) && temp.EndDate.LessOrEqual(endTime) {
				return true
			}
		}
		return false
	})
	return m, true
}

//@param Sponsor username and meeting title
func QuitMeeting(_name string, title string) bool {
	meetings := entity.QueryMeeting(func(m *entity.Meeting) bool {
		return m.Title == title && m.IsParticipator(_name) == true
	})
	if len(meetings) == 0 {
		Log.Println("Quit meeting failed: No such meeting")
		return false
	}

	entity.UpdateMeeting(func(m *entity.Meeting) bool {
		return m.IsParticipator(_name) == true && m.Title == title
	}, func(m *entity.Meeting) {
		m.DeleteParticipator(_name)
	})

	entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return len(m.GetParticipator()) == 0
	})
	
	if err := entity.Sync(); err != nil {
		return false
	}
	
	return true
}

func ClearMeeting(_name string) (int, bool) {
	cm := entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == _name
	})
	if err := entity.Sync(); err != nil {
		Log.Println("Clear meeting failed: Delete operation went wrong")
		return cm, false
	} else {
		return cm, true
	}
}

func AddMeetingParticipator(_name string, title string, participators []string) bool {

	meetings := entity.UpdateMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == _name && m.Title == title
	}, func(m *entity.Meeting) {
		for _, p := range participators {
			m.AddParticipator(p)
		}
	})
	if meetings == 0 {
		Log.Println("Add participator failed: No such meeting")
		return false
	}

	for _, p := range participators {
		uc := entity.QueryUser(func(u *entity.User) bool {
			return u.Name == p
		})
		if len(uc) == 0 {
			Log.Println("Add participator failed: No such user: ", p)
			return false
		}
		qm := entity.QueryMeeting(func(m *entity.Meeting) bool {
			return m.Sponsor == _name && m.Title == title && m.IsParticipator(p)
		})
		if len(qm) != 0 {
			Log.Println("Add participator failed: User ", p, " is already in meeting")
			return false
		}
	}

	if err := entity.Sync(); err != nil {
		return false
	}
	return true
}

func RemoveMeetingParticipator(_name string, title string, participators []string) bool {
	for _, p := range participators {
		uc := entity.QueryUser(func(u *entity.User) bool {
			return u.Name == p
		})
		if len(uc) == 0 {
			Log.Println("Remove participator failed: No such user: ", p)
			return false
		}
		qm := entity.QueryMeeting(func(m *entity.Meeting) bool {
			return m.Sponsor == _name && m.Title == title && m.IsParticipator(p)
		})
		if len(qm) == 0 {
			Log.Println("Remove participator failed: ", p, " is not in Meeting")
			return false
		}
	}
	meetings := entity.UpdateMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == _name && m.Title == title
	}, func(m *entity.Meeting) {
		for _, p := range participators {
			m.DeleteParticipator(p)
		}
	})
	if meetings == 0 {
		Log.Println("Remove participator failed: no such meeting: ", title)
		return false
	}
	entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == _name && len(m.GetParticipator()) == 0
	})
	if err := entity.Sync(); err != nil {
		return false
	}
	return true
}
