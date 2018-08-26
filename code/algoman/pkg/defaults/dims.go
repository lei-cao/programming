package defaults

import (
	"github.com/hajimehoshi/ebiten"
)

var DeviceScale = int(ebiten.DeviceScaleFactor())

var (
	ImageBtnWidth  = 32 * DeviceScale
	ImageBtnHeight = 36 * DeviceScale
)

var (
	ButtonPadding   = 10 * DeviceScale
	ButtonMinHeight = 36 * DeviceScale
	ButtonMinWidth  = 36 * DeviceScale
)

var (
	BarWidth      = 8 * DeviceScale
	BarHeightUnit = 5 * DeviceScale
	BarMargin     = 2 * DeviceScale
)

var (
	CircleRadius      = 11 * DeviceScale
)
