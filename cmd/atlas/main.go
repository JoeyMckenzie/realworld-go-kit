package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

const schemaDefinitionFile = "atlas.hcl"
const atlasCommand = "atlas schema"

var formatSchemaCommand = fmt.Sprintf("%s fmt .", atlasCommand)

// mergeFileIntoSchemaDefinition inspects the definition file partial and appends
// to the atlas definition file defined in the root of the project.
func mergeFileIntoSchemaDefinition(filePath string, schemaDefinitionFile *os.File) {
	log.Printf("Merging file: %s", filePath)
	definitionFilePartial, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("Error opening definition file partial %s - %v", filePath, err)
	}

	defer func(definitionFilePartial *os.File) {
		if err := definitionFilePartial.Close(); err != nil {
			log.Fatal(err)
		}
	}(definitionFilePartial)

	if _, err = io.Copy(schemaDefinitionFile, definitionFilePartial); err != nil {
		log.Fatalf("Could not copy contents of file %s - %v", filePath, err)
	}

	if _, err = schemaDefinitionFile.Write([]byte("\n")); err != nil {
		log.Fatalf("Error writing new line character: %v", err)
	}

	log.Printf("%s merged successfully!", filePath)
}

// main batches the task to merge all atlas HCL DDL files together.
func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error reading current directory - %v", err)
	}

	atlasDirectory := fmt.Sprintf("%s%cschema", cwd, os.PathSeparator)
	atlasFiles, err := os.ReadDir(atlasDirectory)

	if err != nil {
		log.Fatalf("Error reading atlas directory - %v", err)
	}

	var filesToMerge []string

	for _, file := range atlasFiles {
		filePath := fmt.Sprintf("%s%c%s", atlasDirectory, os.PathSeparator, file.Name())
		filesToMerge = append(filesToMerge, filePath)
	}

	log.Printf("Found %d schema files to merge, removing existing definition file...", len(atlasFiles))
	schemaDefinitionFilePath := fmt.Sprintf("%s%c%s", cwd, os.PathSeparator, schemaDefinitionFile)

	if err = os.Remove(schemaDefinitionFilePath); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Could not remove the definition file - %v", err)
	}

	log.Printf("Schema file removed, recreating defintion file at root %s...", schemaDefinitionFilePath)
	rehydratedFile, err := os.Create(schemaDefinitionFilePath)

	defer func(rehydratedFile *os.File) {
		if err := rehydratedFile.Close(); err != nil {
			log.Fatal(err)
		}
	}(rehydratedFile)

	if err != nil {
		log.Fatalf("Could not recreate the definition file - %v", err)
	}

	log.Printf("Merging definition partials into %s...", schemaDefinitionFile)
	for _, filePath := range filesToMerge {
		mergeFileIntoSchemaDefinition(filePath, rehydratedFile)
	}

	log.Print("Success! Formatting schema file...")
	exec.Command(formatSchemaCommand)

	log.Print("Atlas definition file created!")
}
