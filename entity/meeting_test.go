package entity

import (
	"testing"
)

func TestMeeting_init(t *testing.T) {
	type args struct {
		tSponsor       string
		tParticipators []string
		tStartDate     Date
		tEndDate       Date
		tTitle         string
	}
	tests := []struct {
		name     string
		mMeeting Meeting
		args     args
	}{
		{"create meeting test",
			Meeting{"Steve", []string{"Steve", "gyakkun"}, Date{1996, 12, 9, 12, 0}, Date{2019, 1, 1, 12, 00}, "Create meeting test"},
			args{"Steve", []string{"Steve", "gyakkun"}, Date{1996, 12, 9, 12, 0}, Date{2019, 1, 1, 12, 00}, "Create meeting test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mMeeting.init(tt.args.tSponsor, tt.args.tParticipators, tt.args.tStartDate, tt.args.tEndDate, tt.args.tTitle)
		})
	}
}

func TestMeeting_GetSponsor(t *testing.T) {
	tests := []struct {
		name     string
		mMeeting Meeting
		want     string
	}{
		{
			"Get sponsor test",
			Meeting{"Steve", []string{"Steve", "gyakkun"}, Date{1996, 12, 9, 12, 0}, Date{2019, 1, 1, 12, 00}, "Create meeting test"},
			"Steve",
		},
		{
			"Get Chinese sponsor test",
			Meeting{"大佬", []string{"Steve", "gyakkun"}, Date{1996, 12, 9, 12, 0}, Date{2019, 1, 1, 12, 00}, "Create meeting test"},
			"大佬",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mMeeting.GetSponsor(); got != tt.want {
				t.Errorf("Meeting.GetSponsor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeeting_IsParticipator(t *testing.T) {
	type args struct {
		t_username string
	}
	tests := []struct {
		name     string
		mMeeting Meeting
		args     args
		want     bool
	}{
		{"Check if dalao is a participator",
			Meeting{"大佬", []string{"大佬", "gyakkun"}, Date{1996, 12, 9, 12, 0}, Date{2019, 1, 1, 12, 00}, "Create meeting test"},
			args{"大佬"},
			true,
		},
		{"Check if dalao is a participator",
			Meeting{"大佬", []string{"Steve", "gyakkun"}, Date{1996, 12, 9, 12, 0}, Date{2019, 1, 1, 12, 00}, "Create meeting test"},
			args{"萌新"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mMeeting.IsParticipator(tt.args.t_username); got != tt.want {
				t.Errorf("Meeting.IsParticipator() = %v, want %v", got, tt.want)
			}
		})
	}
}
