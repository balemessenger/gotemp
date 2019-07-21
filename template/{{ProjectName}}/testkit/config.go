package testkit

import (
	"fmt"
	"{{ProjectName}}/internal"
	"os"
)

func InitTestConfig(relativePath string) {
	configPath := fmt.Sprintf("%s/%s", os.Getenv("PWD"), relativePath)
	fmt.Println("Config path: ", configPath)
	internal.GetConfig().Initialize(configPath)
}
