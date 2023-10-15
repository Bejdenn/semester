package cmd

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		s *Semester
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "Valid_Values", args: args{s: &Semester{Name: "SEMESTER", Years: []string{"2023"}, Type: "SoSe", Courses: []Course{
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
			"SEMESTER_2023_SoSe/EC_Example_Course_Doe/05-Sonstiges"}, wantErr: false},
		{name: "Empty_Name", args: args{s: &Semester{Name: "", Years: []string{""}, Type: "", Courses: []Course{}}}, want: []string{}, wantErr: true},
		{name: "Empty_Years", args: args{s: &Semester{Name: "SEMESTER", Years: []string{""}, Type: "", Courses: []Course{}}}, want: []string{}, wantErr: true},
		{name: "Empty_Type", args: args{s: &Semester{Name: "SEMESTER", Years: []string{"2023"}, Type: "", Courses: []Course{}}}, want: []string{}, wantErr: true},
		{name: "Empty_Courses", args: args{s: &Semester{Name: "SEMESTER", Years: []string{"2023"}, Type: "SoSe", Courses: []Course{}}}, want: []string{}, wantErr: true},
		{name: "Empty_Course_Name", args: args{s: &Semester{Name: "SEMESTER", Years: []string{"2023"}, Type: "SoSe", Courses: []Course{
			{
				Name:         "",
				Abbreviation: "",
				Teacher:      "",
			},
		}}}, want: []string{}, wantErr: true},
		{name: "Empty_Course_Abbreviation", args: args{s: &Semester{Name: "SEMESTER", Years: []string{"2023"}, Type: "SoSe", Courses: []Course{
			{
				Name:         "Example Course",
				Abbreviation: "",
				Teacher:      "",
			},
		}}}, want: []string{}, wantErr: true},
		{name: "Empty_Course_Teacher", args: args{s: &Semester{Name: "SEMESTER", Years: []string{"2023"}, Type: "SoSe", Courses: []Course{
			{
				Name:         "Example Course",
				Abbreviation: "EC",
				Teacher:      "",
			},
		}}}, want: []string{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generate(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
