package Elevator

import (
	"math"
	"sort"
)

func NewMission() mission {
	return make(map[int]int)
}

type mission map[int]int

func (m mission) Add(level, num int) {
	m[level] = num
}

func (m mission) BestLevels(levelNum int) ([]int, float64) {
	if levelNum > len(m) {
		levelNum = len(m)
	}
	type levelInfo struct {
		id   int
		cost float64
	}
	infoWithLevel := make(map[int]*levelInfo)
	for level, cost := range m.allCost() {
		infoWithLevel[level] = &levelInfo{
			id:   1,
			cost: cost,
		}
	}
	missionWithID := make(map[int]mission)
	for i := 0; i < levelNum; i++ {
		var maxCostLevel int
		var maxCost float64
		for level, info := range infoWithLevel {
			if info.cost > maxCost {
				maxCostLevel = level
				maxCost = info.cost
			}
		}
		if maxCostLevel == 0 {
			break
		}
		for level, cost := range m.allCostByLevel(maxCostLevel) {
			info := infoWithLevel[level]
			if i == 0 || cost < info.cost {
				info.id = maxCostLevel
				info.cost = cost
			}
		}
		missionWithID[maxCostLevel] = NewMission()
	}
	for level, info := range infoWithLevel {
		missionWithID[info.id].Add(level, m[level])
	}
	levels := make([]int, 0, levelNum)
	var costSum float64
	for _, mission := range missionWithID {
		level, cost := mission.BestLevel()
		levels = append(levels, level)
		costSum += cost
	}
	sort.Ints(levels)
	return levels, costSum
}

func (m mission) BestLevel() (int, float64) {
	var underNum, currentNum, upNum int
	levels := make([]int, 0, len(m))
	for level, num := range m {
		upNum += num
		levels = append(levels, level)
	}
	sort.Ints(levels)
	var bestLevel int
	for _, level := range levels {
		underNum += currentNum
		currentNum = m[level]
		upNum -= currentNum
		if underNum+currentNum < upNum {
			continue
		}
		bestLevel = level
		break
	}
	var costSum float64
	costs := m.allCostByLevel(bestLevel)
	for _, costs := range costs {
		costSum += costs
	}
	return bestLevel, costSum
}

func (m mission) allCostByLevel(baseLevel int) map[int]float64 {
	costs := make(map[int]float64)
	for level, num := range m {
		costs[level] = math.Abs(float64(level-baseLevel)) * float64(num)
	}
	return costs
}

func (m mission) allCost() map[int]float64 {
	return m.allCostByLevel(1)
}
