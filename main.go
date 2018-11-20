package main

import (
	"fmt"
	"os"
	"log"
	"os/exec"
	"time"

	"github.com/spf13/viper"
)


func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load %s: %v", viper.ConfigFileUsed(), err)
	} 

	keys := viper.GetStringSlice("team.name")
	path := viper.GetString("team.path")
	fmt.Println("Read team: ", keys)

	t := time.Now().Format("20060102150405")

	for _, k := range keys {
		p := path + "/" + k + "-" + t
		if err := os.Mkdir(p, 0700); err != nil {
			log.Fatalf("Failed to mkdir %s: %v", viper.ConfigFileUsed(), err)
		}

		proc(k, p)
	}
}

type Git struct {
	Url	string
	Path	string
	Name	string
}

func proc(key, path string) error {
	config := viper.GetStringMapString(key)
	fmt.Println("Start team: ", key, config)

	for k, v := range config {
		git := Git{
			Name: k,
			Path: path,
			Url: v,
		}
		if !git.Verify() {
			fmt.Printf("%s config invalid.\n", k)
			continue
		}

		if err := git.Run(); err != nil {
			fmt.Println("Git failed:", err)
		}
	}

	return nil
}

func (git *Git) Verify() bool {
	if git.Url == "" || git.Path == "" || git.Name == "" {
		return false
	}

	if git.Path[len(git.Path) - 1] == '/' {
		git.Path = git.Path[:len(git.Path) - 3]
	}
	
	return true
}

func (git *Git) Run() error {
	fmt.Println("Start git clone:", git.Url, git.Path + "/" +  git.Name)
	cmd := exec.Command("git", "clone", "--depth=1", git.Url, git.Path + "/" +  git.Name)
	log.Printf("Running command and waiting for it to finish...")
	return cmd.Run()
}


