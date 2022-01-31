package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

type Targets struct {
	Targets []Server `json:"targets"`
}

type Server struct {
	Host     string   `json:"host"`
	Friendly string   `json:"friendly"`
	Options  []string `json:"options"`
}

func main() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Print(err)
	}

	var target Targets

	err = json.Unmarshal(data, &target)
	if err != nil {
		fmt.Println("error:", err)
	}

	var servers []string

	for i := 0; i < len(target.Targets); i++ {
		servers = append(servers, target.Targets[i].Host+" | "+target.Targets[i].Friendly)
	}

	prompt := promptui.Select{
		Label: "Select server",
		Items: servers,
		Size:  len(target.Targets),
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	serverIP := strings.TrimSpace(strings.Split(result, "|")[0])

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
