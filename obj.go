package pdftemplating

const (
	AlignmentLeft   = 0
	AlignmentCenter = 1
	AlignmentRight  = 2

	OrientationLandscape = 0
	OrientationPortrait  = 1
)

type Replacement struct {
	Default    string    `yaml:"default" json:"default_value"`
	FontFamily string    `yaml:"font-family" json:"font_family"`
	FontJson   string    `yaml:"font-json" json:"font_json"`
	FontSize   float64   `yaml:"font-size" json:"font_size"`
	PosX       float64   `yaml:"pos-x" json:"x"`
	PosY       float64   `yaml:"pos-y" json:"y"`
	Alignment  Alignment `yaml:"align" json:"align"`
}

type Orientation int

func (o Orientation) String() string {
	switch o {
	case OrientationPortrait:
		return "P"
	case OrientationLandscape:
		return "L"
	}
	return "L"
}

type Alignment int

func (a Alignment) String() string {
	switch a {
	case AlignmentLeft:
		return "L"
	case AlignmentCenter:
		return "C"
	case AlignmentRight:
		return "R"
	}
	return "L"
}
