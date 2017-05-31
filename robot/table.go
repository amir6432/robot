package robot

type Table interface {
	InBounds(x, y int) bool
}

func NewTable(width, height int) Table {
	return &table{width, height}
}

type table struct {
	width  int
	height int
}

func (t *table) InBounds(x, y int) bool {
	return x >= 0 && x < t.width && y >= 0 && y < t.height
}
