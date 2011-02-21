// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

type Stats struct {
	Albums     int
	Artists    int
	Songs      int
	Playtime   int
	Uptime     int
	DbUpdate   int64
	DBPlaytime int64
}

func readStats(a Args) *Stats {
	s := new(Stats)
	s.Albums = a.I("albums")
	s.Artists = a.I("artists")
	s.Songs = a.I("songs")
	s.Playtime = a.I("playtime")
	s.Uptime = a.I("uptime")
	s.DbUpdate = a.I64("db_update")
	s.DBPlaytime = a.I64("db_playtime")
	return s
}
