package hebcal

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// {
//   "title": "Hebcal São Paulo May 2015",
//   "date": "2022-05-30T17:38:40.629Z",
//   "location": {
//     "title": "São Paulo, Sao Paulo, Brazil",
//     "city": "São Paulo",
//     "tzid": "America/Sao_Paulo",
//     "latitude": -23.5475,
//     "longitude": -46.63611,
//     "cc": "BR",
//     "country": "Brazil",
//     "admin1": "Sao Paulo",
//     "asciiname": "Sao Paulo",
//     "geo": "geoname",
//     "geonameid": 3448439
//   },
//   "range": {
//     "start": "2015-05-22",
//     "end": "2015-05-25"
//   },
//   "items": [
//     {
//       "title": "Candle lighting: 5:11pm",
//       "date": "2015-05-22T17:11:00-03:00",
//       "category": "candles",
//       "title_orig": "Candle lighting",
//       "hebrew": "הדלקת נרות",
//       "memo": "Parashat Bamidbar"
//     },
//     {
//       "title": "Erev Shavuot",
//       "date": "2015-05-23",
//       "hdate": "5 Sivan 5775",
//       "category": "holiday",
//       "subcat": "major",
//       "hebrew": "ערב שבועות",
//       "link": "https://hebcal.com/h/shavuot-2015?us=js&um=api",
//       "memo": "Festival of Weeks. Commemorates the giving of the Torah at Mount Sinai"
//     },
//     {
//       "title": "Parashat Bamidbar",
//       "date": "2015-05-23",
//       "hdate": "5 Sivan 5775",
//       "category": "parashat",
//       "hebrew": "פרשת במדבר",
//       "leyning": {
//         "1": "Numbers 1:1-1:19",
//         "2": "Numbers 1:20-1:54",
//         "3": "Numbers 2:1-2:34",
//         "4": "Numbers 3:1-3:13",
//         "5": "Numbers 3:14-3:39",
//         "6": "Numbers 3:40-3:51",
//         "7": "Numbers 4:1-4:20",
//         "torah": "Numbers 1:1-4:20",
//         "haftarah": "Hosea 2:1-22",
//         "maftir": "Numbers 4:17-4:20",
//         "triennial": {
//           "1": "Numbers 2:1-2:9",
//           "2": "Numbers 2:10-2:16",
//           "3": "Numbers 2:17-2:24",
//           "4": "Numbers 2:25-2:31",
//           "5": "Numbers 2:32-2:34",
//           "6": "Numbers 3:1-3:4",
//           "7": "Numbers 3:5-3:13",
//           "maftir": "Numbers 3:11-3:13"
//         }
//       },
//       "link": "https://hebcal.com/s/bamidbar-20150523?us=js&um=api"
//     },
//     {
//       "title": "Candle lighting: 6:05pm",
//       "date": "2015-05-23T18:05:00-03:00",
//       "category": "candles",
//       "title_orig": "Candle lighting",
//       "hebrew": "הדלקת נרות",
//       "memo": "Erev Shavuot"
//     },
//     {
//       "title": "Shavuot I",
//       "date": "2015-05-24",
//       "hdate": "6 Sivan 5775",
//       "category": "holiday",
//       "subcat": "major",
//       "yomtov": true,
//       "hebrew": "שבועות א׳",
//       "leyning": {
//         "1": "Exodus 19:1-19:6",
//         "2": "Exodus 19:7-19:13",
//         "3": "Exodus 19:14-19:19",
//         "4": "Exodus 19:20-20:14",
//         "5": "Exodus 20:15-20:23",
//         "torah": "Exodus 19:1-20:23; Numbers 28:26-31",
//         "haftarah": "Ezekiel 1:1-28, 3:12",
//         "maftir": "Numbers 28:26-28:31"
//       },
//       "link": "https://hebcal.com/h/shavuot-2015?us=js&um=api",
//       "memo": "Festival of Weeks. Commemorates the giving of the Torah at Mount Sinai"
//     },
//     {
//       "title": "Candle lighting: 6:05pm",
//       "date": "2015-05-24T18:05:00-03:00",
//       "category": "candles",
//       "title_orig": "Candle lighting",
//       "hebrew": "הדלקת נרות",
//       "memo": "Shavuot I"
//     },
//     {
//       "title": "Shavuot II",
//       "date": "2015-05-25",
//       "hdate": "7 Sivan 5775",
//       "category": "holiday",
//       "subcat": "major",
//       "yomtov": true,
//       "hebrew": "שבועות ב׳",
//       "leyning": {
//         "1": "Deuteronomy 15:19-15:23",
//         "2": "Deuteronomy 16:1-16:3",
//         "3": "Deuteronomy 16:4-16:8",
//         "4": "Deuteronomy 16:9-16:12",
//         "5": "Deuteronomy 16:13-16:17",
//         "torah": "Deuteronomy 15:19-16:17; Numbers 28:26-31",
//         "haftarah": "Habakkuk 3:1-19",
//         "maftir": "Numbers 28:26-28:31"
//       },
//       "link": "https://hebcal.com/h/shavuot-2015?us=js&um=api",
//       "memo": "Festival of Weeks. Commemorates the giving of the Torah at Mount Sinai"
//     },
//     {
//       "title": "Havdalah: 6:05pm",
//       "date": "2015-05-25T18:05:00-03:00",
//       "category": "havdalah",
//       "title_orig": "Havdalah",
//       "hebrew": "הבדלה",
//       "memo": "Shavuot II"
//     }
//   ]
// }

type Response struct {
	Title    string `json:"title"`
	Date     string `json:"date"`
	Location struct {
		Geo string `json:"geo"`
	} `json:"location"`
	Items []Item `json:"items"`
}

type Item struct {
	Title    string `json:"title"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Subcat   string `json:"subcat"`
	Hebrew   string `json:"hebrew"`
	Memo     string `json:"memo"`
}

// Send the request to the Hebcal API
func SendHebcalRequest(date string) (Response, error) {
	client := resty.New()

	fmt.Println("date: ", date)

	// Send the request to Hebcal based on https://www.hebcal.com/hebcal?v=1&cfg=json&maj=on&min=on&mod=on&nx=on&year=now&month=x&ss=on&mf=on&c=on&geo=geoname&geonameid=3448439&M=on&s=on
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"v":     "1",
			"cfg":   "json",
			"s":     "on",
			"start": date,
			"end":   date,
		}).
		Get("https://www.hebcal.com/hebcal/")

	// return request url
	fmt.Println(response.Request.URL)

	var hebcalResponse Response
	err = json.Unmarshal(response.Body(), &hebcalResponse)
	if err != nil {
		return hebcalResponse, err
	}

	//  Check if the response is empty
	if len(hebcalResponse.Items) == 0 {
		return hebcalResponse, fmt.Errorf("Response is empty")
	}

	return hebcalResponse, nil
}

// Find item with category "parashat" and return the name
func GetParsha(response Response) (Item, error) {
	for _, item := range response.Items {
		if item.Category == "parashat" {
			return item, nil
		}
	}
	return Item{}, fmt.Errorf("Parsha not found")
}

func GetWeeklyPortion(date string) (Item, error) {
	//  Send request to Hebcal
	response, err := SendHebcalRequest(date)
	if err != nil {
		return Item{}, err
	}

	//  Get the name of the Parsha
	parsha, err := GetParsha(response)
	if err != nil {
		return Item{}, err
	}

	fmt.Println("Parsha: ", parsha.Title)
	fmt.Println("Hebrew: ", parsha.Hebrew)

	return parsha, nil
}
