package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	subFolders = []string{"01-Vorlesungen", "02-Uebungen", "03-Tutorium", "04-Literatur", "05-Sonstiges"}
)

type SemesterType string

const (
	SummerSemester SemesterType = "SoSe"
	WinterSemester SemesterType = "WiSe"
)

type Semester struct {
	Name    string
	Years   []string
	Type    SemesterType
	Courses []struct {
		Name         string
		Abbreviation string
		Teacher      string
	}
}

func (s Semester) String() string {
	cs := make([]string, 0)
	for _, c := range s.Courses {
		cs = append(cs, c.Abbreviation)
	}

	return fmt.Sprintf("%s - %s with courses [%s]", s.Name, strings.Join(s.Years, " - "), strings.Join(cs, ", "))
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf(`usage: gensemester <config>
where <config> is a JSON configuration of the semester
`)
		os.Exit(1)
		return
	}

	config, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Printf("could not read file: %v", err)
		os.Exit(1)
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
	err = generate(s)
	if err != nil {
		fmt.Println(fmt.Errorf("gensemester: error while generating: %v", err))
		return
	}
	log.Default().Println("Finished generating")
}

func generate(s *Semester) error {
	root, err := mkdir(".", fmt.Sprintf("%s_%s_%s", s.Name, strings.Join(s.Years, "_"), s.Type))
	if err != nil {
		return err
	}

	ws := regexp.MustCompile(`\s{1,}`)
	for _, c := range s.Courses {
		cDir, err := mkdir(*root, fmt.Sprintf("%s_%s_%s", c.Abbreviation, ws.ReplaceAllString(c.Name, "_"), c.Teacher))
		if err != nil {
			return err
		}

		for _, folder := range subFolders {
			_, err = mkdir(*cDir, folder)
			if err != nil {
				return err
			}
		}

		log.Default().Printf("Generated directory for %s", c.Abbreviation)
	}

	return nil
}

// mkdir creates a new directory with name inside the specified path.
func mkdir(path, name string) (*string, error) {
	dir := path + "/" + name
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not make directory with name %s: %v", name, err)
	}

	return &dir, nil
}
