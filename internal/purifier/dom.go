package purifier

import "github.com/microcosm-cc/bluemonday"

func NewPurifierPolicy() *bluemonday.Policy {
	p := bluemonday.UGCPolicy()

	p.AllowDataAttributes()
	p.AllowAttrs("width").OnElements("iframe")
	p.AllowAttrs("height").OnElements("iframe")
	p.AllowAttrs("allowfullscreen").OnElements("iframe")
	p.AllowAttrs("autoplay").OnElements("iframe")
	p.AllowAttrs("disablekbcontrols").OnElements("iframe")
	p.AllowAttrs("enableiframeapi").OnElements("iframe")
	p.AllowAttrs("endtime").OnElements("iframe")
	p.AllowAttrs("ivloadpolicy").OnElements("iframe")
	p.AllowAttrs("loop").OnElements("iframe")
	p.AllowAttrs("modestbranding").OnElements("iframe")
	p.AllowAttrs("origin").OnElements("iframe")
	p.AllowAttrs("playlist").OnElements("iframe")
	p.AllowAttrs("src").OnElements("iframe")
	p.AllowAttrs("start").OnElements("iframe")

	return p
}

func PurifyHTML(howTo *string, purifier *bluemonday.Policy) *string {
	if howTo == nil {
		return nil
	}

	html := purifier.Sanitize(*howTo)
	return &html
}
