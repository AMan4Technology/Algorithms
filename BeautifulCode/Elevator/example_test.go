package Elevator_test

import (
	"Algorithms/BeautifulCode/elevatorlcode/elevator"
	"fmt"
)

func Example() {
	mission := elevator.NewMission()
	mission.Add(1, 10)
	mission.Add(3, 2)
	mission.Add(4, 2)
	mission.Add(6, 1)
	mission.Add(7, 3)
	mission.Add(8, 4)
	fmt.Println(mission.BestLevel())
	fmt.Println(mission.BestLevels(2))

	// Output:
	// 3 57
	// [1 7] 15
}
