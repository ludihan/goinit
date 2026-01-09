package builder

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rafaeldepontes/goinit/internal/log"
	"github.com/rafaeldepontes/goinit/internal/project/builder/templates"
)

const (
	// Files
	GoModFile = "go.mod"
)

func createGoMod(name string, log *log.Logger) error {
	template := templates.GodModDefault
	scanner := bufio.NewScanner(os.Stdin)

	var gitUsername string
	log.Info("Git username: ")
	if scanner.Scan() {
		gitUsername = scanner.Text()
	}

	if gitUsername != "" {
		template = fmt.Sprintf(templates.GoMod, gitUsername, name)
	}

	if err := os.WriteFile(
		GoModFile,
		[]byte(template),
		OwnerPropertyMode,
	); err != nil {
		return err
	}

	return os.Rename(GoModFile, fmt.Sprintf("%s/%s", name, GoModFile))
}
