package game

import (
	"image"
	"github.com/lei-cao/programming/code/algoman/pkg/ui"
	"github.com/hajimehoshi/ebiten"
	"bytes"
	"log"
	"github.com/lei-cao/programming/code/algoman/resources/images"
)

var (
	forwardOffImg      *ebiten.Image
	forwardOnImg       *ebiten.Image
	backwardOffImg     *ebiten.Image
	backwardOnImg      *ebiten.Image
	stepForwardOffImg  *ebiten.Image
	stepForwardOnImg   *ebiten.Image
	stepBackwardOffImg *ebiten.Image
	stepBackwardOnImg  *ebiten.Image
	redoOffImg            *ebiten.Image
	redoOnImg            *ebiten.Image
)

func init() {
	forwardOffImg = imgFromByte(images.FORWARD_OFF_png)
	forwardOnImg= imgFromByte(images.FORWARD_ON_png)
	backwardOffImg= imgFromByte(images.BACKWARD_OFF_png)
	backwardOnImg= imgFromByte(images.BACKWARD_ON_png)
	stepForwardOffImg= imgFromByte(images.STEP_FORWARD_OFF_png)
	stepForwardOnImg= imgFromByte(images.STEP_FORWARD_ON_png)
	stepBackwardOffImg= imgFromByte(images.STEP_BACKWARD_OFF_png)
	stepBackwardOnImg= imgFromByte(images.STEP_BACKWARD_ON_png)
	redoOffImg = imgFromByte(images.REDO_OFF_png)
	redoOnImg = imgFromByte(images.REDO_ON_png)
}

func imgFromByte(imgByte []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImg, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return ebitenImg
}

func NewController() *Controller {
	c := new(Controller)
	c.PlayToggle = ui.NewImageToggle(image.Rect(10, 250, 40, 280))
	c.NextStepBtn = ui.NewImageButton(image.Rect(50, 250, 80, 280), stepForwardOffImg, stepForwardOnImg)
	c.SpeedDownBtn = ui.NewImageButton(image.Rect(90, 250, 120, 280), backwardOffImg, backwardOnImg)
	c.SpeedUpBtn = ui.NewImageButton(image.Rect(130, 250, 160, 280), forwardOffImg, forwardOnImg)
	c.RestartBtn = ui.NewImageButton(image.Rect(170, 250, 200, 280), redoOffImg, redoOnImg)
	return c
}

type Controller struct {
	PlayToggle   *ui.ImageToggle
	NextStepBtn  *ui.ImageButton
	SpeedUpBtn   *ui.ImageButton
	SpeedDownBtn *ui.ImageButton
	RestartBtn   *ui.ImageButton
}

func (c *Controller) Update() {
	c.PlayToggle.Update()
	c.NextStepBtn.Update()
	c.SpeedDownBtn.Update()
	c.SpeedUpBtn.Update()
	c.RestartBtn.Update()
}

func (c *Controller) Draw(image *ebiten.Image) {
	c.PlayToggle.Draw(image)
	c.NextStepBtn.Draw(image)
	c.SpeedDownBtn.Draw(image)
	c.SpeedUpBtn.Draw(image)
	c.RestartBtn.Draw(image)
}
