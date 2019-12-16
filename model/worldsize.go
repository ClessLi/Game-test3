package model

type WorldSize struct {
	WorldWidth   int32
	WorldHeight  int32
	ScreenWidth  int32
	ScreenHeight int32
	Cell         int32
	DrawHelpText bool
}

func NewWorldSize(width, height, screenW, screenH, cell int32, hasHelp bool) *WorldSize {
	return &WorldSize{
		WorldWidth:   width,
		WorldHeight:  height,
		ScreenWidth:  screenW,
		ScreenHeight: screenH,
		Cell:         cell,
		DrawHelpText: hasHelp,
	}
}
