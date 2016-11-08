package radiko

import (
	"context"
	"encoding/xml"
	"io/ioutil"
)

type Programs struct {
	Stations Stations `xml:"stations"`
}

type Stations struct {
	Stations []Station `xml:"station"`
}

type Station struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name"`
	Scd  Scd    `xml:"scd"`
}

type Scd struct {
	Progs Progs `xml:"progs"`
}

type Progs struct {
	Date  string `xml:"date"`
	Progs []Prog `xml:"prog"`
}

type Prog struct {
	Ft       string `xml:"ft,attr"`
	To       string `xml:"to,attr"`
	Title    string `xml:"title"`
	SubTitle string `xml:"sub_title"`
	Desc     string `xml:"desc"`
	Info     string `xml:"info"`
	URL      string `xml:"url"`
}

// GetNowPrograms returns program's meta-info which are currently on the air.
func (c *Client) GetNowPrograms(ctx context.Context, areaID string) (*Programs, error) {
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
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var programs Programs
	if err = xml.Unmarshal(b, &programs); err != nil {
		return nil, err
	}
	return &programs, err
}

// GetStationMaps returns available station's map.
// This API wraps GetNowPrograms.
func (c *Client) GetStationMaps(ctx context.Context, areaID string) (map[string]string, error) {
	m := make(map[string]string)

	programs, err := c.GetNowPrograms(ctx, areaID)
	if err != nil {
		return m, err
	}

	for _, s := range programs.Stations.Stations {
		m[s.ID] = s.Name
	}
	return m, err
}
