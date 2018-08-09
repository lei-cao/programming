package visualizer

type ColorScheme struct {
	BackgroundColor string
	BarColor        string
	AColor          string
	BColor          string
}

var defaultColor = ColorScheme{
	"#012A36",
	"#A8A7A0",
	"#2AB7B7",
	"#0E7C7B",
}
