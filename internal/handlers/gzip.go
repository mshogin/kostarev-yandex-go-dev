package handlers

import (
	"compress/gzip"
	"io"
	"net/http"
)

type CompressWriter struct {
	w  http.ResponseWriter
	gw *gzip.Writer
}

type CompressReader struct {
	r  io.ReadCloser
	gr *gzip.Reader
}

func NewCompressWriter(w http.ResponseWriter) *CompressWriter {
	gw := gzip.NewWriter(w)
	return &CompressWriter{w, gw}
}

func (cw *CompressWriter) Header() http.Header {
	return cw.w.Header()
}

func (cw *CompressWriter) Write(b []byte) (int, error) {
	return cw.gw.Write(b)
}

func (cw *CompressWriter) WriteHeader(statusCode int) {
	cw.w.WriteHeader(statusCode)
	if statusCode >= 200 && statusCode < 300 {
		cw.w.Header().Set("Content-Encoding", "gzip")
	}
}

func (cw *CompressWriter) Close() error {
	return cw.gw.Close()
}

func NewCompressReader(closer io.ReadCloser) (*CompressReader, error) {
	g, err := gzip.NewReader(closer)
	if err != nil {
		return nil, err
	}

	return &CompressReader{closer, g}, nil
}

func (cr *CompressReader) Read(b []byte) (int, error) {
	return cr.gr.Read(b)
}

func (cr *CompressReader) Close() error {
	if err := cr.r.Close(); err != nil {
		return err
	}

	return cr.r.Close()
}
