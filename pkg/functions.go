package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type gotReposSuccessMsg []Repo
type gotReposErrMsg error

type Repo struct {
	Name string `json:"name"`
	Url  string `json:"html_url"`
}

const UserReposURL = "https://api.github.com/user/repos?visibility=all"

func GetRepos() ([]Repo, gotReposErrMsg) {
	req, err := http.NewRequest(http.MethodGet, UserReposURL, nil)
	if err != nil {
		return gotReposSuccessMsg(nil), gotReposErrMsg(err)
	}

	GH_TOKEN := os.Getenv("GH_TOKEN")

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", GH_TOKEN))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return gotReposSuccessMsg(nil), gotReposErrMsg(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return gotReposSuccessMsg(nil), gotReposErrMsg(err)
	}

	var repos []Repo

	err = json.Unmarshal(data, &repos)
	if err != nil {
		return gotReposSuccessMsg(nil), gotReposErrMsg(err)
	}
	return gotReposSuccessMsg(repos), gotReposErrMsg(nil)

}
