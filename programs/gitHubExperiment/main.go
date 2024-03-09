// The below code snippet fetches GitHub golang repositories data based on specific time intervals, processes the data, and writes it to a CSV file.

package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
)

type githubData struct {
	Name        string
	CloneURL    string
	Description string
}

func main() {
	client := github.NewClient(nil)
	startDate := time.Date(2024, 1, 12, 0, 0, 0, 0, time.UTC)
	for startDate.Unix() < time.Now().Unix() {
		endDate := startDate.Add(24 * 2 * time.Hour)
		query := generateQuery(startDate, endDate)
		data := fetchData(client, query)
		writeToCSV(data)
		startDate = startDate.Add(24 * 2 * time.Hour)
		// To stop the rate limiting on github
		time.Sleep(10 * time.Second)
	}
}

func generateQuery(startDate time.Time, endDate time.Time) string {
	return fmt.Sprintf("language:Go created:\"%d-%02d-%02d .. %d-%02d-%02d\"",
		startDate.Year(), startDate.Month(), startDate.Day(),
		endDate.Year(), endDate.Month(), endDate.Day())
}

func fetchData(client *github.Client, query string) []githubData {
	var data []githubData

	res, _, err := client.Search.Repositories(context.Background(), query, &github.SearchOptions{
		Order: "desc", ListOptions: github.ListOptions{PerPage: 100},
	})
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}

	for _, rep := range res.Repositories {
		desc := ""
		if rep.Description != nil {
			desc = *rep.Description
		}

		y, t, d := rep.CreatedAt.Date()
		fmt.Println(y, t, d)

		data = append(data, githubData{
			Name:        *rep.Name,
			CloneURL:    *rep.CloneURL,
			Description: desc,
		})

	}

	return data
}

func writeToCSV(data []githubData) {
	filePath := "data.csv"
	fileMode := os.O_APPEND | os.O_CREATE | os.O_WRONLY

	file, err := os.OpenFile(filePath, fileMode, os.ModeAppend)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Check if the file is empty to determine whether to write the header
	fileInfo, _ := file.Stat()
	if fileInfo.Size() == 0 {
		header := []string{"Name", "URL", "Description"}
		if err := writer.Write(header); err != nil {
			log.Fatal("Cannot write header to file", err)
		}
	}

	for _, value := range data {
		record := []string{value.Name, value.CloneURL, value.Description}
		if err := writer.Write(record); err != nil {
			log.Fatal("Cannot write to file", err)
		}

	}

	log.Println("Data has been written and appended to data.csv")
}
