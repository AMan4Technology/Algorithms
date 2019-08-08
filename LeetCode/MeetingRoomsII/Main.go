package MeetingRoomsII

import (
	"sort"
)

func minMeetingRooms(intervals []Interval) int {
	length := len(intervals)
	if length < 2 {
		return length
	}
	sort.Sort(Intervals(intervals))
	var heap MinHeap
	heap.Push(intervals[0].End)
	for i := 1; i < length; i++ {
		if intervals[i].Start >= heap.Top() {
			heap.Update(intervals[i].End)
			continue
		}
		heap.Push(intervals[i].End)
	}
	return heap.length
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

type MinHeap struct {
	length int
	data   []int
}

func (h *MinHeap) Push(value int) {
	h.length++
	h.data = append(h.data, value)
	h.up()
}

func (h MinHeap) Update(value int) {
	if h.length == 0 {
		return
	}
	h.data[0] = value
	h.down()
}

func (h MinHeap) Top() int {
	if h.length == 0 {
		return 0
	}
	return h.data[0]
}

func (h MinHeap) up() {
	for curr := h.length - 1; h.data[(curr-1)/2] > h.data[curr]; curr = (curr - 1) / 2 {
		h.data[(curr-1)/2], h.data[curr] = h.data[curr], h.data[(curr-1)/2]
	}
}

func (h MinHeap) down() {
	for curr, min := 0, 0; ; curr = min {
		if 2*curr+1 < h.length && h.data[2*curr+1] < h.data[min] {
			min = 2*curr + 1
		}
		if 2*curr+2 < h.length && h.data[2*curr+2] < h.data[min] {
			min = 2*curr + 2
		}
		if min == curr {
			break
		}
		h.data[min], h.data[curr] = h.data[curr], h.data[min]
	}
}
