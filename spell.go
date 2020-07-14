package main

import (
	"math"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

// fyne custom widget
type clockLayout struct {
	hour, minute, second     *canvas.Line
	hourDot, secondDot, face *canvas.Circle

	canvas fyne.CanvasObject
	stop   bool
}

func (c *clockLayout) rotate(hand fyne.CanvasObject, middle fyne.Position, facePosition float64, offset, length int) {
	rotation := math.Pi * 2 / 60 * facePosition
	x2 := int(float64(length) * math.Sin(rotation))
	y2 := int(float64(-length) * math.Cos(rotation))

	offX := 0
	offY := 0
	if offset > 0 {
		offX += int(float64(offset) * math.Sin(rotation))
		offY += int(float64(-offset) * math.Cos(rotation))
	}

	hand.Move(fyne.NewPos(middle.X+offX, middle.Y+offY))
	hand.Resize(fyne.NewSize(x2, y2))
}

func (c *clockLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	dotRadius := radius / 12
	smallDotRadius := dotRadius / 8

	stroke := float32(diameter) / 40
	midStroke := float32(diameter) / 90
	smallStroke := float32(diameter) / 200

	size = fyne.NewSize(diameter, diameter)
	middle := fyne.NewPos(size.Width/2, size.Height/2)
	topleft := fyne.NewPos(middle.X-radius, middle.Y-radius)

	c.face.Resize(size)
	c.face.Move(topleft)

	c.hour.StrokeWidth = stroke
	c.rotate(c.hour, middle, float64((time.Now().Hour()%12)*5)+(float64(time.Now().Minute())/12), dotRadius, radius/2)
	c.minute.StrokeWidth = midStroke
	c.rotate(c.minute, middle, float64(time.Now().Minute())+(float64(time.Now().Second())/60), dotRadius, int(float64(radius)*.9))
	c.second.StrokeWidth = smallStroke
	c.rotate(c.second, middle, float64(time.Now().Second()), 0, radius-3)

	c.hourDot.StrokeWidth = stroke
	c.hourDot.Resize(fyne.NewSize(dotRadius*2, dotRadius*2))
	c.hourDot.Move(fyne.NewPos(middle.X-dotRadius, middle.Y-dotRadius))
	c.secondDot.StrokeWidth = smallStroke
	c.secondDot.Resize(fyne.NewSize(smallDotRadius*2, smallDotRadius*2))
	c.secondDot.Move(fyne.NewPos(middle.X-smallDotRadius, middle.Y-smallDotRadius))
	c.face.StrokeWidth = smallStroke
}

// canvas layout으로서 MinSize가 작동함
func (c *clockLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(400, 400)
}

func (c *clockLayout) render() *fyne.Container {
	c.hourDot = &canvas.Circle{StrokeColor: theme.TextColor(), StrokeWidth: 5}
	c.secondDot = &canvas.Circle{StrokeColor: theme.PrimaryColor(), StrokeWidth: 3}
	c.face = &canvas.Circle{StrokeColor: theme.TextColor(), StrokeWidth: 1}

	c.hour = &canvas.Line{StrokeColor: theme.TextColor(), StrokeWidth: 5}
	c.minute = &canvas.Line{StrokeColor: theme.TextColor(), StrokeWidth: 3}
	c.second = &canvas.Line{StrokeColor: theme.PrimaryColor(), StrokeWidth: 1}

	container := fyne.NewContainer(c.hourDot, c.secondDot, c.face, c.hour, c.minute, c.second)
	container.Layout = c

	c.canvas = container
	return container
}

func (c *clockLayout) animate(canvas fyne.Canvas, clockWindow fyne.Window) {
	tick := time.NewTicker(time.Microsecond)
	go func() {
		for !c.stop {
			c.Layout(nil, canvas.Size().Subtract(fyne.NewSize(theme.Padding()*2, theme.Padding()*2)))
			canvas.Refresh(c.canvas)
			<-tick.C
		}
	}()
}

