package payload

import "testing"

func TestNilIfEmpty(t *testing.T) {
	if NilIfEmpty("") != nil {
		t.Fatal("expected nil")
	}
	if NilIfEmpty("x") != "x" {
		t.Fatal("expected string")
	}
}

func TestSlugify(t *testing.T) {
	if got := Slugify("Example Speaker", ""); got != "example_speaker" {
		t.Fatalf("got %q", got)
	}
	if got := Slugify("media-player-2", ""); got != "media_player_2" {
		t.Fatalf("got %q", got)
	}
	if got := Slugify("!!!", "host"); got != "host" {
		t.Fatalf("got %q", got)
	}
	if got := Slugify("", ""); got != "device" {
		t.Fatalf("got %q", got)
	}
}
