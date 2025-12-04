package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Given a github user login, return name and number of public repos
func main() {
	fmt.Println(UserInfo("p-andhika"))
}

func demo() {
	resp, err := http.Get("https://api.github.com/users/p-andhika")
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s\n", resp.Status)
		return
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("Content-Type:", ctype)

	// io.Copy(os.Stdout, resp.Body)

	// Anonymous struct
	var reply struct {
		Name        string
		PublicRepos int `json:"public_repos"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&reply); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(reply.Name, reply.PublicRepos)
}

// UserInfo return name and number of pulbic repose from Gitub API
func UserInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login

	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Status)
	}

	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	// Anonymous struct
	var reply struct {
		Name        string
		PublicRepos int `json:"public_repos"`
	}

	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&reply); err != nil {
		return "", 0, err
	}

	return reply.Name, reply.PublicRepos, nil
}

/*
* JSON <-> GO
*
*	Types
*	string <-> string
*	true/false <-> bool
*	number <-> float64, float32, int, int8 ... int 64, uint, uint8, ...
*	array <-> []T, []any
*	object <-> map[string]any, struct
*
*	encoding/json API
*	JSON -> []byte -> Go: Unmarshal
*	Go -> []byte -> JSON: Marshal
*	JSON -> io.Reader -> Go: Decoder
*	Go -> io.Writer -> JSON: Encoder
 */
