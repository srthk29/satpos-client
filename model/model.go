package model

type SatellitePlot struct {
	// Plot Types
	PlotTypes []PlotType `json:"plot_types"`
	// Size
	Size SizeType `json:"size_type"`
	// ImageFormat
	Format ImageFormat `json:"image_format"`
	// Colorscheme
	Colorscheme Colorscheme `json:"color_scheme"`
	// LatLong
	// Altitude
	Locations []*Location `json:"locations"`
}

type PlotType string

const (
	PlotTypePlateCarree         PlotType = "PlateCarree"
	PlotTypeNearsidePerspective PlotType = "NearsidePerspective"
)

type ImageFormat string

const (
	ImageFormatSVG ImageFormat = "svg"
	ImageFormatPNG ImageFormat = "png"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float32 `json:"altitude"`
}

type Colorscheme string

const (
	ColorschemeDefault Colorscheme = "default"
)

type SizeType string

const (
	SizeThumbnail SizeType = "thumbnail"
	SizeSmall     SizeType = "small"
	SizeMedium    SizeType = "medium"
	SizeLarge     SizeType = "large"
	SizePrint     SizeType = "print"
)
