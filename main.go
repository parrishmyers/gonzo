package main

import (
	"fmt"
	"strings"
	"os/exec"

	"gonzo/internal/i3parse"
)

func main() {
	mode, _, err := i3parse.ParseFromRunning("sway", true);
	// config, err := i3parse.GetConfigFromRunningSway()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(mode)
	//fmt.Println(variables)

	var keybindHelp strings.Builder

	for _, value := range mode {
		for _, bind := range value.Bindings {
			var keybind strings.Builder
			for _, mod := range bind.Modifiers {
				keybind.WriteString(mod)
				keybind.WriteString("+")
			}
			keybind.WriteString(bind.Key)
			fmt.Fprintf(&keybindHelp, "%-30s │ %.50s\n", keybind.String(), bind.Command)
		}
	}

	cmd := exec.Command("fuzzel", "--dmenu", "--prompt=\"⌨  Active Shortcuts: \"",
    "--index", "--width=80", "--lines=25", "--layer=overlay", "--tabs=4")
	cmd.Stdin = strings.NewReader(keybindHelp.String())
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(keybindHelp.String())
}
