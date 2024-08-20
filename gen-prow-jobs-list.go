package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ProwJob represents the structure of a prowjob
type ProwJob struct {
	Name        string
	Annotations map[string]string `yaml:"annotations"`
}

// ProwJobsFile represents the structure of a file containing prowjobs
type ProwJobsFile struct {
	Presubmits  map[string][]ProwJob `yaml:"presubmits"`
	Postsubmits map[string][]ProwJob `yaml:"postsubmits"`
	Periodics   []ProwJob            `yaml:"periodics"`
}

func main() {
	root := "./" // Root directory to start the search from

	jobsList := map[string][]string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only YAML files
		if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var prowJobsFile ProwJobsFile
			yaml.Unmarshal(content, &prowJobsFile)

			// Check for presubmits
			for _, jobs := range prowJobsFile.Presubmits {
				for _, job := range jobs {
					if owner, ok := job.Annotations["owner"]; ok {
						jobsList[owner] = append(jobsList[owner], job.Name)
					}
				}
			}

			// Check for postsubmits
			for _, jobs := range prowJobsFile.Postsubmits {
				for _, job := range jobs {
					if owner, ok := job.Annotations["owner"]; ok {
						jobsList[owner] = append(jobsList[owner], job.Name)
					}
				}
			}

			for _, job := range prowJobsFile.Periodics {
				if owner, ok := job.Annotations["onwer"]; ok {
					jobsList[owner] = append(jobsList[owner], job.Name)
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through the directory: %v", err)
	}
	fmt.Println("Jobs:")
	data, _ := yaml.Marshal(jobsList)
	fmt.Println(string(data))
	fmt.Println("Teams:")
	for team := range jobsList {
		fmt.Printf("%s ", team)
	}
}
