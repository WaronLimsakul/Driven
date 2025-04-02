package tasks

import (
	"reflect"
	"testing"
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
)

// Helper function to compare time.Time values ignoring location
func timeEqual(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

func TestGetWeekRange(t *testing.T) {
	tests := []struct {
		name       string
		input      time.Time
		wantMonday time.Time
		wantSunday time.Time
	}{
		{
			name:       "Wednesday",
			input:      time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC), // Wed May 17 2023
			wantMonday: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			wantSunday: time.Date(2023, 5, 21, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Sunday",
			input:      time.Date(2023, 5, 21, 0, 0, 0, 0, time.UTC), // Sun May 21 2023
			wantMonday: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			wantSunday: time.Date(2023, 5, 21, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Monday",
			input:      time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), // Mon May 15 2023
			wantMonday: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			wantSunday: time.Date(2023, 5, 21, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMonday, gotSunday := GetWeekRange(tt.input)
			if !timeEqual(gotMonday, tt.wantMonday) {
				t.Errorf("GetWeekRange() gotMonday = %v, want %v", gotMonday, tt.wantMonday)
			}
			if !timeEqual(gotSunday, tt.wantSunday) {
				t.Errorf("GetWeekRange() gotSunday = %v, want %v", gotSunday, tt.wantSunday)
			}
		})
	}
}

func TestGroupTaskDate(t *testing.T) {
	mon := time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC) // Monday
	tue := mon.Add(24 * time.Hour)                      // Tuesday
	wed := tue.Add(24 * time.Hour)                      // Wednesday

	tasks := []database.Task{
		{Date: mon},
		{Date: tue},
		{Date: wed},
		{Date: mon}, // Another Monday task
	}

	got := GroupTaskDate(tasks)

	if len(got) != 7 {
		t.Fatalf("GroupTaskDate() returned slice with length %d, want 7", len(got))
	}

	// Check Monday (index 0)
	if len(got[0]) != 2 {
		t.Errorf("GroupTaskDate() got[0] has %d tasks, want 2", len(got[0]))
	}

	// Check Tuesday (index 1)
	if len(got[1]) != 1 {
		t.Errorf("GroupTaskDate() got[1] has %d tasks, want 1", len(got[1]))
	}

	// Check Wednesday (index 2)
	if len(got[2]) != 1 {
		t.Errorf("GroupTaskDate() got[2] has %d tasks, want 1", len(got[2]))
	}

	// Check other days should be empty
	for i := 3; i < 7; i++ {
		if len(got[i]) != 0 {
			t.Errorf("GroupTaskDate() got[%d] should be empty, but has %d tasks", i, len(got[i]))
		}
	}
}

func TestGetWeekDayNum(t *testing.T) {
	tests := []struct {
		name string
		day  time.Time
		want int
	}{
		{"Monday", time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), 0},
		{"Tuesday", time.Date(2023, 5, 16, 0, 0, 0, 0, time.UTC), 1},
		{"Wednesday", time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC), 2},
		{"Thursday", time.Date(2023, 5, 18, 0, 0, 0, 0, time.UTC), 3},
		{"Friday", time.Date(2023, 5, 19, 0, 0, 0, 0, time.UTC), 4},
		{"Saturday", time.Date(2023, 5, 20, 0, 0, 0, 0, time.UTC), 5},
		{"Sunday", time.Date(2023, 5, 21, 0, 0, 0, 0, time.UTC), 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeekDayNum(tt.day); got != tt.want {
				t.Errorf("GetWeekDayNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWeekDayStr(t *testing.T) {
	tests := []struct {
		name string
		day  time.Time
		want string
	}{
		{"Monday", time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), "Mon"},
		{"Tuesday", time.Date(2023, 5, 16, 0, 0, 0, 0, time.UTC), "Tue"},
		{"Wednesday", time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC), "Wed"},
		{"Thursday", time.Date(2023, 5, 18, 0, 0, 0, 0, time.UTC), "Thu"},
		{"Friday", time.Date(2023, 5, 19, 0, 0, 0, 0, time.UTC), "Fri"},
		{"Saturday", time.Date(2023, 5, 20, 0, 0, 0, 0, time.UTC), "Sat"},
		{"Sunday", time.Date(2023, 5, 21, 0, 0, 0, 0, time.UTC), "Sun"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeekDayStr(tt.day); got != tt.want {
				t.Errorf("GetWeekDayStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpanWeekDateByFormat(t *testing.T) {
	monday := time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC) // Monday
	want := []string{"15/05", "16/05", "17/05", "18/05", "19/05", "20/05", "21/05"}

	format := func(t time.Time) string {
		return t.Format("02/01")
	}

	got := SpanWeekDateByFormat(monday, format)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SpanWeekDateByFormat() = %v, want %v", got, want)
	}
}

func TestSpanWeekDate(t *testing.T) {
	monday := time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC) // Monday
	want := []time.Time{
		monday,
		monday.Add(24 * time.Hour),
		monday.Add(48 * time.Hour),
		monday.Add(72 * time.Hour),
		monday.Add(96 * time.Hour),
		monday.Add(120 * time.Hour),
		monday.Add(144 * time.Hour),
	}

	got := SpanWeekDate(monday)

	if len(got) != len(want) {
		t.Fatalf("SpanWeekDate() returned slice with length %d, want %d", len(got), len(want))
	}

	for i := range want {
		if !timeEqual(got[i], want[i]) {
			t.Errorf("SpanWeekDate()[%d] = %v, want %v", i, got[i], want[i])
		}
	}
}
