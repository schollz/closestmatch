package cmclient

import (
	"fmt"
	"testing"
)

var testingServer = "http://localhost:8051"

func TestClosest(t *testing.T) {
	conn, _ := Open(testingServer)
	match, err := conn.Closest("The War of the Worlds by H.G. Wells")
	if err != nil {
		t.Error(err)
	}
	if match != "The Time Machine/The War of the Worlds by H.G. Wells" {
		t.Error(match)
	}
}

func TestClosestN(t *testing.T) {
	conn, _ := Open(testingServer)
	matches, err := conn.ClosestN("The War of the Worlds by H.G. Wells", 10)
	if err != nil {
		t.Error(err)
	}
	if len(matches) != 10 {
		t.Errorf("Got %d", len(matches))
	}
	fmt.Println(matches)
}
