package m3u8

/*
 Part of M3U8 parser & generator library.
 This file defines functions related to playlist generation.

 Copyright 2013-2017 The Project Developers.
 See the AUTHORS and LICENSE files at the top-level directory of this distribution
 and at https://github.com/grafov/m3u8/

 ॐ तारे तुत्तारे तुरे स्व
*/

import (
	"errors"
	"math"
	"time"
)

var (
	ErrPlaylistFull = errors.New("playlist is full")
)

// Set version of the playlist accordingly with section 7
func version(ver *uint8, newver uint8) {
	if *ver < newver {
		*ver = newver
	}
}

// Create new empty master playlist.
// Master playlist consists of variants.
func NewMasterPlaylist() *MasterPlaylist {
	p := new(MasterPlaylist)
	p.ver = minver
	return p
}

// Append variant to master playlist.
// This operation does reset playlist cache.
func (p *MasterPlaylist) Append(uri string, chunklist *MediaPlaylist, params VariantParams) {
	v := new(Variant)
	v.URI = uri
	v.Chunklist = chunklist
	v.VariantParams = params
	p.Variants = append(p.Variants, v)
	if len(v.Alternatives) > 0 {
		// From section 7:
		// The EXT-X-MEDIA tag and the AUDIO, VIDEO and SUBTITLES attributes of
		// the EXT-X-STREAM-INF tag are backward compatible to protocol version
		// 1, but playback on older clients may not be desirable.  A server MAY
		// consider indicating a EXT-X-VERSION of 4 or higher in the Master
		// Playlist but is not required to do so.
		version(&p.ver, 4) // so it is optional and in theory may be set to ver.1
		// but more tests required
	}
	p.buf.Reset()
}

// Creates new media playlist structure.
// Winsize defines how much items will displayed on playlist generation.
// Capacity is total size of a playlist.
func NewMediaPlaylist(winsize uint, capacity uint) (*MediaPlaylist, error) {
	p := new(MediaPlaylist)
	p.ver = minver
	p.capacity = capacity
	if err := p.SetWinSize(winsize); err != nil {
		return nil, err
	}
	p.Segments = make([]*MediaSegment, capacity)
	return p, nil
}

// last returns the previously written segment's index
func (p *MediaPlaylist) last() uint {
	if p.tail == 0 {
		return p.capacity - 1
	}
	return p.tail - 1
}

// Append general chunk to the tail of chunk slice for a media playlist.
// This operation does reset playlist cache.
func (p *MediaPlaylist) Append(uri string, duration float64, title string) error {
	seg := new(MediaSegment)
	seg.URI = uri
	seg.Duration = duration
	seg.Title = title
	return p.AppendSegment(seg)
}

// AppendSegment appends a MediaSegment to the tail of chunk slice for a media playlist.
// This operation does reset playlist cache.
func (p *MediaPlaylist) AppendSegment(seg *MediaSegment) error {
	if p.head == p.tail && p.count > 0 {
		return ErrPlaylistFull
	}
	p.Segments[p.tail] = seg
	p.tail = (p.tail + 1) % p.capacity
	p.count++
	if p.TargetDuration < seg.Duration {
		p.TargetDuration = math.Ceil(seg.Duration)
	}
	p.buf.Reset()
	return nil
}

// Count tells us the number of items that are currently in the media playlist
func (p *MediaPlaylist) Count() uint {
	return p.count
}

// Set limit and offset for the current media segment (EXT-X-BYTERANGE support for protocol version 4).
func (p *MediaPlaylist) SetRange(limit, offset int64) error {
	if p.count == 0 {
		return errors.New("playlist is empty")
	}
	version(&p.ver, 4) // due section 3.4.1
	p.Segments[p.last()].Limit = limit
	p.Segments[p.last()].Offset = offset
	return nil
}

// SetSCTE sets the SCTE cue format for the current media segment.
//
// Deprecated: Use SetSCTE35 instead.
func (p *MediaPlaylist) SetSCTE(cue string, id string, time float64) error {
	return p.SetSCTE35(&SCTE{Syntax: SCTE35_67_2014, Cue: cue, ID: id, Time: time})
}

// SetSCTE35 sets the SCTE cue format for the current media segment
func (p *MediaPlaylist) SetSCTE35(scte35 *SCTE) error {
	if p.count == 0 {
		return errors.New("playlist is empty")
	}
	p.Segments[p.last()].SCTE = scte35
	return nil
}

// Set discontinuity flag for the current media segment.
// EXT-X-DISCONTINUITY indicates an encoding discontinuity between the media segment
// that follows it and the one that preceded it (i.e. file format, number and type of tracks,
// encoding parameters, encoding sequence, timestamp sequence).
func (p *MediaPlaylist) SetDiscontinuity() error {
	if p.count == 0 {
		return errors.New("playlist is empty")
	}
	p.Segments[p.last()].Discontinuity = true
	return nil
}

// Set program date and time for the current media segment.
// EXT-X-PROGRAM-DATE-TIME tag associates the first sample of a
// media segment with an absolute date and/or time.  It applies only
// to the current media segment.
// Date/time format is YYYY-MM-DDThh:mm:ssZ (ISO8601) and includes time zone.
func (p *MediaPlaylist) SetProgramDateTime(value time.Time) error {
	if p.count == 0 {
		return errors.New("playlist is empty")
	}
	p.Segments[p.last()].ProgramDateTime = value
	return nil
}

// SetWinSize overwrites the playlist's window size.
func (p *MediaPlaylist) SetWinSize(winsize uint) error {
	if winsize > p.capacity {
		return errors.New("capacity must be greater than winsize or equal")
	}
	p.winsize = winsize
	return nil
}
