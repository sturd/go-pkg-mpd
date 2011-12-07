// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

type Song struct {
	File         string
	Artist       string
	Album        string
	Title        string
	LastModified int64
	Id           int
	Pos          int
	Track        int
	Time         int
}

func readSong(a Args) *Song {
	s := new(Song)
	s.File = a.S("file")
	s.Id = a.I("Id")
	s.Pos = a.I("Pos")
	s.Artist = a.S("Artist")
	s.Album = a.S("Album")
	s.Title = a.S("Title")
	s.Track = a.I("Track")
	s.Time = a.I("Time")
	s.LastModified = a.I64("Last-Modified")
	return s
}
