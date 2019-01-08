package MeetingRoomsII

import (
	"fmt"
	"testing"
)

func TestMinMeetingRooms(t *testing.T) {
	fmt.Println(minMeetingRooms(Intervals{
		{5, 10}, {0, 30}, {10, 15},
	}))
}
