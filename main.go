package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {

	jsonPath := "resume.json"

	resumeData := GetJsonData(jsonPath)
	log.Println(resumeData)
}

func GetJsonData(jsonPath string) JsonResume {

	var resumeData JsonResume

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Panic(err.Error())
	}

	err = json.Unmarshal(data, &resumeData)
	if err != nil {
		log.Panic(err.Error())
	}

	return resumeData
}
