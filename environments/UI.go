package environments

import (
	"fmt"
	"github.com/ebitenui/ebitenui"
	_ "github.com/ebitenui/ebitenui"
	e_m "github.com/ebitenui/ebitenui/image"
	widget "github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
)

type PlayerUI struct {
	*ebitenui.UI
	HP        *widget.ProgressBar
	Bag       *widget.Container
	bagToggle *widget.Button
}

type QuestUI struct {
	toggle *widget.Button
	Quests []*widget.Label
}

func NewPlayerUI() *PlayerUI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.Transparent)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	progressBarsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(20),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
		),
	)

	hProgressbar := widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(200, 20),
		),
		widget.ProgressBarOpts.Images(
			&widget.ProgressBarImage{
				Idle:  e_m.NewNineSliceColor(color.NRGBA{138, 28, 82, 100}),
				Hover: e_m.NewNineSliceColor(color.NRGBA{138, 28, 82, 100}),
			},
			&widget.ProgressBarImage{
				Idle:  e_m.NewNineSliceColor(color.NRGBA{222, 0, 0, 100}),
				Hover: e_m.NewNineSliceColor(color.NRGBA{222, 0, 0, 100}),
			},
		),
		widget.ProgressBarOpts.Values(0, 10, 2),
		widget.ProgressBarOpts.TrackPadding(widget.Insets{
			Top:    2,
			Bottom: 2,
		}),
	)
	progressBarsContainer.AddChild(hProgressbar)
	bag := NewGridContainer(9)

	bagWindow := NewPopUpWindow("Bag", bag)
	eUI := &ebitenui.UI{Container: rootContainer}
	bagToggle := NewButton("Open Bag", eUI, bagWindow)

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

func NewQuestUI() *ebitenui.UI {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.Transparent)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(20)),
			widget.RowLayoutOpts.Spacing(20),
		)),
	)
	window := NewPopUpWindow("QUEST", container)
	window.SetLocation(image.Rect(1300, 200, 1500, 300))
	ui := ebitenui.UI{Container: container}
	ui.AddWindow(window)
	return &ui
}

func CreateAQuest(label, content string, container *widget.Container) {
	questLabel := widget.NewLabel(
		widget.LabelOpts.Text(label, loadFont(20), &widget.LabelColor{
			Idle:     color.White,
			Disabled: color.NRGBA{100, 100, 100, 255},
		}),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				}),
			),
		),
	)
	questContent := widget.NewLabel(
		widget.LabelOpts.Text(content, loadFont(13), &widget.LabelColor{
			Idle:     color.White,
			Disabled: color.NRGBA{100, 100, 100, 255},
		}),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				}),
			),
		),
	)
	container.AddChild(questLabel)
	container.AddChild(questContent)
}

func NewPopUpWindow(label string, contentContainer *widget.Container) *widget.Window {
	titleContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.NRGBA{150, 150, 150, 255})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleContainer.AddChild(widget.NewText(
		widget.TextOpts.Text(label, loadFont(20), color.NRGBA{0, 0, 255, 255}),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	))
	bagWindow := widget.NewWindow(
		widget.WindowOpts.Contents(contentContainer),
		widget.WindowOpts.TitleBar(titleContainer, 50),
	)
	return bagWindow
}

func NewGridContainer(size int) *widget.Container {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(4),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(30)),
			widget.GridLayoutOpts.Spacing(0, 0),
		)),
	)
	for i := 0; i < size; i++ {
		innerContainer := widget.NewContainer(
			widget.ContainerOpts.BackgroundImage(e_m.NewNineSliceColor(color.NRGBA{83, 136, 162, 255})),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(32, 32),
			),
		)
		rootContainer.AddChild(innerContainer)
	}
	return rootContainer
}

func NewButton(label string, ui *ebitenui.UI, window *widget.Window) *widget.Button {
	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			// instruct the ui's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ButtonOpts.Image(loadButtonImage()),
		widget.ButtonOpts.Text(label, loadFont(12), &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			window.SetLocation(image.Rect(500, 500, 750, 750))
			ui.AddWindow(window)
		}),
	)
	button.SetLocation(image.Rect(100, 100, 150, 150))
	return button
}

func loadButtonImage() *widget.ButtonImage {
	idle := e_m.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := e_m.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := e_m.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
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
