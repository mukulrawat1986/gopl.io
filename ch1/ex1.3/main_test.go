package main

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	args   = []string{"exec", "arg0", "arg1", "arg2", "arg3", "arg4", "arg5"}
	buffer = &bytes.Buffer{}
)

func Test_echo1(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "echo1",
			args: args{
				args: []string{"execute", "hello", "world"},
			},
			wantW: "hello world\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			echo1(w, tt.args.args)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("echo1() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_echo2(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "echo2",
			args: args{
				args: []string{"execute", "hello", "world"},
			},
			wantW: "hello world\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			echo2(w, tt.args.args)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("echo2() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_echo3(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "echo3",
			args: args{
				args: []string{"execute", "hello", "world"},
			},
			wantW: "hello world\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			echo3(w, tt.args.args)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("echo3() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func BenchmarkEcho1(b *testing.B) {
	fmt.Println(args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo1(buffer, args)
		}
	} else {
		fmt.Println("No args given")
	}
}

func BenchmarkEcho2(b *testing.B) {
	fmt.Println(args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo2(buffer, args)
		}
	} else {
		fmt.Println("No args given")
	}
}

func BenchmarkEcho3(b *testing.B) {
	fmt.Println(args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo3(buffer, args)
		}
	} else {
		fmt.Println("No args given")
	}
}
