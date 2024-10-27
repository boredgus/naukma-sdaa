package src

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
)

func DrawPoints(points []Point) {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 400, 400))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.BeginPath()

	var imageWidth float64 = 400

	drawBackground(gc, imageWidth)
	drawAxes(gc, imageWidth)
	drawPoints(gc, points, imageWidth)

	if err := draw2dimg.SaveToPngFile("lab4_graph.png", dest); err != nil {
		panic(err)
	}
}

func drawBackground(gc *draw2dimg.GraphicContext, width float64) {
	gc.SetStrokeColor(color.RGBA{255, 255, 255, 0})
	gc.SetLineWidth(1)
	gc.SetFillColor(color.RGBA{255, 255, 255, 0xff})

	gc.MoveTo(0, 0)
	gc.LineTo(width, 0)
	gc.LineTo(width, width)
	gc.LineTo(0, width)
	gc.Close()
	gc.FillStroke()
}

func drawAxes(gc *draw2dimg.GraphicContext, width float64) {
	gc.SetStrokeColor(color.RGBA{255, 255, 255, 0xaf})
	gc.SetLineWidth(1)
	gc.SetFillColor(color.RGBA{0, 0, 0, 0})

	gc.MoveTo(0, width/2)
	gc.LineTo(width, width/2)
	gc.MoveTo(width/2, width)
	gc.LineTo(width/2, 0)
	gc.FillStroke()
}

func drawPoints(gc *draw2dimg.GraphicContext, points []Point, imageWidth float64) {
	gc.SetFillColor(color.RGBA{0, 0, 0, 0})
	gc.SetStrokeColor(color.RGBA{255, 0, 0, 0xff})
	gc.SetLineWidth(2)

	for idx, point := range points {
		x, y := point.Transplant(150, imageWidth)

		if idx == 0 {
			gc.MoveTo(x, y)
		} else {
			gc.LineTo(x, y)
		}
		gc.FillString("str")
	}

	gc.FillStroke()
}
