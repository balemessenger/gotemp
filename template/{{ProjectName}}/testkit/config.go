package testkit

import (
	"fmt"
	"{{ProjectName}}/internal"
	"os"
)

func InitTestConfig(relativePath string) *internal.Config{
	configPath := fmt.Sprintf("%s/%s", os.Getenv("PWD"), relativePath)
	fmt.Println("Config path: ", configPath)
	return internal.NewConfig(configPath)
}

