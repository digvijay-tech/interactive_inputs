package utilities

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/term"
)

// Returns width and height of current termminal window, if it fails for any reason it will return (0,0) instead.
func GetWindowSize() (width, height int) {
	w, h, err := term.GetSize(0)
	if err != nil {
		return 0, 0
	}

	return w, h
}

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

// It reads keyboard input directly into the provided byte slice, changing its content.
func RecordKeyStroke(byteArr []byte) {
	oldState := enterRawMode()
	os.Stdin.Read(byteArr)
	exitRawMode(oldState)
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

func HideDefaultTerminalCursor() { fmt.Print("\033[?25l") }
func ShowDefaultTerminalCursor() { fmt.Print("\033[?25h") }

func IsTerminalCapable() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

/************** IDENTIFYING KEYSTRONG **************/

// Returns true if byte array evaluates to enter keystroke
func IsEnter(byteArr []byte) bool {
	return byteArr[0] == 13
}

// Returns true if byte array evaluates to Ctrl+c keystroke
func IsCtrlC(byteArr []byte) bool {
	return byteArr[0] == 3
}

// Returns true if byte array evaluates to up arrow keystroke
func IsUpArrow(byteArr []byte) bool {
	return byteArr[0] == 27 && byteArr[1] == 91 && byteArr[2] == 65
}

// Returns true if byte array evaluates to down arrow keystroke
func IsDownArrow(byteArr []byte) bool {
	return byteArr[0] == 27 && byteArr[1] == 91 && byteArr[2] == 66
}

// Returns true if byte array evaluates to spacebar keystroke
func IsSpacebar(byteArr []byte) bool {
	return byteArr[0] == 32
}
