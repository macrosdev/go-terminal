package termui

type Block struct {
	X             int
	Y             int
	Border        labeledBorder
	IsDisplay     bool
	HasBorder     bool
	BgColor       Attribute
	Width         int
	Height        int
	innerWidth    int
	innerHeight   int
	innerX        int
	innerY        int
	PaddingTop    int
	PaddingBottom int
	PaddingLeft   int
	PaddingRight  int
}

func NewBlock() *Block {
	d := Block{}
	d.IsDisplay = true
	d.HasBorder = true
	d.Width = 2
	d.Height = 2
	return &d
}

// compute box model
func (d *Block) align() {
	d.innerWidth = d.Width - d.PaddingLeft - d.PaddingRight
	d.innerHeight = d.Height - d.PaddingTop - d.PaddingBottom
	d.innerX = d.X + d.PaddingLeft
	d.innerY = d.Y + d.PaddingTop

	if d.HasBorder {
		d.innerHeight -= 2
		d.innerWidth -= 2
		d.Border.X = d.X
		d.Border.Y = d.Y
		d.Border.Width = d.Width
		d.Border.Height = d.Height
		d.innerX += 1
		d.innerY += 1
	}
}

func (d *Block) Buffer() []Point {
	d.align()

	ps := []Point{}
	if !d.IsDisplay {
		return ps
	}

	if d.HasBorder {
		ps = d.Border.Buffer()
	}

	for i := 0; i < d.innerWidth; i++ {
		for j := 0; j < d.innerHeight; j++ {
			p := Point{}
			p.X = d.X + 1 + i
			p.Y = d.Y + 1 + j
			p.Ch = ' '
			p.Bg = d.BgColor
			ps = append(ps, p)
		}
	}
	return ps
}
