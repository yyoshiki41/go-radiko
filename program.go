package radiko

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/yyoshiki41/go-radiko/internal"
)

// Stations is a slice of Station.
type Stations []Station

// Station is a struct.
type Station struct {
	ID    string `xml:"id,attr"`
	Name  string `xml:"name"`
	Scd   Scd    `xml:"scd,omitempty"`
	Progs Progs  `xml:"progs,omitempty"`
}

// Scd is a struct.
type Scd struct {
	Progs Progs `xml:"progs"`
}

// Progs is a slice of Prog.
type Progs struct {
	Date  string `xml:"date"`
	Progs []Prog `xml:"prog"`
}

// Prog is a struct.
type Prog struct {
	Ft       string `xml:"ft,attr"`
	To       string `xml:"to,attr"`
	Ftl      string `xml:"ftl,attr"`
	Tol      string `xml:"tol,attr"`
	Dur      string `xml:"dur,attr"`
	Title    string `xml:"title"`
	SubTitle string `xml:"sub_title"`
	Desc     string `xml:"desc"`
	Pfm      string `xml:"pfm"`
	Info     string `xml:"info"`
	URL      string `xml:"url"`
}

// GetStationsByAreaID returns program's meta-info.
func (c *Client) GetStationsByAreaID(ctx context.Context, areaID string, date time.Time) (Stations, error) {
	apiEndpoint := path.Join(apiV3,
		"program/date", internal.ProgramsDate(date),
		fmt.Sprintf("%s.xml", areaID))

	req, err := c.newRequest(ctx, "GET", apiEndpoint, &Params{})
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var entity stationsEntity
	if err = xml.Unmarshal(b, &entity); err != nil {
		return nil, err
	}
	return entity.stations(), nil
}

// GetStations returns program's meta-info in the current location.
// This API wraps GetStationsByAreaID.
func (c *Client) GetStations(ctx context.Context, date time.Time) (Stations, error) {
	areaID, err := AreaID()
	if err != nil {
		return nil, err
	}

	return c.GetStationsByAreaID(ctx, areaID, date)
}

// GetNowProgramsByAreaID returns program's meta-info which are currently on the air.
func (c *Client) GetNowProgramsByAreaID(ctx context.Context, areaID string) (Stations, error) {
	apiEndpoint := apiPath(apiV2, "program/now")

	req, err := c.newRequest(ctx, "GET", apiEndpoint, &Params{
		query: map[string]string{
			"area_id": areaID,
		},
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var entity stationsEntity
	if err = xml.Unmarshal(b, &entity); err != nil {
		return nil, err
	}

	return entity.stations(), nil
}

// GetNowPrograms returns program's meta-info in the current location.
// This API wraps GetNowProgramsByAreaID.
func (c *Client) GetNowPrograms(ctx context.Context) (Stations, error) {
	areaID, err := AreaID()
	if err != nil {
		return nil, err
	}

	return c.GetNowProgramsByAreaID(ctx, areaID)
}

// GetProgramByStartTime returns a specified program.
// This API wraps GetStations.
func (c *Client) GetProgramByStartTime(ctx context.Context, stationID string, start time.Time) (*Prog, error) {
	if stationID == "" {
		return nil, errors.New("StationID is empty.")
	}

	stations, err := c.GetStations(ctx, start)
	if err != nil {
		return nil, err
	}

	ft := internal.Datetime(start)
	var prog *Prog
	for _, s := range stations {
		if s.ID == stationID {
			for _, p := range s.Progs.Progs {
				if p.Ft == ft {
					prog = &p
					break
				}
			}
		}
	}
	if prog == nil {
		return nil, errors.New("Can not find a program.")
	}
	return prog, nil
}

// stationsEntity includes a response struct for client's users.
type stationsEntity struct {
	XMLName     xml.Name `xml:"radiko"`
	XMLStations struct {
		XMLName  xml.Name `xml:"stations"`
		Stations Stations `xml:"station"`
	} `xml:"stations"`
}

// stations returns Stations which is a response struct for client's users.
func (e *stationsEntity) stations() Stations {
	return e.XMLStations.Stations
}
