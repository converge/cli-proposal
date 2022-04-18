package main

// cli POC

import (
	"cli-proposal/configs"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog/log"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type VaultConfig struct {
	url    string
	method string
	token  string
}

type GitLabConfig struct {
	url        string
	username   string
	password   string
	ssoEnabled bool
}

type ArgoCDConfig struct {
	url        string
	username   string
	password   string
	ssoEnabled bool
}

// checkRequiredEnvValues look for required environment variables
func checkRequiredEnvValues(config *configs.Config) error {

	elem := reflect.ValueOf(config).Elem()
	for i := 0; i < elem.NumField(); i++ {
		fieldValue := elem.Field(i).Interface()
		if fieldValue == "" {
			return fmt.Errorf("%s is not set", elem.Type().Field(i).Name)
		}
	}
	return nil
}

func banner() {
	docStyle := lipgloss.NewStyle().Padding(1, 2, 1, 2)

	highlight := lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#FFFFFF"}
	columnWidth := 170

	doc := strings.Builder{}

	banner := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Foreground(lipgloss.Color("#000000")).
		Background(highlight).
		Margin(0, 0, 0, 0).
		Padding(5, 5).
		Height(10).
		Width(columnWidth)

	kubefirst :=
		`
 __        ___.           _____.__                 __   
|  | ____ _\_ |__   _____/ ____\__|______  _______/  |_ 
|  |/ /  |  \ __ \_/ __ \   __\|  \_  __ \/  ___/\   __\
|    <|  |  / \_\ \  ___/|  |  |  ||  | \/\___ \  |  |  
|__|_ \____/|___  /\___  >__|  |__||__|  /____  > |__|  
     \/         \/     \/                     \/        

v1.7.0
`
	doc.WriteString(banner.Copy().Align(lipgloss.Center).Render(kubefirst))
	fmt.Println(docStyle.Render(doc.String()))

}

func vaultBuilder(isReady chan bool, config *VaultConfig) {

	sleepRandom()

	config.url = "https://vault.example.com/"
	config.method = "token"
	config.token = "jwt-token-here"

	isReady <- true
	return
}

func gitLabBuilder(isReady chan bool, config *GitLabConfig) {

	sleepRandom()

	config.url = "https://gitlab.example.com"
	config.username = "test"
	config.password = "test"
	config.ssoEnabled = true

	isReady <- true
}
func sleepRandom() {
	rand.Seed(time.Now().UnixNano())
	randomTime := 3 + rand.Intn(20-3+1)
	time.Sleep(time.Duration(randomTime) * time.Second)
}
func argoCDBuilder(isReady chan bool, config *ArgoCDConfig) {

	sleepRandom()

	config.url = "https://argocd.example.com"
	config.username = "test-argocd"
	config.password = "test-argocd"
	config.ssoEnabled = true

	isReady <- true
}

func showVaultConfig(config *VaultConfig) {
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("Vault is Ready!")
	fmt.Printf("URL: %s\n", config.url)
	fmt.Printf("username: %s\n", config.method)
	fmt.Printf("password: %s\n", config.token)
	fmt.Println(strings.Repeat("-", 70))
}

func showGitLabConfig(config *GitLabConfig) {
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("GitLab is Ready!")
	fmt.Printf("URL: %s\n", config.url)
	fmt.Printf("username: %s\n", config.username)
	fmt.Printf("password: %s\n", config.password)
	fmt.Printf("sso enabled: %t\n", config.ssoEnabled)
	fmt.Println(strings.Repeat("-", 70))
}

func showArgoCDConfig(config *ArgoCDConfig) {
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("ArgoCD is Ready!")
	fmt.Printf("URL: %s\n", config.url)
	fmt.Printf("username: %s\n", config.username)
	fmt.Printf("password: %s\n", config.password)
	fmt.Printf("sso enabled: %t\n", config.ssoEnabled)
	fmt.Println(strings.Repeat("-", 70))
}

func main() {

	banner()
	config := configs.ReadConfig()

	err := checkRequiredEnvValues(config)
	if err != nil {
		log.Err(err).Msg("")
		return
	}

	fmt.Println("working...")

	vaultBuilderChan := make(chan bool)
	var vaultConfig VaultConfig
	go vaultBuilder(vaultBuilderChan, &vaultConfig)

	gitLabBuilderChan := make(chan bool)
	var gitLabConfig GitLabConfig
	go gitLabBuilder(gitLabBuilderChan, &gitLabConfig)

	argoCDBuilderChan := make(chan bool)
	var argoCDConfig ArgoCDConfig
	go argoCDBuilder(argoCDBuilderChan, &argoCDConfig)

	for i := 0; i < 3; i++ {
		select {
		case isVaultReady := <-vaultBuilderChan:
			if isVaultReady {
				showVaultConfig(&vaultConfig)
			}
		case isGitLabReady := <-gitLabBuilderChan:
			if isGitLabReady {
				showGitLabConfig(&gitLabConfig)
			}
		case isArgoCDReady := <-argoCDBuilderChan:
			if isArgoCDReady {
				showArgoCDConfig(&argoCDConfig)
			}
		}
	}

}
