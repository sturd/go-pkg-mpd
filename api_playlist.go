// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"fmt"
	"os"
)

// Add a single file from the database to the playlist. This command increments
// the playlist version by 1 for each song added to the playlist.
// @path: A single directory or file. If this is a directory, all files in it
// are added recursively.
func (this *Client) Add(path string) (err os.Error) {
	_, err = this.request("add \"%s\"", path)
	return
}

// Same as 'add', but this returns a playlistid and allows specifying a position
// at which to insert the file(s).
// @path: A single directory or file. If this is a directory, all files in it
// are added recursively.
// @pos: The location at which to insert the file(s) into the playlist. Pass
// -1 to insert at the end of the list.
func (this *Client) AddId(path string, pos int) (err os.Error) {
	if pos > -1 {
		_, err = this.request("addid \"%s\" %d", path, pos)
	} else {
		_, err = this.request("addid \"%s\"", path)
	}
	return
}

// Clears the current playlist. Increments the playlist version by 1.
func (this *Client) Clear() (err os.Error) {
	_, err = this.request("clear")
	return
}

// Reports the metadata of the currently playing song.
func (this *Client) Current() (Args, os.Error) {
	return this.request("currentsong")
}

// Deletes the specified song from the playlist. increments the playlist version by 1.
// @pos: Position of the song in the playlist.
func (this *Client) Delete(pos int) (err os.Error) {
	_, err = this.request("delete %d", pos)
	return
}

// Deletes the specified song from the playlist. increments the playlist version by 1.
// @pos: Id of the song to delete.
func (this *Client) DeleteId(id int) (err os.Error) {
	_, err = this.request("deleteid %d", id)
	return
}

// Load the playlist @name from the playlist directory, Increments the playlist
// version by the number of songs added.
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
func (this *Client) Load(name string) (err os.Error) {
	_, err = this.request("load \"%s\"", name)
	return
}

// Renames a playlist from @oldname to @newname. Names should be supplied 
// *without* the path and file extension. eg: '/path/to/all.m3u' -> 'all'.
// @oldname: Current name of the playlist.
// @newname: New name of the playlist.
func (this *Client) Rename(oldname, newname string) (err os.Error) {
	_, err = this.request("rename \"%s\" \"%s\"", oldname, newname)
	return
}

// Moves a song with id @src to position @dest.
// @src: Source position.
// @dst: Target position.
func (this *Client) Move(src, dst int) (err os.Error) {
	_, err = this.request("move %d %d", src, dst)
	return
}

// Moves a song with id position @src to position @dest.
// @src: Id of source song.
// @dst: Target position.
func (this *Client) MoveId(src, dst int) (err os.Error) {
	_, err = this.request("moveid %d %d", src, dst)
	return
}

// Reports metadata for songs in the playlist.
// @pos: An optional number that specifies a single song to display information
// for. Specify -1 to report for all songs.
func (this *Client) PlaylistInfo(pos int) (list []*Song, err os.Error) {
	var a []Args
	var str string

	if pos == -1 {
		str = "playlistinfo"
	} else {
		str = fmt.Sprintf("playlistinfo %d", pos)
	}

	if a, err = this.requestList(str); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// Reports changed songs in the playlist since @version.
// @version: The playlist version to display changed songs for.
func (this *Client) PlaylistChanges(version int) (list []*Song, err os.Error) {
	var a []Args

	if a, err = this.requestList("plchanges %d", version); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// Removes the playlist called @name from the playlist directory.
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
func (this *Client) PlaylistRm(name string) (err os.Error) {
	_, err = this.request("rm \"%s\"", name)
	return
}

// Saves the current playlist to @name in the playlist directory.
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
func (this *Client) Save(name string) (err os.Error) {
	_, err = this.request("save \"%s\"", name)
	return
}

// Shuffles the current playlist, increments playlist version by 1.
func (this *Client) Shuffle() (err os.Error) {
	_, err = this.request("shuffle")
	return
}

// Swap positions of songs at positions @pos1 and @pos2. Increments playlist version by 1.
// @src: Source position.
// @dst: Target position.
func (this *Client) Swap(src, dst int) (err os.Error) {
	_, err = this.request("swap %d %d", src, dst)
	return
}

// Swap positions of songs with the specified IDs. Increments playlist version by 1.
// @src: Source ID.
// @dst: Target ID.
func (this *Client) SwapId(src, dst int) (err os.Error) {
	_, err = this.request("swapid %d %d", src, dst)
	return
}

// Reports files in playlist named @name.
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
func (this *Client) ListPlaylistFiles(name string) (v []string, err os.Error) {
	var a []Args
	if a, err = this.requestList("listplaylist \"%s\"", name); err != nil {
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

// Reports songs in playlist named @name.
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
func (this *Client) ListPlaylistSongs(name string) (list []*Song, err os.Error) {
	var a []Args

	if a, err = this.requestList("listplaylistinfo \"%s\"", name); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// Adds @path to playlist @name.
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
// @path: Path of file(s) to add to the given playlist.
func (this *Client) PlaylistAdd(name, path string) (err os.Error) {
	_, err = this.request("playlistadd \"%s\" \"%s\"", name, path)
	return
}

// Clear playlist with given @name.
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
func (this *Client) PlaylistClear(name string) (err os.Error) {
	_, err = this.request("playlistclear \"%s\"", name)
	return
}

// Deletes song with given @id from playlist @name.
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
// @id: ID of song to delete.
func (this *Client) PlaylistDelete(name string, id int) (err os.Error) {
	_, err = this.request("playlistdelete \"%s\" %d", name, id)
	return
}

// Moves song with given @id in playlist @name to position @pos
// @name: Name of the playlist file *without* the path and file extension.
// eg: '/path/to/all.m3u' -> 'all'.
// @id: ID of song to move.
// @pos: Position to move song to.
func (this *Client) PlaylistMove(name string, id, pos int) (err os.Error) {
	_, err = this.request("playlistmove \"%s\" %d %d", name, id, pos)
	return
}

// Case-insensitive playlist search.
// @tag: Tag to search in.
// @term: Term to search for in @tag.
func (this *Client) PlaylistSearch(tag, term string) (list []*Song, err os.Error) {
	var a []Args

	if a, err = this.requestList("playlistsearch \"%s\" \"%s\"", tag, term); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}
