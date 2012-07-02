// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

// Toggle toggles between play/pause.
func (c *Client) Toggle() (err error) {
	var arg Args
	if arg, err = c.request("status"); err != nil {
		return
	}

	if arg["state"] == "play" {
		_, err = c.request("pause 1")
	} else {
		_, err = c.request("play")
	}
	return
}

// Crossfade sets crossfading (mixing) between songs.
//
//     time: Crossfade time in seconds.
func (c *Client) Crossfade(time int) (err error) {
	_, err = c.request("crossfade %d", time)
	return
}

// Pause toggles pause on/off.
//
//     stoggle: Specifies whether to pause or resume playback.
func (c *Client) Pause(toggle bool) (err error) {
	v := 0
	if toggle {
		v = 1
	}
	_, err = c.request("pause %d", v)
	return
}

// Play plays the song at the specified position.
// 
//     pos: Position of song to play.
func (c *Client) Play(pos int) (err error) {
	_, err = c.request("play %d", pos)
	return
}

// PlayId plays the song with the specified id.
//
//     id: Id of the song to play.
func (c *Client) PlayId(id int) (err error) {
	_, err = c.request("playid %d", id)
	return
}

// Previous skips to previous song.
func (c *Client) Previous() (err error) {
	_, err = c.request("previous")
	return
}

// Next skips to next song.
func (c *Client) Next() (err error) {
	_, err = c.request("next")
	return
}

// Random toggles random mode on/of.
//
//     toggle: Specifies whether to used random or normal playback.
func (c *Client) Random(toggle bool) (err error) {
	v := 0
	if toggle {
		v = 1
	}
	_, err = c.request("random %d", v)
	return
}

// Repeat toggles repeat mode on/off.
//
//     toggle: Specifies whether to ise repeat or not.
func (c *Client) Repeat(toggle bool) (err error) {
	v := 0
	if toggle {
		v = 1
	}
	_, err = c.request("repeat %d", v)
	return
}

// Seek skips to specific point in time in a song at position `pos`.
//
//      pos: Position of song.
//     time: Time in seconds to jump to.
func (c *Client) Seek(pos, time int) (err error) {
	_, err = c.request("seek %d %d", pos, time)
	return
}

// SeekId skips to specific point in time in a song with the given id.
//
//       id: Id of song.
//     time: Time in seconds to jump to.
func (c *Client) SeekId(id, time int) (err error) {
	_, err = c.request("seekid %d %d", id, time)
	return
}

// Volume performs volume adjustment. Allows setting of explicit volume value
// as well as a relative increase and decrease of current volume.
//
//          vol: New volume value in range 0-100.
//     relative: Indicates if our value is an absolute volume, or relative
//               adjustment from the current volume.
func (c *Client) Volume(vol byte, relative bool) (err error) {
	if relative {
		var a Args

		if a, err = c.request("status"); err != nil {
			return
		}

		vol += a.U8("volume")
	}

	if vol < 0 {
		vol = 0
	}

	if vol > 100 {
		vol = 100
	}

	_, err = c.request("setvol %d", vol)
	return
}

// Stop stops playback.
func (c *Client) Stop() (err error) {
	_, err = c.request("stop")
	return
}
