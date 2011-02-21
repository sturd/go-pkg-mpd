// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import "os"

// Reports the current status of MPD, as well as the current settings of some
// playback options.
func (this *Client) Status() (s *Status, err os.Error) {
	var a Args

	if a, err = this.request("status"); err != nil {
		return
	}

	s = readStatus(a)
	return
}

// Reports database and playlist statistics.
func (this *Client) Stats() (s *Stats, err os.Error) {
	var a Args

	if a, err = this.request("stats"); err != nil {
		return
	}

	s = readStats(a)
	return
}

// Reports information about all known audio output devices.
func (this *Client) Outputs() (list []*Output, err os.Error) {
	var a []Args

	if a, err = this.requestList("outputs"); err != nil {
		return
	}

	list = make([]*Output, 0, len(a))

	for _, m := range a {
		list = append(list, readOutput(m))
	}

	return
}

// Reports which commands the current user has access to.
func (this *Client) Commands() (v []string, err os.Error) {
	var a []Args
	if a, err = this.requestList("commands"); err != nil {
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

// Reports which commands the current user has *no* access to.
func (this *Client) NotCommands() (v []string, err os.Error) {
	var a []Args
	if a, err = this.requestList("notcommands"); err != nil {
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

// Reports a list of available song metadata fields.
func (this *Client) TagTypes() (v []string, err os.Error) {
	var a []Args
	if a, err = this.requestList("tagtypes"); err != nil {
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

// Reports a list of available URL handlers.
func (this *Client) UrlHandlers() (v []string, err os.Error) {
	var a []Args
	if a, err = this.requestList("urlhandlers"); err != nil {
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
