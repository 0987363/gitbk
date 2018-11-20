package main

import (
	"fmt"
	"log"
	"reflect"
	"os/exec"

	"github.com/spf13/viper"
)


func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load %s: %v", viper.ConfigFileUsed(), err)
	} 

	keys := viper.GetStringSlice("team")
	fmt.Println("Read team: ", keys)

	for _, k := range keys {
		proc(k)
	}

}

type Git struct {
	Url	string
	Path	string
	Name	string
}

func proc(key string) error {
	config := viper.GetStringMap(key)
	fmt.Println("Start team: ", key)

	for k, v := range config {
		vo := reflect.ValueOf(v)
		if vo.Kind() != reflect.Map {
			continue
		}

		m, ok := v.(map[string]interface{})
		if !ok {
			fmt.Println("Convert map failed:", v)
			continue
		}

		var git Git
		for field, data := range m {
			if field == "url" {
				git.Url = data.(string)
				continue
			}
			if field == "path" {
				git.Path = data.(string)
				continue
			}
			if field == "name" {
				git.Name = data.(string)
				continue
			}
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
	cmd := exec.Command("git", "clone", "--depth=1", git.Url, git.Path + "/" +  git.Name)
	log.Printf("Running command and waiting for it to finish...")
	return cmd.Run()
}


