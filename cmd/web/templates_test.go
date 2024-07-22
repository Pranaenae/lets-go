package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2024, 7, 22, 10, 48, 30, 0, time.UTC)

	hd := humanDate(tm)

	if hd != "22 Jul 2024 at 10:48" {
		t.Errorf("got %q; want %q", hd, "22 July 2024 at 10:48")
	}
}
