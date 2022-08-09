package table

var Styles = map[int][]string{
	/*
		style 1
		┌┬┐  ┌─┐
		├┼┤  │┼│
		└┴┘  └─┘
	*/
	0: {"┌", "┬", "┐", "├", "┼", "┤", "└", "┴", "┘", "─", "│"},

	/*
		style 2
		┏┳┓  ┏━┓
		┣╋┫  ┃╋┃
		┗┻┛  ┗━┛
	*/
	1: {"┏", "┳", "┓", "┣", "╋", "┫", "┗", "┻", "┛", "━", "┃"},
	/*
		style 2
		╔╦╗  ╔═╗
		╠╬╣  ║╬║
		╚╩╝  ╚═╝
	*/
	2: {"╔", "╦", "╗", "╠", "╬", "╣", "╚", "╩", "╝", "═", "║"},

	/*
		style 3
		+++  +-+
		+++  |+|
		+++  +-+
	*/
	3: {"+", "+", "+", "+", "+", "+", "+", "+", "+", "-", "|"},
}

var (
	LineLEdge = map[int]int{
		0: 0, // line0 left edge
		1: 3, // line1 left edge
		2: 6, // line2 left edge
	}

	LineREdge = map[int]int{
		0: 2, // line0 right edge
		1: 5, // line1 right edge
		2: 8, // line2 right edge
	}

	LineSplit = map[int]int{
		0: 1, // line0 split
		1: 4, // line1 split
		2: 7, // line2 split
	}
)
