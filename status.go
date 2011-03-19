// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

import (
	"strings"
	"strconv"
)

type PlayState uint8

const (
	Playing PlayState = iota
	Paused
	Stopped
)

type Status struct {
	State          PlayState
	Volume         byte
	Playlist       int
	PlaylistLength int
	Song           int
	SongId         int
	NextSong       int
	NextSongId     int
	CrossFade      int
	Bitrate        int
	MixRampDB      float32
	MixRampDelay   float32
	Elapsed        float32
	Single         bool
	Repeat         bool
	Random         bool
	Consume        bool
	Time           []int
	Audio          []int
}

func readStatus(a Args) *Status {
	s := new(Status)
	s.Volume = a.U8("volume")
	s.Bitrate = a.I("bitrate")
	s.PlaylistLength = a.I("playlistlength")
	s.Time = splitI(a.S("time"), ":")
	s.Repeat = a.I("repeat") == 1
	s.NextSong = a.I("nextsong")
	s.MixRampDB = a.F32("mixrampdb")
	s.Playlist = a.I("playlist")
	s.Song = a.I("song")
	s.NextSongId = a.I("nextsongid")
	s.MixRampDelay = a.F32("mixrampdelay")
	s.Single = a.I("single") == 1
	s.CrossFade = a.I("xfade")
	s.Elapsed = a.F32("elapsed")
	s.SongId = a.I("songid")
	s.Random = a.I("random") == 1
	s.Consume = a.I("consume") == 1
	s.Audio = splitI(a.S("audio"), ":")

	switch a.S("state") {
	case "play":
		s.State = Playing
	case "stop":
		s.State = Stopped
	case "pause":
		s.State = Paused
	}
	return s
}

func splitI(v string, delim string) []int {
	var list []int
	var el []string

	if el = strings.Split(v, delim, -1); len(el) == 0 {
		return nil
	}

	list = make([]int, len(el))
	for i := range el {
		list[i], _ = strconv.Atoi(el[i])
	}

	return list
}
