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
)

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
)
