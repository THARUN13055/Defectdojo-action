package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/truemilk/go-defectdojo/defectdojo"
)

type datas struct {
	filePath          string
	EngagementName    string
	ScanType          string
	version           string
	github_buildId    string
	github_branchTag  string
	github_commitHash string
	ProductName       string
}

func files(filePath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("The error is %v\n", err)
	}

	// Join the current working directory with the file path
	fullPath := filepath.Join(cwd, filePath)

	return fullPath
}

func importscan(data datas) defectdojo.ImportScan {
	return defectdojo.ImportScan{
		MinimumSeverity:   ptrString("Info"),
		Active:            ptrBool(true),
		Verified:          ptrBool(true),
		File:              &data.filePath,
		EngagementName:    &data.EngagementName,
		ScanType:          &data.ScanType,
		Environment:       ptrString("Development"),
		AutoCreateContext: ptrBool(true),
		ProductName:       &data.ProductName,
		ProductTypeName:   ptrString("Research and Development"),
		Version:           &data.version,
		BuildId:           &data.github_buildId,
		CommitHash:        &data.github_commitHash,
		BranchTag:         &data.github_branchTag,
	}
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
}

func main() {

	//File Path of Your report
	//If you add Another Report You need to Add file path and also DataList
	filename := os.Getenv("FILE_PATH_WITH_FILE_NAME")
	filepath := files(filename)
	product_name := os.Getenv("PRODUCT_NAME")

	// Scan Type of the file
	scantype := os.Getenv("SCAN_TYPE")

	// Retrieve the github CI/CD environment variables
	reponame := os.Getenv("GITHUB_REPOSITORY")
	version := os.Getenv("GITHUB_RUN_NUMBER")
	github_buildId := os.Getenv("GITHUB_RUN_ID")
	github_commitHash := os.Getenv("GITHUB_SHA")
	github_branchTag := os.Getenv("GITHUB_REF")

	// User Auth
	url := os.Getenv("DEFECTDOJO_URL")
	authorizedToken := os.Getenv("DEFECTDOJO_TOKEN")
	EngagementNamed := reponame + "_" + version

	// Define multiple data structs
	dataList := []datas{
		{
			filePath:          filepath,
			EngagementName:    EngagementNamed,
			ScanType:          scantype,
			version:           version,
			github_buildId:    github_buildId,
			github_commitHash: github_commitHash,
			github_branchTag:  github_branchTag,
			ProductName:       product_name,
		},
	}

	// Creating the Client
	dj, err := defectdojo.NewDojoClient(url, authorizedToken, nil)
	if err != nil {
		log.Fatalf("Failed to create DefectDojo client: %v", err)
	}

	// Condition for failing the unwanted report upload.

	fileuploadcondition := false

	// Looping the Data
	for _, data := range dataList {
		// Check if the file exists
		if _, err := os.Stat(data.filePath); os.IsNotExist(err) {
			log.Fatalf("File does not exist: %v", err)
			continue
		}

		fileuploadcondition = true
		// Import the scan
		scan := importscan(data)

		// Use the Create method from the ImportScan service
		_, err := dj.ImportScan.Create(context.Background(), &scan)
		if err != nil {
			// Print error details
			fmt.Printf("Failed to import scan results for %s: %v\n", data.filePath, err)
		}
	}

	// Give as Output if No file is founded.
	if !fileuploadcondition {
		fmt.Println("No file Path has not been Founded")
	}

}

func ptrString(s string) *string {
	return &s
}

func ptrBool(b bool) *bool {
	return &b
}
