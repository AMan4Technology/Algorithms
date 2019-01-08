package KthLargestElementInAStream

func Constructor(k int, nums []int) KthLargest {
	kL := KthLargest{
		minHeap: make([]int, 0, k),
	}
	for _, val := range nums {
		kL.Add(val)
	}
	return kL
}

type KthLargest struct {
	minHeap []int
}

func (k *KthLargest) Add(val int) int {
	if len(k.minHeap) < cap(k.minHeap) {
		k.minHeap = append(k.minHeap, val)
		k.up(len(k.minHeap) - 1)
	} else if val > k.minHeap[0] {
		k.minHeap[0] = val
		k.down(0)
	}
	return k.minHeap[0]
}

func (k *KthLargest) down(i int) {
	min := i
	for {
		if 2*i+1 < len(k.minHeap) && k.minHeap[min] > k.minHeap[2*i+1] {
			min = 2*i + 1
		}
		if 2*i+2 < len(k.minHeap) && k.minHeap[min] > k.minHeap[2*i+2] {
			min = 2*i + 2
		}
		if min == i {
			break
		}
		k.minHeap[i], k.minHeap[min] = k.minHeap[min], k.minHeap[i]
		i = min
	}
}

func (k *KthLargest) up(i int) {
	for k.minHeap[i] < k.minHeap[(i-1)/2] {
		k.minHeap[i], k.minHeap[(i-1)/2] = k.minHeap[(i-1)/2], k.minHeap[i]
		i = (i - 1) / 2
	}
}
