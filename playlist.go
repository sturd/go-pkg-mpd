package mpd

type Playlist struct {
	Name string
	LastModified string
}

func readPlaylist(a Args) *Playlist {
	p := new(Playlist)
	p.Name = a.S("playlist")
	p.LastModified = a.S("Last-Modified")
	return p
}