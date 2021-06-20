package person

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestOldestAge(t *testing.T) {
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
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "with one person",
			args:    args{p1},
			want:    34,
			wantErr: false,
		},
		{
			name:    "with several persons",
			args:    args{p2},
			want:    43,
			wantErr: false,
		},
		{
			name:    "without persons",
			args:    args{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OldestAge(tt.args.p...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OldestAge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OldestAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOldestObject(t *testing.T) {
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
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "with one person",
			args:    args{i1},
			want:    p1,
			wantErr: false,
		},
		{
			name:    "with several persons (customer oldest)",
			args:    args{i2},
			want:    p2,
			wantErr: false,
		},
		{
			name:    "with several persons (employee oldest)",
			args:    args{i3},
			want:    p4,
			wantErr: false,
		},
		{
			name:    "without persons",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OldestObject(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OldestObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OldestObject() = %v, want %v", got, tt.want)
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
