package pretty

import (
	"os"

	"github.com/mattn/go-isatty"
	"github.com/robocorp/rcc/common"
)

var (
	Iconic      bool
	Disabled    bool
	Interactive bool
	White       string
	Grey        string
	Black       string
	Red         string
	Green       string
	Yellow      string
	Cyan        string
	Reset       string
	Sparkles    string
	Rocket      string
)

func Setup() {
	stdin := isatty.IsTerminal(os.Stdin.Fd())
	stdout := isatty.IsTerminal(os.Stdout.Fd())
	stderr := isatty.IsTerminal(os.Stderr.Fd())
	Interactive = stdin && stdout && stderr

	localSetup()

	common.Trace("Interactive mode enabled: %v; colors enabled: %v; icons enabled: %v", Interactive, !Disabled, Iconic)
	if Interactive && !Disabled {
		White = csi("97m")
		Grey = csi("90m")
		Black = csi("30m")
		Red = csi("91m")
		Green = csi("92m")
		Cyan = csi("96m")
		Yellow = csi("93m")
		Reset = csi("0m")
	}
	if Iconic {
		Sparkles = "\u2728 "
		Rocket = "\U0001F680 "
	}
}
