package proxy

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

//go:embed templ.gotmpl
var tmpl string

func (s *Server) parseHTML(data []byte) (*MetaInfo, error) {
	h, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	i := MetaInfo{}
	doc := goquery.NewDocumentFromNode(h)
	var ok bool
	i.SiteName, ok = doc.Find("meta[property='og:site_name']").Attr("content")
	if !ok {
		return nil, errors.New("can't find site name")
	}
	i.IgermanURL, ok = doc.Find("meta[property='og:video']").Attr("content")
	if !ok {
		return nil, errors.New("can't find video url")
	}
	i.YoutubeURL, ok = doc.Find("meta[property='og:url']").Attr("content")
	if !ok {
		return nil, errors.New("can't find youtube video url")
	}
	alt := doc.Find("link[rel='alternate']")
	i.AlternateURL, ok = alt.Attr("href")
	if !ok {
		return nil, errors.New("can't find alternate url")
	}
	i.AlternateTitle, ok = alt.Attr("title")
	if !ok {
		return nil, errors.New("can't find alternate title")
	}
	i.YoutubeVideoID, err = s.idExtractor.GetVideoID(i.YoutubeURL)
	if err != nil {
		return nil, err
	}
	i.Description = i.SiteName
	i.Image = fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", i.YoutubeVideoID)
	return &i, nil
}

func buildOG(w io.Writer, info *MetaInfo) error {
	t, err := template.New("meta").Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(w, info)
}
