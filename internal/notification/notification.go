package notification

import (
	"io"
	"io/fs"
	"os"

	"github.com/gen2brain/beeep"
)

func Notify(title string, message string, icon fs.File) error {
	tmp, err := os.CreateTemp("", "*.png")
	if err != nil {
		return err
	}

	name := tmp.Name()
	defer os.Remove(name)
	defer tmp.Close()

	if _, err = io.Copy(tmp, icon); err != nil {
		return err
	}

	if err := beeep.Notify(title, message, name); err != nil {
		return err
	}

	return nil
}
