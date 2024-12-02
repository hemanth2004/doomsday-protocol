package core

type UrlGetter struct {
	Key           string // used to identify resouce in the server
	RecentURLUsed string
	UpdatedURLURL []string
	DefaultURLs   []string
}

func (u *UrlGetter) GetUrl() string {
	if len(u.UpdatedURLURL) == 0 {
		u.RecentURLUsed = u.DefaultURLs[0]

	} else {
		u.RecentURLUsed = u.UpdatedURLURL[0]
	}

	return u.RecentURLUsed
}
