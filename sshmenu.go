package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"io/ioutil"
	"log"
	"moul.io/banner"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type Targets struct {
	Targets []Server `json:"targets"`
}

type Server struct {
	Host     string   `json:"host"`
	Friendly string   `json:"friendly"`
	Options  []string `json:"options"`
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	checkError(err)
}

func main() {
	clear()
	fmt.Println(banner.Inline("ssh menu"))
	appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkError(err)
	configPath := path.Join(appDir, "config.json")
	data, err := ioutil.ReadFile(configPath)
	checkError(err)

	var target Targets

	err = json.Unmarshal(data, &target)
	checkError(err)

	var servers []string

	for i := 0; i < len(target.Targets); i++ {
		servers = append(servers, target.Targets[i].Host+" | "+target.Targets[i].Friendly)
	}

	var qs = []*survey.Question{
		{
			Name: "server",
			Prompt: &survey.Select{
				Message: "Choose server:",
				Options: servers,
				Default: servers[0],
			},
		},
	}

	answers := struct {
		Server string `survey:"server"`
	}{}

	err = survey.Ask(qs, &answers, survey.WithPageSize(len(servers)))
	checkError(err)

	serverIP := strings.TrimSpace(strings.Split(answers.Server, "|")[0])

	hostOptions := getSettings(target, serverIP)
	connectSettings := append(hostOptions, serverIP)

	connectSSH(connectSettings...)

}

func getSettings(target Targets, serverIP string) []string {
	for _, host := range target.Targets {
		if host.Host == serverIP {
			return host.Options
		}
	}
	return []string{}
}

func connectSSH(param ...string) {
	cmd := exec.Command("ssh", param...)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
