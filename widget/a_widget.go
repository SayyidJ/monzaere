package widget

// hold a state of widget
type Widget interface {
	CreateRender() RenderLayout
	DisposeRender()
}
