package person

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestHighestAge(t *testing.T) {
	c1 := Customer{age: 34}
	c2 := Customer{age: 43}
	e3 := Employee{age: 27}

	var p1, p2 []Person
	p1 = append(p1, &c1)
	p2 = append(p2, &c1, &c2, &e3)

	type args struct {
		p []Person
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with one person",
			args: args{p1},
			want: 34,
		},
		{
			name: "with several persons",
			args: args{p2},
			want: 43,
		},
		{
			name: "without persons",
			args: args{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HighestAge(tt.args.p...)
			if got != tt.want {
				t.Errorf("HighestAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEldestPerson(t *testing.T) {
	p1 := Customer{age: 34}
	p2 := Customer{age: 43}
	p3 := Employee{age: 27}
	p4 := Employee{age: 35}

	var i1, i2, i3 []interface{}
	i1 = append(i1, p1)
	i2 = append(i1, p1, p2, p3)
	i3 = append(i1, p1, p3, p4)

	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "with one person",
			args: args{i1},
			want: p1,
		},
		{
			name: "with several persons (customer oldest)",
			args: args{i2},
			want: p2,
		},
		{
			name: "with several persons (employee oldest)",
			args: args{i3},
			want: p4,
		},
		{
			name: "without persons",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EldestPerson(tt.args.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EldestPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	var i1, i2, i3, i4 []interface{}
	i1 = append(i1, 1, "string")
	i2 = append(i2, errors.New("new error"), true, 'a')
	i3 = append(i2, 0.5, "")
	i4 = append(i3, 1, false, "str1", " ", "str2")

	type args struct {
		args []interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name:  "Test #1",
			args:  args{i1},
			wantW: "string",
		},
		{
			name:  "Test #2",
			args:  args{i2},
			wantW: "",
		},
		{
			name:  "Test #2",
			args:  args{i3},
			wantW: "",
		},
		{
			name:  "Test #3",
			args:  args{i4},
			wantW: "str1 str2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Print(w, tt.args.args...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Print() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
