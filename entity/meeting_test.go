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
		// TODO: Add test cases.
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

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mMeeting.GetSponsor(); got != tt.want {
				t.Errorf("Meeting.GetSponsor() = %v, want %v", got, tt.want)
			}
		})
	}
}
