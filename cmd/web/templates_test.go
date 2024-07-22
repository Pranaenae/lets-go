package main

import (
	"testing"
	"time"

	"github.com/Pranaenae/lets-go/internal/assert"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 7, 22, 10, 48, 30, 0, time.UTC),
			want: "22 Jul 2024 at 10:48",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 7, 22, 10, 48, 30, 0, time.FixedZone("CET", 1*60*60)),
			want: "22 Jul 2024 at 09:48",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
