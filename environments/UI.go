package environments

import (
	"fmt"
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

type PlayerUI struct {
	*ebitenui.UI
	HP        *widget2.ProgressBar
	Bag       *widget2.Container
	bagToggle *widget2.Button
}

func NewPlayerUI() *PlayerUI {
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
	bag := NewGridContainer(9)

	bagWindow := NewPopUpWindow("Bag", bag)
	eUI := &ebitenui.UI{Container: rootContainer}
	bagToggle := NewButton(eUI, bagWindow)

	rootContainer.AddChild(progressBarsContainer)
	rootContainer.AddChild(bagToggle)

	ui := PlayerUI{
		UI:        eUI,
		HP:        hProgressbar,
		Bag:       bag,
		bagToggle: bagToggle,
	}
	return &ui
}

func NewPopUpWindow(label string, windowContainer *widget2.Container) *widget2.Window {
	titleContainer := widget2.NewContainer(
		widget2.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.NRGBA{150, 150, 150, 255})),
		widget2.ContainerOpts.Layout(widget2.NewAnchorLayout()),
	)
	titleContainer.AddChild(widget2.NewText(
		widget2.TextOpts.Text(label, loadFont(20), color.NRGBA{0, 0, 255, 255}),
		widget2.TextOpts.WidgetOpts(widget2.WidgetOpts.LayoutData(widget2.AnchorLayoutData{
			HorizontalPosition: widget2.AnchorLayoutPositionCenter,
			VerticalPosition:   widget2.AnchorLayoutPositionCenter,
		})),
	))
	bagWindow := widget2.NewWindow(
		widget2.WindowOpts.Contents(windowContainer),
		widget2.WindowOpts.TitleBar(titleContainer, 50),
	)
	return bagWindow
}

func NewGridContainer(size int) *widget2.Container {
	rootContainer := widget2.NewContainer(
		widget2.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100})),
		widget2.ContainerOpts.Layout(widget2.NewGridLayout(
			widget2.GridLayoutOpts.Columns(4),
			widget2.GridLayoutOpts.Padding(widget2.NewInsetsSimple(30)),
			widget2.GridLayoutOpts.Spacing(0, 0),
		)),
	)
	for i := 0; i < size; i++ {
		innerContainer := widget2.NewContainer(
			widget2.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.NRGBA{83, 136, 162, 255})),
			widget2.ContainerOpts.WidgetOpts(
				widget2.WidgetOpts.MinSize(32, 32),
			),
		)
		rootContainer.AddChild(innerContainer)
	}
	return rootContainer
}

func NewButton(ui *ebitenui.UI, window *widget2.Window) *widget2.Button {
	button := widget2.NewButton(
		widget2.ButtonOpts.WidgetOpts(
			// instruct the ui's anchor layout to center the button both horizontally and vertically
			widget2.WidgetOpts.LayoutData(widget2.AnchorLayoutData{
				HorizontalPosition: widget2.AnchorLayoutPositionCenter,
				VerticalPosition:   widget2.AnchorLayoutPositionCenter,
			}),
		),
		widget2.ButtonOpts.Image(loadButtonImage()),
		widget2.ButtonOpts.Text("Bag", loadFont(12), &widget2.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget2.ButtonOpts.TextPadding(widget2.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		widget2.ButtonOpts.ClickedHandler(func(args *widget2.ButtonClickedEventArgs) {
			window.SetLocation(image.Rect(500, 500, 750, 750))
			ui.AddWindow(window)
		}),
	)
	button.SetLocation(image.Rect(100, 100, 150, 150))
	return button
}

func loadButtonImage() *widget2.ButtonImage {
	idle := e_m.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := e_m.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := e_m.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget2.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}
}

func loadFont(size float64) font.Face {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		fmt.Println("Load font failed.")
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}
