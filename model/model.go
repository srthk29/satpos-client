package model

type SatellitePlot struct {
	NoradId int64 `json:"norad"`
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
	Positions []*GeodeticPosition `json:"locations"`
	// At Now UTC
	NowPosition *GeodeticPosition `json:"now_location"`
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

type GeodeticPosition struct {
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longitude"`
	Altitude   float64 `json:"altitude"`
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
