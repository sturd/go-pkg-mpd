// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

import "fmt"

// Add adds a single file from the database to the playlist. This command
// increments the playlist version by 1 for each song added to the playlist.
//
//     path: A single directory or file. If path is a directory, all files in it
//           are added recursively.
func (c *Client) Add(path string) (err error) {
	_, err = c.request("add %q", path)
	return
}

// AddId is the same as `add`, but this returns a playlistid and allows
// specifying a position at which to insert the file(s).
//
//     path: A single directory or file. If path is a directory, all files in it
//           are added recursively.
//      pos: The location at which to insert the file(s) into the playlist.
//           Supply -1 to insert at the end of the list.
func (c *Client) AddId(path string, pos int) (err error) {
	if pos > -1 {
		_, err = c.request("addid %q %d", path, pos)
	} else {
		_, err = c.request("addid %q", path)
	}
	return
}

// Clear clears the current playlist. Increments the playlist version by 1.
func (c *Client) Clear() (err error) {
	_, err = c.request("clear")
	return
}

// Current reports the metadata of the currently playing song.
func (c *Client) Current() (Args, error) {
	return c.request("currentsong")
}

// Delete deletes the specified song from the playlist. Increments the playlist
// version by 1.
//
//     pos: Position of the song in the playlist.
func (c *Client) Delete(pos int) (err error) {
	_, err = c.request("delete %d", pos)
	return
}

// DeleteId deletes the specified song from the playlist. Increments the
// playlist version by 1.
//
//     id: Id of the song to delete.
func (c *Client) DeleteId(id int) (err error) {
	_, err = c.request("deleteid %d", id)
	return
}

// Load loads the playlist name from the playlist directory, Increments the
// playlist version by the number of songs added.
//
//     name: Name of the playlist file *without* the path and file extension.
//           eg: `/path/to/all.m3u` -> `all`.
func (c *Client) Load(name string) (err error) {
	_, err = c.request("load %q", name)
	return
}

// Rename renames a playlist from `oldname` to `newname`. Names should be
// supplied  *without* the path and file extension.
// eg: `/path/to/all.m3u` -> `all`.
//
//     oldname: Current name of the playlist.
//     newname: New name of the playlist.
func (c *Client) Rename(oldname, newname string) (err error) {
	_, err = c.request("rename %q %q", oldname, newname)
	return
}

// Move moves a song at position `src` to position `dst`.
//
//     src: Source position.
//     dst: Target position.
func (c *Client) Move(src, dst int) (err error) {
	_, err = c.request("move %d %d", src, dst)
	return
}

// MoveId moves a song with id `src` to position `dst`.
//
//     src: Id of source song.
//     dst: Target position.
func (c *Client) MoveId(src, dst int) (err error) {
	_, err = c.request("moveid %d %d", src, dst)
	return
}

// PlaylistInfo reports metadata for songs in the playlist.
//
//     pos: An optional number that specifies a single song to display
//          information for. Specify -1 to report for all songs.
func (c *Client) PlaylistInfo(pos int) (list []*Song, err error) {
	var a []Args
	var str string

	if pos == -1 {
		str = "playlistinfo"
	} else {
		str = fmt.Sprintf("playlistinfo %d", pos)
	}

	if a, err = c.requestList(str); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// PlaylistChanges reports changed songs in the playlist since version.
//
//     version: The playlist version to display changed songs for.
func (c *Client) PlaylistChanges(version int) (list []*Song, err error) {
	var a []Args

	if a, err = c.requestList("plchanges %d", version); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// PlaylistRm removes the playlist called name from the playlist directory.
//
//     name: Name of the playlist file *without* the path and file extension.
//           eg: `/path/to/all.m3u` -> `all`.
func (c *Client) PlaylistRm(name string) (err error) {
	_, err = c.request("rm %q", name)
	return
}

// Save saves the current playlist to name in the playlist directory.
//
//     name: Name of the playlist file *without* the path and file extension.
//           eg: `/path/to/all.m3u` -> `all`.
func (c *Client) Save(name string) (err error) {
	_, err = c.request("save %q", name)
	return
}

// Shuffle shuffles the current playlist and increments playlist version by 1.
func (c *Client) Shuffle() (err error) {
	_, err = c.request("shuffle")
	return
}

// Swap swaps positions of songs at positions `src` and `dst`. Increments
// playlist version by 1.
//
//     src: Source position.
//     dst: Target position.
func (c *Client) Swap(src, dst int) (err error) {
	_, err = c.request("swap %d %d", src, dst)
	return
}

// SwapId swaps positions of songs with the specified IDs. Increments playlist
// version by 1.
//
//     src: Source id.
//     dst: Target id.
func (c *Client) SwapId(src, dst int) (err error) {
	_, err = c.request("swapid %d %d", src, dst)
	return
}

// ListPlaylistFiles reports files in playlist named `name`.
//
//     name: Name of the playlist file *without* the path and file extension.
//           eg: `/path/to/all.m3u` -> `all`.
func (c *Client) ListPlaylistFiles(name string) (v []string, err error) {
	var a []Args
	if a, err = c.requestList("listplaylist %q", name); err != nil {
		return
	}

	var k string
	for _, m := range a {
		for _, k = range m {
			v = append(v, k)
		}
	}

	return
}

// ListPlaylists lists all playlists stored on the server
//
func (c *Client) ListPlaylists() (p []*Playlist, err error) {
	var a []Args
	if a, err = c.requestList("listplaylists"); err != nil {
		return
	}

	for _, k := range a {
		p = append(p, readPlaylist(k))
	}

	return
}

// ListPlaylistSongs reports songs in playlist named `name`.
//
//     name: Name of the playlist file *without* the path and file extension.
//           eg: `/path/to/all.m3u` -> `all`.
func (c *Client) ListPlaylistSongs(name string) (list []*Song, err error) {
	var a []Args

	if a, err = c.requestList("listplaylistinfo %q", name); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// PlaylistAdd adds `path` to playlist `name`.
//
//     name: Name of the playlist file *without* the path and file extension.
//           eg: `/path/to/all.m3u` -> `all`.
//     path: Path of file(s) to add to the given playlist.
func (c *Client) PlaylistAdd(name, path string) (err error) {
	_, err = c.request("playlistadd %q %q", name, path)
	return
}

// PlaylistClear clears playlist with given `name`.
//
//     name: Name of the playlist file *without* the path and file extension.
//           eg: `/path/to/all.m3u` -> `all`.
func (c *Client) PlaylistClear(name string) (err error) {
	_, err = c.request("playlistclear %q", name)
	return
}

// PlaylistDelete deletes the song with given `id` from playlist `name`.
//
//      name: Name of the playlist file *without* the path and file extension.
//            eg: `/path/to/all.m3u` -> `all`.
//        id: ID of song to delete.
func (c *Client) PlaylistDelete(name string, id int) (err error) {
	_, err = c.request("playlistdelete %q %d", name, id)
	return
}

// PlaylistMove moves the song with given `id` in playlist `name` to
// position `pos`.
//
//     name: Name of the playlist file *without* the path and file extension.
//           eg: `/path/to/all.m3u` -> `all`.
//       id: ID of song to move.
//      pos: Position to move song to.
func (c *Client) PlaylistMove(name string, id, pos int) (err error) {
	_, err = c.request("playlistmove %q %d %d", name, id, pos)
	return
}

// PlaylistSearch performs a case-insensitive playlist search.
//
//      tag: Tag to search in.
//     term: Term to search for in tag.
func (c *Client) PlaylistSearch(tag, term string) (list []*Song, err error) {
	var a []Args

	if a, err = c.requestList("playlistsearch %q %q", tag, term); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}
