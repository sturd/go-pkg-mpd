package mpd

type SubSystem uint8

const (
	DatabaseSystem SubSystem = iota
	UpdateSystem
	StoredPlaylistSystem
	PlaylistSystem
	PlayerSystem
	MixerSystem
	OutputSystem
	StickerSystem
	SubscriptionSystem
	MessageSystem
)

func readSubSystem(a Args) SubSystem {
	var s SubSystem

	switch a.S("changed") {
	case "database":
		s = DatabaseSystem
	case "update":
		s = UpdateSystem
	case "stored_playlist":
		s = StoredPlaylistSystem
	case "playlist":
		s = PlaylistSystem
	case "player":
		s = PlayerSystem
	case "mixer":
		s = MixerSystem
	case "output":
		s = OutputSystem
	case "sticker":
		s = StickerSystem
	case "subscription":
		s = SubscriptionSystem
	case "message":
		s = MessageSystem
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

func (c *Client) IdleSubSystem(subsystem SubSystem) (changed bool, err error) {
	var a Args
	var s string

	switch subsystem {
	case DatabaseSystem:
		s = "database"
	case UpdateSystem:
		s = "update"
	case StoredPlaylistSystem:
		s = "stored_playlist"
	case PlaylistSystem:
		s = "playlist"
	case PlayerSystem:
		s = "player"
	case MixerSystem:
		s = "mixer"
	case OutputSystem:
		s = "output"
	case StickerSystem:
		s = "sticker"
	case SubscriptionSystem:
		s = "subscription"
	case MessageSystem:
		s = "message"
	}

	if a, err = c.request("idle %s", s); err != nil {
		return
	}

	changed = a != nil
	return
}
