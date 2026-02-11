package model

type Plot struct {
	MediaType ImageFormat `json:"media_type"`
	PlotType  PlotType    `json:"plot_type"`
	// Content   []byte      `json:"content"`
	Content string `json:"content"`
}

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

	ShowIcon      bool      `json:"icon"`
	AddNightShade bool      `json:"nightshade"`
	Features      []Feature `json:"features"`
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

type SizeType string

const (
	SizeThumbnail SizeType = "thumbnail"
	SizeSmall     SizeType = "small"
	SizeMedium    SizeType = "medium"
	SizeLarge     SizeType = "large"
	SizePrint     SizeType = "print"
)

/*
| Name      | Description |
|-----------|-------------|
| BORDERS   | Country boundaries. |
| COASTLINE | Coastline, including major islands. |
| LAKES     | Natural and artificial lakes. |
| LAND      | Land polygons, including major islands. |
| OCEAN     | Ocean polygons. |
| RIVERS    | Single-line drainages, including lake centerlines. |
| STATES    | Internal, first-order administrative boundaries (limited to the United States at this scale). Natural Earth have first-order admin boundaries for most countries at the 1:10,000,000 scale; these may be accessed with `cartopy.feature.STATES.with_scale('10m')`. |
*/

type Feature string

const (
	FeatureBorders   Feature = "BORDERS"
	FeatureCoastLine Feature = "COASTLINE"
	FeatureLakes     Feature = "LAKES"
	FeatureLand      Feature = "LAND"
	FeatureOcean     Feature = "OCEAN"
	FeatureRivers    Feature = "RIVERS"
	FeatureStates    Feature = "STATES"
)
