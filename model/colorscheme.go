package model

type Colorscheme struct {
	Theme  `json:"theme"`
	Accent `json:"accent"`
}

type Theme string

const (
	ThemeLight          Theme = "light"
	ThemeDark           Theme = "dark"
	ThemeDim            Theme = "dim"
	ThemeSolarizedLight Theme = "solarized-light"
	ThemeSolarizedDark  Theme = "solarized-dark"
	ThemeMuted          Theme = "muted"
	ThemeOcean          Theme = "ocean"
	ThemeNebula         Theme = "nebula"
	ThemeAstrophage     Theme = "astrophage"
	ThemeMidnight       Theme = "midnight"
	ThemeGraphite       Theme = "graphite"
	ThemeForest         Theme = "forest"
	ThemePlum           Theme = "plum"
)

func Themes() []Theme {
	return []Theme{
		ThemeLight,
		ThemeDark,
		ThemeDim,
		ThemeSolarizedLight,
		ThemeSolarizedDark,
		ThemeMuted,
		ThemeOcean,
		ThemeNebula,
		ThemeAstrophage,
		ThemeMidnight,
		ThemeGraphite,
		ThemeForest,
		ThemePlum,
	}
}

type Accent string

const (
	AccentDefault  Accent = "default"
	AccentWhite    Accent = "white"
	AccentGreen    Accent = "green"
	AccentPink     Accent = "pink"
	AccentRed      Accent = "red"
	AccentDarkRed  Accent = "darkred"
	AccentPurple   Accent = "purple"
	AccentBlue     Accent = "blue"
	AccentDarkBlue Accent = "darkblue"
	AccentDarkPink Accent = "darkpink"
	AccentEarth    Accent = "earth"
	AccentCyan     Accent = "cyan"
	AccentTeal     Accent = "teal"
	AccentSky      Accent = "sky"
	AccentViolet   Accent = "violet"
	AccentAmber    Accent = "amber"
	AccentYellow   Accent = "yellow"
	AccentLime     Accent = "lime"
	AccentOrange   Accent = "orange"
)

func Accents() []Accent {
	return []Accent{
		AccentDefault,
		AccentWhite,
		AccentGreen,
		AccentPink,
		AccentRed,
		AccentDarkRed,
		AccentPurple,
		AccentBlue,
		AccentDarkBlue,
		AccentDarkPink,
		AccentEarth,
		AccentCyan,
		AccentTeal,
		AccentSky,
		AccentViolet,
		AccentAmber,
		AccentYellow,
		AccentLime,
		AccentOrange,
	}
}
