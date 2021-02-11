package stack

type KthLargest struct {
	values []int
	length int
	k      int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{}
	kthLargest.k = k
	kthLargest.length = 0
	kthLargest.values = make([]int, 0, k)
	for i := range nums {
		kthLargest.Add(nums[i])
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if this.length < this.k {
		this.values = append(this.values, val)
		this.length++
		this.stackly()
	} else if val > this.values[0] {
		this.values[0], this.values[this.length-1] = this.values[this.length-1], this.values[0]
		this.values[this.length-1] = val
		this.stackly()
	}
	return this.values[0]
}

func (this *KthLargest) stackly() {
	length := len(this.values)
	for i := length / 2; i >= 0; i-- {
		left, right := i<<1+1, i<<1+2
		if left < length && this.values[left] <= this.values[i] {
			this.values[i], this.values[left] = this.values[left], this.values[i]
		}
		if right < length && this.values[right] <= this.values[i] {
			this.values[i], this.values[right] = this.values[right], this.values[i]
		}
	}
}
