package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const InboxId = "943367551"
const ApiUrl = "https://api.todoist.com/rest/v1"

func main() {
	apiToken := os.Getenv("TODOIST_API_TOKEN")
	if apiToken == "" {
		log.Fatal("API Token required")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Task name: ")

	input, e := reader.ReadString('\n')
	if e != nil {
		log.Fatal(e)
	}

	if input == "\n" {
		log.Fatal("Task name can't be blank")
	}

	taskName := input[0 : len(input)-1]

	body := strings.NewReader(fmt.Sprintf("{\"content\": \"%s\"}", taskName))
	req, e := http.NewRequest("POST", fmt.Sprintf("%s/tasks", ApiUrl), body)
	if e != nil {
		log.Fatal(e)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	resBody, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(string(resBody))
}
