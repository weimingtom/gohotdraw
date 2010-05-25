package gohotdraw

type Event interface {
	GetSource() interface{}
}

type EventObject struct {
	source interface{}
}

func (this *EventObject) GetSource() interface{} {
	return this.source
}

type FigureEvent struct {
	*EventObject
	rect *Rectangle
}

func NewFigureEvent(source Figure) *FigureEvent {
	event := &FigureEvent{}
	event.EventObject = &EventObject{source}
	event.rect = &Rectangle{0, 0, 0, 0}
	return event
}

func NewFigureEventRect(source Figure, rect *Rectangle) *FigureEvent {
	event := &FigureEvent{}
	event.EventObject = &EventObject{source}
	event.rect = rect
	return event
}

func (this *FigureEvent) GetFigure() Figure {
	return this.GetSource().(Figure)
}

func (this *FigureEvent) GetInvalidatedRect() *Rectangle {
	return this.rect
}

type EventHandler struct {
	view DrawingView
}

func NewEventHandler(view DrawingView) *EventHandler {
	return &EventHandler{view}
}

func (this *EventHandler) FigureAdded(event *FigureEvent) {
	this.view.Repaint()
}

func (this *EventHandler) FigureChanged(event *FigureEvent) {
	this.view.Repaint()
}

func (this *EventHandler) FigureRemoved(event *FigureEvent) {
	this.view.Repaint()
}

func (this *EventHandler) FigureRequestRemove(event *FigureEvent) {
	this.view.Repaint()
}

type InputEvent interface{}

const (
	SHIFT_KEY = 1 << 0
)

type MouseEvent struct {
	X           int
	Y           int
	Button      int
	KeyModifier int
}

func (this *MouseEvent) GetPoint() *Point {
	return &Point{this.X, this.Y}
}

func (this *MouseEvent) IsShiftDown() bool {
	return this.KeyModifier == SHIFT_KEY
}

type KeyEvent struct {
	KeyCode     int
	KeyModifier int
}

type ExposeEvent struct {
	X      int
	Y      int
	Width  int
	Height int
}

type ToolEvent struct {
	invalidatedArea *Rectangle
	*EventObject
	view DrawingView
}

//func NewToolEvent(tool Tool, view DrawingView, invalidatedArea *Rectangle) *ToolEvent {
//	event := &ToolEvent{}
//	event.invalidatedArea = invalidatedArea
//	event.EventObject = &EventObject{tool}
//	event.view = view
//	return event
//}

//func (this *ToolEvent) GetTool() Tool {
//	return this.GetSource().(Tool)
//}

//func (this *ToolEvent) GetView() DrawingView {
//	return this.view
//}

//func (this *ToolEvent) GetInvalidatedArea() *Rectangle {
//	return this.invalidatedArea
//}
