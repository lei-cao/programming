package controller

func (c *Controller) BubbleSort() {
	for i := 0; i < c.Config.Size; i++ {
		for j := 0; j < c.Config.Size-1; j++ {
			if c.nums[j] > c.nums[j+1] {
				c.swap(j, j+1)
			} else {
				c.pass(j, j+1)
			}
		}
	}
}

func (c *Controller) SelectionSort() {
	for i := 0; i < c.Config.Size; i++ {
		for j := i + 1; j < c.Config.Size; j++ {
			if c.nums[i] > c.nums[j] {
				c.swap(i, j)
			} else {
				c.pass(i, j)
			}
		}
	}
}

func (c *Controller) InsertionSort() {
	for i := 0; i < c.Config.Size; i++ {
		temp := c.nums[i]
		for j := i - 1; j >= 0; j-- {
			if c.nums[j] > temp {
				c.swap(j+1, j)
			} else {
				c.pass(j+1, j)
			}
		}
	}
}

func (c *Controller) QuickSort() {
	c.quickSort(0, len(c.nums)-1)
}

func (c *Controller) quickSort(lo, hi int) {
	if lo < hi {
		p := c.partition(lo, hi)
		c.quickSort(lo, p-1)
		c.quickSort(p+1, hi)
	}
}

func (c *Controller) partition(lo, hi int) int {
	pivot := c.nums[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if c.nums[j] < pivot {
			i += 1
			c.swap(i, j)
		} else {
			c.pass(hi, j)
		}
	}
	c.swap(hi, i+1)
	return i + 1
}

func (c *Controller) pass(a, b int) {
	step := &Step{A: a, B: b}
	c.LastStep.Next = step
	c.LastStep = step
}

func (c *Controller) swap(a, b int) {
	step := &Step{A: a, B: b, DoSwap: true}
	c.nums[a], c.nums[b] = c.nums[b], c.nums[a]
	c.LastStep.Next = step
	c.LastStep = step
}
