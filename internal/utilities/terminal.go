package utilities

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/term"
)

func ClearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func EnterRawMode() (oldState *term.State) {
	currentState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return currentState
}

func ExitRawMode(oldState *term.State) {
	term.Restore(int(os.Stdin.Fd()), oldState)
}

func HideDefaultTerminalCursor() { fmt.Print("\033[?25l") }
func ShowDefaultTerminalCursor() { fmt.Print("\033[?25h") }
