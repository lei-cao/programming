# Algorithm Visualization

I made a very simple canvas app for visualising some basic sorting algorithms by [GopherJS](https://github.com/gopherjs/gopherjs)

```go
func (s *BubbleSort) Sort(nums []int) {
	obj := createCanvas("bubbleSort", len(nums))

	draw(nums, 0, 0, obj)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			draw(nums, i, j, obj)
			time.Sleep(1 * time.Millisecond)
		}
	}
}
```

Then I just loaded it in a Vue Component. Hacky.. But works now.

```html
<ACanvas/>
```

<ACanvas/>
