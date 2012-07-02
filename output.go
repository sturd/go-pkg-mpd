// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

type Output struct {
	Name    string
	Id      int
	Enabled bool
}

func readOutput(a Args) *Output {
	s := new(Output)
	s.Id = a.I("outputid")
	s.Name = a.S("outputname")
	s.Enabled = a.I("outputenabled") == 1
	return s
}
