package config

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"os"
)

func InitializeConfiguration(environment string) (*Configuration, error) {
	var config Configuration

	currentDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	filePath := fmt.Sprintf("%s%cconfig%c%s.hcl",
		currentDir,
		os.PathSeparator,
		os.PathSeparator,
		environment)

	if err = hclsimple.DecodeFile(filePath, nil, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
