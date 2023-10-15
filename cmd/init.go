/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	subFolders = []string{"01-Vorlesungen", "02-Uebungen", "03-Tutorium", "04-Literatur", "05-Sonstiges"}
)

type SemesterType string

const (
	SummerSemester SemesterType = "SoSe"
	WinterSemester SemesterType = "WiSe"
)

type Course struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Teacher      string `json:"teacher"`
}

type Semester struct {
	Name    string       `json:"name"`
	Years   []string     `json:"years"`
	Type    SemesterType `json:"type"`
	Courses []Course     `json:"courses"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := os.ReadFile(cmd.Flag("config").Value.String())
		if err != nil {
			fmt.Printf("could not read file: %v\n", err)
			return
		}

		var s *Semester
		err = json.Unmarshal(config, &s)
		if err != nil {
			fmt.Printf("could not unmarshall file: %v", err)
			os.Exit(1)
			return
		}

		log.Default().Printf("Starting to generate directories for %s", *s)
		dirs, err := generate(s)
		if err != nil {
			fmt.Printf("gensemester: error while generating: %v", err)
			os.Exit(1)
			return
		}

		for _, d := range dirs {
			err := os.Mkdir(d, os.ModePerm)
			if err != nil {
				fmt.Printf("gensemester: error while generating: %v", err)
				os.Exit(1)
				return
			}

			log.Default().Printf("Generated directory %s\n", d)
		}

		log.Default().Println("Finished generating")
	},
}

func generate(s *Semester) ([]string, error) {
	dirs := make([]string, 0)

	if isBlank(s.Name) {
		return []string{}, fmt.Errorf("name must be a non-blank string")
	}
	if containsBlank(s.Years) {
		return []string{}, fmt.Errorf("years must contain at least one non-blank string")
	}
	if isBlank(string(s.Type)) {
		return []string{}, fmt.Errorf("type must be a non-blank string")
	}
	if len(s.Courses) == 0 {
		return []string{}, fmt.Errorf("courses must contain at least one object")
	}

	root := path.Join(".", fmt.Sprintf("%s_%s_%s", s.Name, strings.Join(s.Years, "_"), s.Type))
	dirs = append(dirs, root)

	ws := regexp.MustCompile(`\s{1,}`)
	for i, c := range s.Courses {
		if isBlank(c.Name) {
			return []string{}, fmt.Errorf("courses[%d].name must be a non-blank string", i)
		}
		if isBlank(c.Abbreviation) {
			return []string{}, fmt.Errorf("courses[%d].abbreviation must be a non-blank string", i)
		}
		if isBlank(c.Teacher) {
			return []string{}, fmt.Errorf("courses[%d].teacher must be a non-blank string", i)
		}

		courseDir := path.Join(root, fmt.Sprintf("%s_%s_%s", c.Abbreviation, ws.ReplaceAllString(c.Name, "_"), c.Teacher))
		dirs = append(dirs, courseDir)

		for _, folder := range subFolders {
			dirs = append(dirs, path.Join(courseDir, folder))
		}
	}

	return dirs, nil
}

func isBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func containsBlank(list []string) bool {
	for _, l := range list {
		if isBlank(l) {
			return true
		}
	}

	return false
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().StringP("config", "c", "", "Path pointing to a config file")
}
