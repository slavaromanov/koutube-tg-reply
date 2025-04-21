package proxy

import (
	"io"
	"net/http"
	"net/url"
)

type VideoIDExtractor interface {
	GetVideoID(s string) (string, error)
}

type ProxyPort string

type Server struct {
	port        string
	idExtractor VideoIDExtractor
	*http.ServeMux
}

func NewServer(port ProxyPort, extractor VideoIDExtractor) *Server {
	return &Server{
		idExtractor: extractor,
		port:        string(port),
		ServeMux:    http.NewServeMux(),
	}
}

func (s *Server) Run() error {
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := url.URL{
			Host:     "koutu.be",
			Scheme:   "https",
			Path:     r.URL.Path,
			RawQuery: r.URL.RawQuery,
		}
		resp, err := http.Get(u.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		i, err := s.parseHTML(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		iresp, err := http.Get(i.IgermanURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		truVideoURL, err := iresp.Location()
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		i.IgermanURL = truVideoURL.String()
		if err := buildOG(w, i); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	return http.ListenAndServe(":"+s.port, s)
}
