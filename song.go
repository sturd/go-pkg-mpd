// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

type Song struct {
	File         string
	Id           int
	Pos          int
	Artist       string
	Album        string
	Title        string
	Track        int
	Time         int
	LastModified int64
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
