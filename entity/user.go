package entity

// User : class user with member name, pw, email and phone
type User struct {
	Name, Password, Email, Phone string
}

func (mUser User) init(tName, tPassword, tEmail, tPhone string) {
	mUser.Name = tName
	mUser.Password = tPassword
	mUser.Email = tEmail
	mUser.Phone = tPhone
}

func (mUser User) CopyUser(t_user User) {
	mUser.Name = t_user.Name
	mUser.Password = t_user.Password
	mUser.Email = t_user.Email
	mUser.Phone = t_user.Phone
}

func (mUser User) GetName() string {
	return mUser.Name
}

func (mUser User) SetName(tName string) {
	mUser.Name = tName
}

func (mUser User) GetPassword() string {
	return mUser.Password
}

func (mUser User) SetPassword(tPassword string) {
	mUser.Password = tPassword
}

func (mUser User) GetEmail() string {
	return mUser.Email
}

func (mUser User) SetEmail(tEmail string) {
	mUser.Email = tEmail
}

func (mUser User) GetPhone() string {
	return mUser.Phone
}

func (mUser User) SetPhone(tPhone string) {
	mUser.Phone = tPhone
}

func (mUser User) IsSameUser(tUser User) bool {
	if mUser.GetEmail() == tUser.GetEmail() &&
		mUser.GetEmail() == tUser.GetEmail() &&
		mUser.GetName() == tUser.GetName() &&
		mUser.GetPassword() == tUser.GetPassword() &&
		mUser.GetPhone() == tUser.GetPhone() {

		return true
	}
	return false
}
