package entity

import (
	"reflect"
	"testing"
)

func TestDate_init(t *testing.T) {
	type args struct {
		tYear   int
		tMonth  int
		tDay    int
		tHour   int
		tMinute int
	}
	tests := []struct {
		name  string
		mDate Date
		args  args
	}{
		{"leap year", Date{1996, 12, 9, 12, 12}, args{1996, 12, 9, 12, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mDate.init(tt.args.tYear, tt.args.tMonth, tt.args.tDay, tt.args.tHour, tt.args.tMinute)
		})
	}
}

func TestStringToDate(t *testing.T) {
	type args struct {
		tDateString string
	}
	tests := []struct {
		name    string
		args    args
		want    Date
		wantErr bool
	}{
		{"2019-01-19/12:12", args{"2019-01-19/12:12"}, Date{2019, 1, 19, 12, 12}, false},
		{"Invalid date format", args{"20-01-19/12:12"}, Date{0, 0, 0, 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToDate(tt.args.tDateString)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateToString(t *testing.T) {
	type args struct {
		tDate Date
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1996-12-09", args{Date{1996, 12, 9, 12, 0}}, "1996-12-09/12:00", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateToString(tt.args.tDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DateToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_MoreThan(t *testing.T) {
	type args struct {
		tDate Date
	}
	tests := []struct {
		name  string
		mDate Date
		args  args
		want  bool
	}{
		{"more: 1 > 2", Date{2019,1,1,12,00}, args{Date{2019,2,1,12,00}}, false},
		{"less: 1999 < 2019", Date{1999,1,1,12,00}, args{Date{2019,2,1,12,00}}, false},
		{"equal: 1 = 1", Date{2019,2,1,12,00}, args{Date{2019,2,1,12,00}}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mDate.MoreThan(tt.args.tDate); got != tt.want {
				t.Errorf("Date.MoreThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
