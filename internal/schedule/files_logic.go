package schedule

import (
	// "days-remaining/internal/data"
	"log"
	"os"
)

const directoryPath = "data"
const dataPath = directoryPath + "/data.json"
const directoryPermissions = 0777
const dataPermissions = 0644

func createDirectory() {
	folderErr := os.Mkdir(directoryPath, os.FileMode(directoryPermissions))
	if (folderErr != nil) {
		log.Print(folderErr.Error())
	}
}

func writeFile(data []byte) {
	dataPermissions := 0644
	err := os.WriteFile(dataPath, data, os.FileMode(dataPermissions))
	if err != nil {
		panic(err)
	}
}