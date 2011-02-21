// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

type Output struct {
	Id      int
	Name    string
	Enabled bool
}

func readOutput(a Args) *Output {
	s := new(Output)
	s.Id = a.I("outputid")
	s.Name = a.S("outputname")
	s.Enabled = a.I("outputenabled") == 1
	return s
}
