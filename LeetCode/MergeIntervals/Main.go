package MergeIntervals

import (
	"sort"

	. "Algorithms/LeetCode/internal/box/math"
)

func merge(intervals []Interval) (result []Interval) {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Sort(Intervals(intervals))
	result = []Interval{intervals[0]}
	for length, curr, i := len(intervals), 0, 1; i < length; i++ {
		if result[curr].End < intervals[i].Start {
			result = append(result, intervals[i])
			curr++
			continue
		}
		result[curr].End = MaxOfTwo(result[curr].End, intervals[i].End)
	}
	return
}

type Intervals []Interval

func (is Intervals) Len() int {
	return len(is)
}

func (is Intervals) Less(i, j int) bool {
	return is[i].Start < is[j].Start
}

func (is Intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type Interval struct {
	Start int
	End   int
}
