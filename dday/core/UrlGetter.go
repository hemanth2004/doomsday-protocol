package core

type UrlGetter struct {
	Key           string // used to identify resouce in the server
	UpdatedURLURL []string
	DefaultURLs   []string
}

func (u UrlGetter) GetUrl() string {
	if len(u.UpdatedURLURL) == 0 {
		return u.DefaultURLs[0]
	} else {
		return u.UpdatedURLURL[0]
	}
}
