/*
# GetMovieRating go file accepts movie name as argument from command line and returns Rotten Tomatoes rating for the movie on command line
# Author: Manju Subramanyam
# Date: Jan 15th 2020
# Copyright (c) 2018 Cisco Systems. All rights reserved.
*/

package main

import (
	"testing"
)

func Test_inputReqArg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{want: "star"},
		{want: "red"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inputReqArg(); got == tt.want {
				t.Errorf("inputReqArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spaceReplace(t *testing.T) {
	type args struct {
		replaceString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{replaceString: "star wars"}, want: "star%20wars"},
		{args: args{replaceString: "red sea"}, want: "red%20sea"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spaceReplace(tt.args.replaceString); got != tt.want {
				t.Errorf("spaceReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidSource(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args: args{source: "Rotten Tomatoes"}, want: true},
		{args: args{source: "Rotten Tomatoes false"}, want: false},
		{args: args{source: " "}, want: false},
		{args: args{source: "Rotten Tomatoes ***"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSource(tt.args.source); got != tt.want {
				t.Errorf("IsValidSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
