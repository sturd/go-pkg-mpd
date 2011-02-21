// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import "os"

// Toggles between play/pause.
func (this *Client) Toggle() (err os.Error) {
	var arg Args
	if arg, err = this.request("status"); err != nil {
		return
	}

	if arg["state"] == "play" {
		_, err = this.request("pause 1")
	} else {
		_, err = this.request("play")
	}
	return
}

// Sets crossfading (mixing) between songs.
// @time: Crossfade time in seconds.
func (this *Client) Crossfade(time int) (err os.Error) {
	_, err = this.request("crossfade %d", time)
	return
}

// Toggle pause on/off.
// @toggle: Specifies whether to pause or resume playback.
func (this *Client) Pause(toggle bool) (err os.Error) {
	v := 0
	if toggle {
		v = 1
	}
	_, err = this.request("pause %d", v)
	return
}

// Play the song at the specified position.
// @pos: Position of song to play.
func (this *Client) Play(pos int) (err os.Error) {
	_, err = this.request("play %d", pos)
	return
}

// Play the song with the specified id.
// @id: Id of the song to play.
func (this *Client) PlayId(id int) (err os.Error) {
	_, err = this.request("playid %d", id)
	return
}

// Skip to previous song.
func (this *Client) Previous() (err os.Error) {
	_, err = this.request("previous")
	return
}

// Skip to next song.
func (this *Client) Next() (err os.Error) {
	_, err = this.request("next")
	return
}

// Toggle random mode on/of
// @toggle: Specifies whether to used random or normal playback.
func (this *Client) Random(toggle bool) (err os.Error) {
	v := 0
	if toggle {
		v = 1
	}
	_, err = this.request("random %d", v)
	return
}

// Toggle repeat mode on/off.
// @toggle: Specifies whether to ise repeat or not.
func (this *Client) Repeat(toggle bool) (err os.Error) {
	v := 0
	if toggle {
		v = 1
	}
	_, err = this.request("repeat %d", v)
	return
}

// Skip to specific point in time in song at position @pos.
// @pos: Position of song.
// @time: Time in seconds to jump to.
func (this *Client) Seek(pos, time int) (err os.Error) {
	_, err = this.request("seek %d %d", pos, time)
	return
}

// Skip to specific point in time in song at position @pos.
// @pos: Id of song.
// @time: Time in seconds to jump to.
func (this *Client) SeekId(id, time int) (err os.Error) {
	_, err = this.request("seekid %d %d", id, time)
	return
}

// Volume adjustment. Allows setting of explicit volume value as well as a
// relative increase and decrease of current volume.
// @vol: New volume value in range 0-100.
// @relative: Indicates if our value is an absolute volume, or relative
// adjustment from the current volume.
func (this *Client) Volume(vol byte, relative bool) (err os.Error) {
	if relative {
		var a Args

		if a, err = this.request("status"); err != nil {
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

	_, err = this.request("setvol %d", vol)
	return
}

// Stop playback.
func (this *Client) Stop() (err os.Error) {
	_, err = this.request("stop")
	return
}
