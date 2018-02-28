// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

type Song struct {
	File         string
	Directory  	 string
	Artist       string
	AlbumArtist  string
	Album        string
	Title        string
	Genre        string
	LastModified string
	MBArtistID   string
	MBAArtistID  string
	MBAlbumID    string
	Date         int
	Id           int
	Pos          int
	Track        int
	Time         int
}

func readSong(a Args) *Song {
	s := new(Song)
	s.File = a.S("file")
	s.Directory = a.S("directory")
	s.Id = a.I("Id")
	s.Pos = a.I("Pos")
	s.Date = a.I("Date")
	s.Artist = a.S("Artist")
	s.AlbumArtist = a.S("AlbumArtist")
	s.Album = a.S("Album")
	s.Title = a.S("Title")
	s.Genre = a.S("Genre")
	s.Track = a.I("Track")
	s.Time = a.I("Time")
	s.LastModified = a.S("Last-Modified")
	s.MBArtistID = a.S("MUSICBRAINZ_ARTISTID")
	s.MBAArtistID = a.S("MUSICBRAINZ_ALBUMARTISTID")
	s.MBAlbumID = a.S("MUSICBRAINZ_ALBUMID")
	return s
}
