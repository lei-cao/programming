package defaults

import "github.com/hajimehoshi/ebiten"

var DeviceScale = int(ebiten.DeviceScaleFactor())

var ImageBtnWidth = 32 * DeviceScale
var ImageBtnHeight = 36 * DeviceScale

var ButtonPadding = 10 * DeviceScale
var ButtonMinHeight = 36 * DeviceScale
var ButtonMinWidth = 36 * DeviceScale
