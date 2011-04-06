// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	var err os.Error
	var c *Client

	if c, err = Dial("127.0.0.1:6600", ""); err != nil {
		t.Error(err.String())
		return
	}

	defer c.Close()

	if status, err := c.Status(); err != nil {
		t.Error(err.String())
		return
	} else {
		// do something with our status data.
		_ = status
	}

	if songs, err := c.PlaylistSearch("artist", "tool"); err != nil {
		t.Error(err.String())
		return
	} else {
		var song *Song
		for _, song = range songs {
			// do something with our song data.
			_ = song
		}
	}
}
