package helper

import (
	"log"
	"path"

	"github.com/kardianos/osext"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getProjectDirectory() string {
	currentWd, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	parent := path.Dir(currentWd)
	projectDirectory := path.Dir(parent)

	return projectDirectory
}

func GetTemplatesDirectory() string {
	projectDirectory := getProjectDirectory()
	return path.Join(projectDirectory, "templates")
}
