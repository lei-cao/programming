package controller

func (c *Controller) BubbleSort() {

	for i := 0; i < c.Config.Size; i++ {
		for j := 0; j < c.Config.Size-1; j++ {
			step := &Step{}
			step.A = j
			step.B = j + 1
			if c.nums[j] > c.nums[j+1] {
				c.nums[j], c.nums[j+1] = c.nums[j+1], c.nums[j]
				step.DoSwap = true
			}
			c.LastStep.Next = step
			c.LastStep = step
		}
	}
}

func (c *Controller) SelectionSort() {

	for i := 0; i < c.Config.Size; i++ {
		for j := i + 1; j < len(c.nums); j++ {
			step := &Step{}
			step.A = i
			step.B = j
			if c.nums[i] > c.nums[j] {
				c.nums[i], c.nums[j] = c.nums[j], c.nums[i]
				step.DoSwap = true
			}
			c.LastStep.Next = step
			c.LastStep = step
		}
	}
}