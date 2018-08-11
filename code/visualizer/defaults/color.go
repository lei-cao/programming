package defaults

type ColorScheme struct {
	BackgroundColor string
	BarColor        string
	AColor          string
	BColor          string
	CColor          string
}

var DefaultColor = ColorScheme{
	"#012A36",
	"#A8A7A0",
	"#2AB7B7",
	"#0E7C7B",
	"#F25243",
}
