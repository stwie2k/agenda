package entity

import (
	"testing"
)

func TestUser_init(t *testing.T) {
	type args struct {
		tName     string
		tPassword string
		tEmail    string
		tPhone    string
	}
	tests := []struct {
		name  string
		mUser User
		args  args
	}{
		{
			"init test",
			User{"a", "a", "a", "a"},
			args{"a", "a", "a", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mUser.init(tt.args.tName, tt.args.tPassword, tt.args.tEmail, tt.args.tPhone)
		})
	}
}

func TestUser_IsSameUser(t *testing.T) {
	type args struct {
		tUser User
	}
	tests := []struct {
		name  string
		mUser User
		args  args
		want  bool
	}{
		{
			"Is same user",
			User{"name", "phone", "email", "pw"},
			args{User{"naem", "phone", "email", "pw"}},
			false,			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mUser.IsSameUser(tt.args.tUser); got != tt.want {
				t.Errorf("User.IsSameUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
