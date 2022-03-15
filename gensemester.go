package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	subFolders = []string{"01-Vorlesungen", "02-Uebungen", "03-Tutorium", "04-Literatur", "05-Sonstiges"}
)

var ErrNotEnoughArgs = errors.New("usage: gensemester <semester-key> <courses>" + "\n" +
	"where <courses> is a comma-separated enumeration of your courses (atleast one)")

func main() {
	if len(os.Args) < 3 {
		fmt.Println(ErrNotEnoughArgs)
		return
	}

	err := generateSemesterDirs()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func generateSemesterDirs() (err error) {
	semesterKey := replaceWhitespace(os.Args[1])

	semesterDir, err := mkdir("./", semesterKey)
	if err != nil {
		return
	}

	courses := os.Args[2:]
	for _, c := range courses {
		err := generateCourseFolder(semesterDir, c)
		if err != nil {
			log.Println(err)
		}
	}

	return
}

func replaceWhitespace(str string) string {
	return strings.ReplaceAll(str, " ", "_")
}

func mkdir(path, dirName string) (dirPath string, err error) {
	dirPath = path + "/" + dirName

	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		err = fmt.Errorf("could not make directory with name %s: %v", dirName, err)
		return
	}

	return
}

func generateCourseFolder(path, courseName string) (err error) {
	c := replaceWhitespace(courseName)

	subPath, err := mkdir(path, c)
	if err != nil {
		return err
	}

	for _, folder := range subFolders {
		_, err = mkdir(subPath, folder)
		if err != nil {
			return
		}
	}

	return
}
