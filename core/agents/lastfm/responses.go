package lastfm

type Response struct {
	Artist         Artist         `json:"artist"`
	SimilarArtists SimilarArtists `json:"similarartists"`
	TopTracks      TopTracks      `json:"toptracks"`
	Error          int            `json:"error"`
	Message        string         `json:"message"`
	Token          string         `json:"token"`
	Session        Session        `json:"session"`
	NowPlaying     NowPlaying     `json:"nowplaying"`
	Scrobbles      Scrobbles      `json:"scrobbles"`
}

type Artist struct {
	Name       string        `json:"name"`
	MBID       string        `json:"mbid"`
	URL        string        `json:"url"`
	Image      []ArtistImage `json:"image"`
	Streamable string        `json:"streamable"`
	Stats      struct {
		Listeners string `json:"listeners"`
		Plays     string `json:"plays"`
	} `json:"stats"`
	Similar SimilarArtists `json:"similar"`
	Tags    struct {
		Tag []ArtistTag `json:"tag"`
	} `json:"tags"`
	Bio ArtistBio `json:"bio"`
}

type SimilarArtists struct {
	Artists []Artist `json:"artist"`
	Attr    Attr     `json:"@attr"`
}

type Attr struct {
	Artist string `json:"artist"`
}

type ArtistImage struct {
	URL  string `json:"#text"`
	Size string `json:"size"`
}

type ArtistTag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ArtistBio struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

type Track struct {
	Name string `json:"name"`
	MBID string `json:"mbid"`
}

type TopTracks struct {
	Track []Track `json:"track"`
	Attr  Attr    `json:"@attr"`
}

type Session struct {
	Name       string `json:"name"`
	Key        string `json:"key"`
	Subscriber int    `json:"subscriber"`
}

type NowPlaying struct {
	Artist struct {
		Corrected string `json:"corrected"`
		Text      string `json:"#text"`
	} `json:"artist"`
	IgnoredMessage struct {
		Code string `json:"code"`
		Text string `json:"#text"`
	} `json:"ignoredMessage"`
	Album struct {
		Corrected string `json:"corrected"`
		Text      string `json:"#text"`
	} `json:"album"`
	AlbumArtist struct {
		Corrected string `json:"corrected"`
		Text      string `json:"#text"`
	} `json:"albumArtist"`
	Track struct {
		Corrected string `json:"corrected"`
		Text      string `json:"#text"`
	} `json:"track"`
}

type Scrobbles struct {
	Attr struct {
		Accepted int `json:"accepted"`
		Ignored  int `json:"ignored"`
	} `json:"@attr"`
	Scrobble struct {
		Artist struct {
			Corrected string `json:"corrected"`
			Text      string `json:"#text"`
		} `json:"artist"`
		IgnoredMessage struct {
			Code string `json:"code"`
			Text string `json:"#text"`
		} `json:"ignoredMessage"`
		AlbumArtist struct {
			Corrected string `json:"corrected"`
			Text      string `json:"#text"`
		} `json:"albumArtist"`
		Timestamp string `json:"timestamp"`
		Album     struct {
			Corrected string `json:"corrected"`
			Text      string `json:"#text"`
		} `json:"album"`
		Track struct {
			Corrected string `json:"corrected"`
			Text      string `json:"#text"`
		} `json:"track"`
	} `json:"scrobble"`
}
