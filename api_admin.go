// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

// Turns an audio-output source off.
// @id: Id of the output device. Use the 'outputs' command to find all valid Ids
func (this *Client) DisableOutput(id int) (err error) {
	_, err = this.request("disableoutput %d", id)
	return
}

// Turns an audio-output source on.
// @id: Id of the output device. Use the 'outputs' command to find all valid Ids.
func (this *Client) EnableOutput(id int) (err error) {
	_, err = this.request("enableoutput %d", id)
	return
}

// Stops MPD from running; in a safe way. Writes a state file if defined.
func (this *Client) Kill() (err error) {
	_, err = this.request("kill")
	return
}

// Scans the music directory as defined in the MPD configuration file's
// music_directory  setting. Adds new files and their metadata (if any) to the
// MPD database and removes files and metadata from the database that are no
// longer in the directory.
// @path: An optional argument that picks an exact directory or file to
// update, otherwise the root of the music_directory in your MPD configuration
// file is assumed.
func (this *Client) Update(path string) (err error) {
	if len(path) == 0 {
		_, err = this.request("update")
	} else {
		_, err = this.request("update \"%s\"", path)
	}
	return
}
