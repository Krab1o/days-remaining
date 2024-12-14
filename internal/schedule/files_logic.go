package schedule

import (
	"days-remaining/internal/data"
	"log"
	"os"
)

var directoryPath = "data"
var dataPath = directoryPath + "/data.json"
var directoryPermissions = 0777
var dataPermissions = 0644

func createDirectory() {
	folderErr := os.Mkdir(directoryPath, os.FileMode(directoryPermissions))
	if (folderErr != nil) {
		log.Print(folderErr.Error())
	}
}

func remove(s []data.Sending, i int) []data.Sending {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func writeFile(data []byte) {
	dataPermissions := 0644
	err := os.WriteFile(dataPath, data, os.FileMode(dataPermissions))
	if err != nil {
		panic(err)
	}
}