package radiko

import (
	"compress/zlib"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	playerURL = "http://radiko.jp/apps/js/flash/myplayer-release.swf"
)

func downloadBinary() ([]byte, error) {
	resp, err := http.Get(playerURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return swfExtract(resp.Body)
}

const targetID = 12   // swfextract -b "12"
const targetCode = 87 // swfextract "-b" 12

const headerCWS = 8
const headerRect = 5
const rectNum = 4
const headerRest = 2 + 2
const binaryOffset = 6

func swfExtract(body io.Reader) ([]byte, error) {
	io.CopyN(ioutil.Discard, body, headerCWS)
	zf, err := zlib.NewReader(body)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(zf)
	if err != nil {
		return nil, err
	}

	offset := 0

	// Skip Rect
	rectSize := int(buf[offset] >> 3)
	rectOffset := (headerRect + rectNum*rectSize + 7) / 8

	offset += rectOffset

	// Skip the rest header
	offset += headerRest

	// Read tags
	for i := 0; ; i++ {
		// tag code
		code := int(buf[offset+1])<<2 + int(buf[offset])>>6

		// tag length
		len := int(buf[offset] & 0x3f)

		// Skip tag header
		offset += 2

		// tag length (if long version)
		if len == 0x3f {
			len = int(buf[offset])
			len += int(buf[offset+1]) << 8
			len += int(buf[offset+2]) << 16
			len += int(buf[offset+3]) << 24

			// skip tag lentgh header
			offset += 4
		}

		// Not found...
		if code == 0 {
			return nil, errors.New("swf extract failed")
		}
		// tag ID
		id := int(buf[offset]) + int(buf[offset+1])<<8

		// Found?
		if code == targetCode && id == targetID {
			return buf[offset+binaryOffset : offset+len], nil
		}

		offset += len
	}
}
