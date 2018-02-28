// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

import (
	"errors"
	"fmt"
)

// Find finds songs in the database with a case sensitive, exact match to `term`.
//
//      tag: This is the type of metadata you wish to use to refine the search.
//           Examples would be album, artist, title or any.
//     term: This is the value that is being searched for in tag.
func (c *Client) Find(tag, term string) (list []*Song, err error) {
	var a []Args

	if a, err = c.requestList("find %q %q", tag, term); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// List reports all metadata of type `tag1`.
//
//     tag1: The type of metadata to list.
//     tag2: Used together with `term`. This specifies to look for `tag2` in
//           the list of `tag1` results.
//     term: Used together with `tag2`. This specifies to look for matches
//           of term in the list of `tag2` results.
func (c *Client) List(tag1, tag2, term string) (list []*Song, err error) {
	var a []Args

	str := fmt.Sprintf("list %q", tag1)

	if tag2 > "" {
		if len(term) == 0 {
			return nil, errors.New("Missing parameter @term if parameter @tag2 has been supplied.")
		}

		str += fmt.Sprintf(" %q %q", tag1, tag2, term)
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

// ListFiles reports all directories and filenames in path recursively.
//
//     path: An optional directory path to act as the root of the list.
//           If omitted, we assume the music root as defined in mpd.conf.
func (c *Client) ListFiles(path string) (list []string, err error) {
	var a []Args

	str := "listall"
	if path != "" {
		str += fmt.Sprintf(" %q", path)
	}

	if a, err = c.requestList(str); err != nil {
		return
	}

	list = make([]string, 0, len(a))
	for _, m := range a {
		list = append(list, m["file"])
	}

	return
}

// ListInfo reports all information in the database about all music files
// in path recursively.
//
//     path: An optional directory path to act as the root of the list.
//           If omitted, we assume the music root as defined in mpd.conf.
func (c *Client) ListInfo(path string) (list []*Song, err error) {
	var a []Args
	str := "listallinfo"

	if len(path) > 0 {
		str += fmt.Sprintf(" %q", path)
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

// Ls reports a list of files/directories in `path`, from the database.
//
//     path: An optional directory path to act as the root of the list.
//           If omitted, we assume the music root as defined in mpd.conf.
func (c *Client) Ls(path string) (list []*Song, err error) {
	var a []Args

	str := "lsinfo"

	if path != "" {
		str += fmt.Sprintf(" %q", path)
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

// Search finds songs in the database with a case insensitive match to `term`.
//
//      tag: This is the type of metadata you wish to use to refine the search.
//     term: This is the value that is being searched for in tag.
func (c *Client) Search(tag, term string) (list []*Song, err error) {
	var a []Args

	if a, err = c.requestList("search %q %q", tag, term); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// Count reports the number of songs and their total playtime in the
// database matching `term`.
//
//      tag: This is the type of metadata you wish to use to refine the search.
//     term: This is the value that is being searched for in tag.
func (c *Client) Count(tag, term string) (songs, playtime int, err error) {
	var a []Args

	if a, err = c.requestList("count %q %q", tag, term); err != nil || len(a) == 0 {
		return
	}

	songs = a[0].I("songs")
	playtime = a[0].I("playtime")
	return
}
