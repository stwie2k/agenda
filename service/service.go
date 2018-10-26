package service

import (
 
  "fmt"
   "agenda/entity"
)
func StartAgenda() bool {
	
	return true
}
func QuitAgenda()bool {
	return false
}
func UserRegister(username string, password string, email string, phone string) (bool, error) {
	user := entity.QueryUser(func (u *entity.User) bool {
		return u.Name == username
	})
	if len(user) == 1 {
		fmt.Println("User Register: Already exist username")
		return false, nil
}
	entity.CreateUser(entity.User{username, password, email, phone})
	
	return true, nil
}