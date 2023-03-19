package main

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		s *Semester
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test regular semester", args: args{s: &Semester{Name: "SEMESTER", Years: []string{"2023"}, Type: "SoSe", Courses: []struct {
			Name         string
			Abbreviation string
			Teacher      string
		}{
			{
				Name:         "Example Course",
				Abbreviation: "EC",
				Teacher:      "Doe",
			},
		}}}, want: []string{
			"SEMESTER_2023_SoSe",
			"SEMESTER_2023_SoSe/EC_Example_Course_Doe",
			"SEMESTER_2023_SoSe/EC_Example_Course_Doe/01-Vorlesungen",
			"SEMESTER_2023_SoSe/EC_Example_Course_Doe/02-Uebungen",
			"SEMESTER_2023_SoSe/EC_Example_Course_Doe/03-Tutorium",
			"SEMESTER_2023_SoSe/EC_Example_Course_Doe/04-Literatur",
			"SEMESTER_2023_SoSe/EC_Example_Course_Doe/05-Sonstiges"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
