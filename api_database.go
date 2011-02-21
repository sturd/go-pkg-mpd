// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"os"
	"fmt"
)

// Finds songs in the database with a case sensitive, exact match to @term.
// @tag: This is the type of metadata you wish to use to refine the search.
// Examples would be album, artist, title or any.
// @term: This is the value that is being searched for in @tag.
func (this *Client) Find(tag, term string) (list []*Song, err os.Error) {
	var a []Args

	if a, err = this.requestList("find \"%s\" \"%s\"", tag, term); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// Reports all metadata of @tag1
// @tag1: This lists all metadata of @tag1.
// @tag2: Used together with @term. This specifies to look for @tag2 in the list
// of @tag1 results.
// @term: Used together with @tag2. This specifies to look for matches of @term
// in the list of @tag2 results.
func (this *Client) List(tag1, tag2, term string) (list []*Song, err os.Error) {
	var str string
	var a []Args

	if tag2 == "" {
		str = fmt.Sprintf("list \"%s\"", tag1)
	} else {
		if len(term) == 0 {
			return nil, os.NewError("Missing parameter @term if parameter @tag2 has been supplied.")
		}
		str = fmt.Sprintf("list \"%s\" \"%s\" \"%s\"", tag1, tag2, term)
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

// Reports all directories and filenames in @path recursively.
// @path: An optional directory path to act as the root of the list. If omitted,
// we assume the music root as defined in mpd.conf.
func (this *Client) ListFiles(path string) (list []string, err os.Error) {
	var a []Args

	str := "listall"
	if path != "" {
		str = fmt.Sprintf("listall \"%s\"", path)
	}

	if a, err = this.requestList(str); err != nil {
		return
	}

	list = make([]string, 0, len(a))
	for _, m := range a {
		list = append(list, m["file"])
	}

	return
}

// Reports all information in the database about all music files in @path
// recursively.
// @path: An optional directory path to act as the root of the list. If omitted,
// we assume the music root as defined in mpd.conf.
func (this *Client) ListInfo(path string) (list []*Song, err os.Error) {
	var a []Args
	str := "listallinfo"

	if len(path) > 0 {
		str = fmt.Sprintf("listallinfo \"%s\"", path)
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

// Reports list of files/directories in @path, from the database.
// @path: An optional directory path to act as the root of the list. If omitted,
// we assume the music root as defined in mpd.conf.
func (this *Client) Ls(path string) (list []string, err os.Error) {
	var a []Args
	var str, v string

	if path != "" {
		str = fmt.Sprintf("lsinfo \"%s\"", path)
	} else {
		str = "lsinfo"
	}

	if a, err = this.requestList(str); err != nil {
		return
	}

	list = make([]string, 0, len(a))
	for _, m := range a {
		for _, v = range m {
			list = append(list, v)
		}
	}

	return
}

// Finds songs in the database with a case insensitive match to @term.
// @tag: This is the type of metadata you wish to use to refine the search.
// @term: This is the value that is being searched for in @tag.
func (this *Client) Search(tag, term string) (list []*Song, err os.Error) {
	var a []Args

	if a, err = this.requestList("search \"%s\" \"%s\"", tag, term); err != nil {
		return
	}

	list = make([]*Song, 0, len(a))
	for _, m := range a {
		list = append(list, readSong(m))
	}

	return
}

// Reports the number of songs and their total playtime in the database matching @what.
// @tag: This is the type of metadata you wish to use to refine the search.
// @term: This is the value that is being searched for in @tag.
func (this *Client) Count(tag, term string) (songs, playtime int, err os.Error) {
	var a []Args

	if a, err = this.requestList("count \"%s\" \"%s\"", tag, term); err != nil || len(a) == 0 {
		return
	}

	songs = a[0].I("songs")
	playtime = a[0].I("playtime")
	return
}
