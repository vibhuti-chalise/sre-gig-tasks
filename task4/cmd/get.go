/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets joke from an external source",
	Long:  `get a joke from some Jokes API`,
	Run: func(cmd *cobra.Command, args []string) {
		getJoke()
		//	joke := strings.Join(args, " ")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

type Jokes struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:status`
}

// Gets one joke at a time from https://icanhazdadjoke.com/
func getJokefromAPI(baseAPI string) []byte {
	req, err := http.NewRequest(http.MethodGet, baseAPI, nil)

	if err != nil {
		log.Printf("Could not request a joke - %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "jokes-app CLI (github.com/vibhuti-chalise/jokes-app)")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Could not make a request to get a joke - %v", err)
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read response body- %v", err)
	}
	return resBytes
}

// copies the joke to a struct variable
func getJoke() {
	url := "https://icanhazdadjoke.com/"
	resBytes := getJokefromAPI(url)
	j1 := Jokes{}
	if err := json.Unmarshal(resBytes, &j1); err != nil {
		log.Printf("Could not unmarshal respone - %v", err)
	}

	fmt.Printf("%v \n", j1.Joke)

	saveJoketoFile(j1.Joke)
}

// saves joke to a file in the current directory
func saveJoketoFile(j string) {
	file, err := os.OpenFile("jokes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)
	if err != nil {
		log.Printf("unable to open the file")
	}

	defer file.Close()
	if _, err := file.WriteString(j + "\n"); err != nil {
		log.Println("Unable to add joke to the file")
	}
}
