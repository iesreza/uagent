package classifiers

import "regexp"

var regexDataParts = regexp.MustCompile(`\(.+?\)`)
var regexModel = regexp.MustCompile(`([a-zA-Z0-9\s\.-]+)`)

var Classifiers = map[string][]ClassifierModule{
	"OS":      []ClassifierModule{},
	"Model":   []ClassifierModule{},
	"Device":  []ClassifierModule{Console{}, iPad{}, iPhone{}, iPod{}, WinMobile{}, blackBerry{}, WinPC{}, Android{}, Mac{}, LinuxOS{}},
	"Browser": []ClassifierModule{WebBrowser{}},
}

type Describer struct {
	OS       string
	Model    string
	Device   string
	Platform string
	Browser  string
}

type Devices uint8

const (
	WindowsPhone  = "Windows Phone"
	Windows       = "Windows"
	MacOS         = "MacOS"
	Linux         = "Linux"
	SmartTV       = "Smart TV"
	IoS           = "iOS"
	Tablet        = "Tablet"
	Mobile        = "Mobile"
	WindowsMobile = "Windows Mobile"
	MusicPlayer   = "Music Player"
	Computer      = "Computer"
	Laptop        = "Laptop"
	Desktop       = "Desktop"
	GameConsole   = "Game Console"
	WebOS         = "WebOS"
	EBookReader   = "EBook Reader"
)

type ClassifierModule interface {
	Eval(ua string, desc *Describer) bool
}
