// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

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
