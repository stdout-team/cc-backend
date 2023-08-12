package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ResolveCoordinates(c *gin.Context) {
	type validator struct {
		Address string `binding:"required" form:"address"`
	}

	req := &validator{}

	if err := c.BindQuery(req); err != nil {
		c.Error(err)
		return
	}

	coordinates, err := resolveCoordinatesByStreetAddress(req.Address)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"lat": coordinates[0], "lon": coordinates[1]})
}

type errorGISDecoder struct {
	Meta struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	} `json:"meta"`
}

type resultGISDecoder struct {
	Result struct {
		Items []struct {
			Name  string `json:"name"`
			Point struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"point"`
		} `json:"items"`
	} `json:"result"`
}

func resolveCoordinatesByStreetAddress(address string) ([2]float64, error) {
	// https://catalog.api.2gis.com/3.0/items?q=%D1%8F%D0%BD%D0%B4%D0%B5%D0%BA%D1%81&sort_point=60.603395%2C56.841776&key=296c3a1a-574e-4239-953f-a763fe09e544&fields=items.point

	req := url.Values{}
	req.Set("q", address)
	req.Set("sort_point", "60.603395,56.841776")
	req.Set("key", os.Getenv("2GIS_TOKEN"))
	req.Set("fields", "items.point")

	resp, err := http.Get("https://catalog.api.2gis.com/3.0/items?" + req.Encode())
	if err != nil {
		return [2]float64{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return [2]float64{}, errors.New("error")
	}

	obj := &resultGISDecoder{}

	buf, err := io.ReadAll(resp.Body)
	str := string(buf)

	dec := json.NewDecoder(strings.NewReader(str))
	err = dec.Decode(obj)
	if err != nil {
		return [2]float64{}, err
	}

	for _, item := range obj.Result.Items {
		if item.Point.Lon != 0 || item.Point.Lat != 0 {
			return [2]float64{item.Point.Lat, item.Point.Lon}, nil
		}
	}

	return [2]float64{}, errors.New("not found")
}
