package inputs

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/term"
)

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func enterRawMode() (oldState *term.State) {
	currentState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return currentState
}

func exitRawMode(oldState *term.State) {
	term.Restore(int(os.Stdin.Fd()), oldState)
}
