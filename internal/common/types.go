package common

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/dreadl0ck/tlsx"
	"github.com/gaukas/clienthellod"
	"github.com/wi1dcard/fingerproxy/pkg/ja4"
	"github.com/wi1dcard/fingerproxy/pkg/metadata"
)

type CollectedInfoFingerProxy struct {
	ClientHello []byte `json:"clienthello"`
	JA3         string `json:"ja3"`
	JA4         string `json:"ja4"`
	HTTP2       string `json:"http2"`
	Detail      struct {
		JA3      *tlsx.ClientHelloBasic             `json:"ja3"`
		JA4      *ja4.JA4Fingerprint                `json:"ja4"`
		HTTP2    metadata.HTTP2FingerprintingFrames `json:"http2"`
		MetaData *metadata.Metadata                 `json:"metadata"`
	} `json:"detail"`
}

type CollectedInfoClienthellod struct {
	Raw           []byte                            `json:"raw"`
	TLS           *clienthellod.ClientHello         `json:"tls,omitempty"`
	QUIC          *clienthellod.ClientInitialPacket `json:"quic,omitempty"`
	NID           *int64                            `json:"nid,omitempty"`
	NIDNormalized *int64                            `json:"nid_normalized,omitempty"`
	ID            string                            `json:"id,omitempty"`
	IDNormalized  string                            `json:"id_normalized,omitempty"`
}

type CollectedInfo struct {
	BrowserType BrowserType          `json:"browser"`
	Date        time.Time            `json:"date"`
	URL         string               `json:"url"`
	UserAgent   string               `json:"user-agent"`
	Headers     http.Header          `json:"headers,omitempty"`
	Proto       string               `json:"proto"`
	TLS         *tls.ConnectionState `json:"tls,omitempty"`

	ID string `json:"id"`

	FingerProxy  *CollectedInfoFingerProxy  `json:"fingerproxy,omitempty"`
	Clienthellod *CollectedInfoClienthellod `json:"clienthellod,omitempty"`
}

type BrowserType string

const (
	BrowserTypeChrome  BrowserType = "chrome"
	BrowserTypeFirefox BrowserType = "firefox"
	BrowserTypeEdge    BrowserType = "edge"
)

func Ptr[T any](v T) *T {
	return &v
}
