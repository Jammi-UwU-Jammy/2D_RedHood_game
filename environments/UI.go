package environments

import (
	"github.com/ebitenui/ebitenui"
	_ "github.com/ebitenui/ebitenui"
	e_m "github.com/ebitenui/ebitenui/image"
	widget2 "github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
)

type UIWindow struct {
	ebitenui.UI
}

func NewUI() *ebitenui.UI {
	ui := ebitenui.UI{}
	buttonImg, _ := loadButtonImage()
	face, _ := loadFont(20)
	titleFace, _ := loadFont(12)

	windowContainer := widget2.NewContainer(
		widget2.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255})),
		widget2.ContainerOpts.Layout(widget2.NewAnchorLayout()),
	)
	windowContainer.AddChild(widget2.NewText(
		widget2.TextOpts.Text("Hello.", face, color.NRGBA{R: 254, G: 254, B: 254, A: 254}),
		widget2.TextOpts.WidgetOpts(widget2.WidgetOpts.LayoutData(widget2.AnchorLayoutData{
			HorizontalPosition: widget2.AnchorLayoutPositionCenter,
			VerticalPosition:   widget2.AnchorLayoutPositionCenter,
		})),
	))
	titleContainer := widget2.NewContainer(
		widget2.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.NRGBA{R: 150, G: 150, B: 150, A: 255})),
		widget2.ContainerOpts.Layout(widget2.NewAnchorLayout()),
	)

	titleContainer.AddChild(widget2.NewText(
		widget2.TextOpts.Text("Window Title", titleFace, color.NRGBA{R: 254, G: 255, B: 255, A: 255}),
		widget2.TextOpts.WidgetOpts(widget2.WidgetOpts.LayoutData(widget2.AnchorLayoutData{
			HorizontalPosition: widget2.AnchorLayoutPositionCenter,
			VerticalPosition:   widget2.AnchorLayoutPositionCenter,
		})),
	))

	window := widget2.NewWindow(
		widget2.WindowOpts.Contents(windowContainer),
		widget2.WindowOpts.TitleBar(titleContainer, 25),
		widget2.WindowOpts.Modal(),
		widget2.WindowOpts.CloseMode(widget2.CLICK_OUT),
		widget2.WindowOpts.MaxSize(300, 200),
	)

	rootContainer := widget2.NewContainer(
		widget2.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.Transparent)),
		widget2.ContainerOpts.Layout(widget2.NewAnchorLayout()),
	)

	button := widget2.NewButton(
		widget2.ButtonOpts.WidgetOpts(
			widget2.WidgetOpts.LayoutData(widget2.AnchorLayoutData{
				HorizontalPosition: widget2.AnchorLayoutPositionCenter,
				VerticalPosition:   widget2.AnchorLayoutPositionCenter,
			}),
		),
		widget2.ButtonOpts.Image(buttonImg),
		widget2.ButtonOpts.Text("Open Bag", face, &widget2.ButtonTextColor{
			Idle: color.NRGBA{R: 0xdf, G: 0xf4, B: 0xff, A: 0xff},
		}),
		widget2.ButtonOpts.TextPadding(widget2.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		widget2.ButtonOpts.ClickedHandler(func(args *widget2.ButtonClickedEventArgs) {
			//Get the preferred size of the content
			x, y := window.Contents.PreferredSize()
			r := image.Rect(0, 0, x, y)
			r = r.Add(image.Point{100, 50})
			window.SetLocation(r)
			ui.AddWindow(window)
		}),
	)
	rootContainer.AddChild(button)
	ui.Container = rootContainer

	return &ui
}

func PlayerUI() *ebitenui.UI {
	rootContainer := widget2.NewContainer(
		widget2.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.Transparent)),
		widget2.ContainerOpts.Layout(widget2.NewAnchorLayout()),
	)

	progressBarsContainer := widget2.NewContainer(
		widget2.ContainerOpts.Layout(widget2.NewRowLayout(
			widget2.RowLayoutOpts.Direction(widget2.DirectionVertical),
			widget2.RowLayoutOpts.Spacing(20),
		)),
		widget2.ContainerOpts.WidgetOpts(
			widget2.WidgetOpts.LayoutData(widget2.AnchorLayoutData{
				HorizontalPosition: widget2.AnchorLayoutPositionStart,
				VerticalPosition:   widget2.AnchorLayoutPositionStart,
			}),
		),
	)

	hProgressbar := widget2.NewProgressBar(
		widget2.ProgressBarOpts.WidgetOpts(
			widget2.WidgetOpts.MinSize(200, 20),
		),
		widget2.ProgressBarOpts.Images(
			&widget2.ProgressBarImage{
				Idle:  e_m.NewNineSliceColor(color.NRGBA{138, 28, 82, 100}),
				Hover: e_m.NewNineSliceColor(color.NRGBA{138, 28, 82, 100}),
			},
			&widget2.ProgressBarImage{
				Idle:  e_m.NewNineSliceColor(color.NRGBA{222, 0, 0, 100}),
				Hover: e_m.NewNineSliceColor(color.NRGBA{222, 0, 0, 100}),
			},
		),
		widget2.ProgressBarOpts.Values(0, 10, 2),
		widget2.ProgressBarOpts.TrackPadding(widget2.Insets{
			Top:    2,
			Bottom: 2,
		}),
	)
	progressBarsContainer.AddChild(hProgressbar)
	rootContainer.AddChild(progressBarsContainer)
	ui := ebitenui.UI{
		Container: rootContainer,
	}
	return &ui
}

func loadButtonImage() (*widget2.ButtonImage, error) {
	idle := e_m.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := e_m.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := e_m.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget2.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
