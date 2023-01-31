package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type Student struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func jsonStructDemo() {
	// json struct marshal and unmarshal
	s := Student{Name: "hxia", Age: 18}
	data, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	var student Student
	if err := json.Unmarshal(data, &student); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("unmarshal json: Name: %s, Age: %d\n", student.Name, student.Age)
}

func jsonStructSliceDemo() {
	// json slice struct marshal and unmarshal
	ss := []Student{
		{Name: "hxia", Age: 18},
		{Name: "xiaoming", Age: 20},
	}
	datas, err := json.MarshalIndent(ss, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(datas))

	var students []Student
	if err := json.Unmarshal(datas, &students); err != nil {
		fmt.Println(err)
	}

	for i, student := range students {
		fmt.Printf("student %d: Name %s Age %d\n", i, student.Name, student.Age)
	}
}

func SearchGithubIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	//jsonStructDemo()
	//jsonStructSliceDemo()

	result, err := SearchGithubIssues(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %-9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
