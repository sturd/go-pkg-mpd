// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

// Reports the current status of MPD, as well as the current settings of some
// playback options.
func (c *Client) Status() (s *Status, err error) {
	var a Args

	if a, err = c.request("status"); err != nil {
		return
	}

	s = readStatus(a)
	return
}

// Reports database and playlist statistics.
func (c *Client) Stats() (s *Stats, err error) {
	var a Args

	if a, err = c.request("stats"); err != nil {
		return
	}

	s = readStats(a)
	return
}

// Reports information about all known audio output devices.
func (c *Client) Outputs() (list []*Output, err error) {
	var a []Args

	if a, err = c.requestList("outputs"); err != nil {
		return
	}

	list = make([]*Output, 0, len(a))

	for _, m := range a {
		list = append(list, readOutput(m))
	}

	return
}

// Reports which commands the current user has access to.
func (c *Client) Commands() (v []string, err error) {
	var a []Args
	if a, err = c.requestList("commands"); err != nil {
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
func (c *Client) NotCommands() (v []string, err error) {
	var a []Args
	if a, err = c.requestList("notcommands"); err != nil {
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
func (c *Client) TagTypes() (v []string, err error) {
	var a []Args
	if a, err = c.requestList("tagtypes"); err != nil {
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
func (c *Client) UrlHandlers() (v []string, err error) {
	var a []Args
	if a, err = c.requestList("urlhandlers"); err != nil {
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
