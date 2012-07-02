// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

type Stats struct {
	Albums     int
	Artists    int
	Songs      int
	Playtime   int
	Uptime     int
	DbUpdate   int64
	DbPlaytime int64
}

func readStats(a Args) *Stats {
	s := new(Stats)
	s.Albums = a.I("albums")
	s.Artists = a.I("artists")
	s.Songs = a.I("songs")
	s.Playtime = a.I("playtime")
	s.Uptime = a.I("uptime")
	s.DbUpdate = a.I64("db_update")
	s.DbPlaytime = a.I64("db_playtime")
	return s
}
