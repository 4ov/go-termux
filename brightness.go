package termux

import (
	"fmt"
	"os/exec"
)

func Brightness(brightness int) error {
	if brightness < 0 || brightness > 255 {
		return fmt.Errorf("Brightness must be between 0 and 100")
	}

	return exec.Command("termux-brightness", fmt.Sprintf("%d", brightness)).Run()
}
