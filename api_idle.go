package mpd

type SubSystem uint8

const (
	DatabaseSys SubSystem = iota
	UpdateSys
	StoredPlaylistSys
	PlaylistSys
	PlayerSys
	MixerSys
	OutputSys
	StickerSys
	SubscriptionSys
	MessageSys
)

func readSubSystem(a Args) SubSystem {
	var s SubSystem

	switch a.S("changed") {
	case "database":
		s = DatabaseSys
	case "update":
		s = UpdateSys
	case "stored_playlist":
		s = StoredPlaylistSys
	case "playlist":
		s = PlaylistSys
	case "player":
		s = PlayerSys
	case "mixer":
		s = MixerSys
	case "output":
		s = OutputSys
	case "sticker":
		s = StickerSys
	case "subscription":
		s = SubscriptionSys
	case "message":
		s = MessageSys
	}
	return s
}

func (c *Client) Idle() (s SubSystem, err error) {
	var a Args

	if a, err = c.request("idle"); err != nil {
		return
	}

	s = readSubSystem(a)
	return
}

func (c *Client) IdleSubSystem(subsystem string) (changed bool, err error) {
	var a Args

	if a, err = c.request("idle %s", subsystem); err != nil {
		return
	}

	changed = a != nil
	return
}
