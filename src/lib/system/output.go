package system

import (
	"github.com/fatih/color"
	"os"
)

func Exit(err error) {
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
}
