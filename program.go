package radiko

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path"
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

// GetStations returns program's meta-info.
func (c *Client) GetStations(ctx context.Context, areaID, date string) (*Stations, error) {
	apiEndpoint := path.Join(apiV3,
		"program/date", date,
		fmt.Sprintf("%s.xml", areaID))

	req, err := c.newRequest("GET", apiEndpoint, &Params{})
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	resp, err := c.HTTPClient.Do(req)
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
	return entity.stations(), err
}

// GetNowPrograms returns program's meta-info which are currently on the air.
func (c *Client) GetNowPrograms(ctx context.Context, areaID string) (*Stations, error) {
	apiEndpoint := apiPath(apiV2, "program/now")

	req, err := c.newRequest("GET", apiEndpoint, &Params{
		query: map[string]string{
			"area_id": areaID,
		},
	})
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	resp, err := c.HTTPClient.Do(req)
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

	return entity.stations(), err
}

// stationsEntity includes a response struct for client's users.
type stationsEntity struct {
	XMLName     xml.Name `xml:"radiko"`
	XMLStations struct {
		XMLName  xml.Name  `xml:"stations"`
		Stations *Stations `xml:"station"`
	} `xml:"stations"`
}

// stations return Stations which is a response struct for client's users.
func (e *stationsEntity) stations() *Stations {
	return e.XMLStations.Stations
}
