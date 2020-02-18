package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

)

type GeoData struct {
	Response struct {
		MetaInfo struct {
			TimeStamp string `json:"TimeStamp"`
		} `json:MetaInfo`
		View []struct {
			Result []struct {
				MatchLevel string `json:"MatchLevel"`
				Location   struct {
					Address struct {
						Label       string `json:"Label"`
						Country     string `json:"Country"`
						State       string `json:"State"`
						County      string `json:"County"`
						City        string `json:"City"`
						District    string `json:"District"`
						Street      string `json:"Street"`
						HouseNumber string `json:"HouseNumber"`
						PostalCode  string `json:"PostalCode"`
					} `json:"Address"`
				} `json:"Location"`
			} `json:"Result"`
		} `json:"View"`
	} `json:"Response"`
}

type Position struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Geocoder struct {
	AppId   string `json:"add_id"`
	AppCode string `json:"add_code"`
}

func (geocoder *Geocoder) reverse(position Position) (GeoData, error) {
	endpoint, _ := url.Parse("https://reverse.geocoder.api.here.com/6.2/reversegeocode.json")
	queryParams := endpoint.Query()
	queryParams.Set("app_id", geocoder.AppId)
	queryParams.Set("app_code", geocoder.AppCode)
	queryParams.Set("mode", "retrieveAddresses")
	queryParams.Set("prox", position.Latitude+","+position.Longitude)
	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())
	if err != nil {
		return GeoData{}, err
	}

	data, _ := ioutil.ReadAll((response.Body))
	var geoData GeoData
	json.Unmarshal(data, &geoData)
	return geoData, nil

}

func main() {
	latitude := flag.String("lat", "37.7397", "Latitude")
	longitude := flag.String("lng", "-12.4252", "Longitude")
	flag.Parse()
	geocoder := Geocoder{AppId: "APP-ID-HERE", AppCode: "APP-CODE-HERE"}
	result, err := geocoder.reverse((Position{Latitude: *latitude, Longitude: *longitude}))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}
	if len(result.Response.View) > 0 && len(result.Response.View[0].Result) > 0 {
		data, _ := json.Marshal(result.Response.View[0].Result[0])
		fmt.Println(string(data))
	}
}
