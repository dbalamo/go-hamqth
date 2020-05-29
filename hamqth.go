package hamqth

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strconv"
	"strings"

	"github.com/dbalamo/go-hamqth/commons"
	"github.com/dbalamo/go-hamqth/datastructs"
)

type HamQth struct {
	User      string
	Password  string
	SessionId string
}

//Bands, to be used in Dx Cluster Spots Feed queries
const (
	Band160M = "160M"
	Band80M  = "80M"
	Band40M  = "40M"
	Band30M  = "30M"
	Band20M  = "20M"
	Band17M  = "17M"
	Band15M  = "15M"
	Band12M  = "12M"
	Band10M  = "10M"
	Band6M   = "6M"
	Band2M   = "2M"
	Band70CM = "70CM"
)

//Bands, to be used in Reverse Beacon Network queries
const (
	RbnBand160M = "160"
	RbnBand80M  = "80"
	RbnBand40M  = "40"
	RbnBand30M  = "30"
	RbnBand20M  = "20"
	RbnBand17M  = "17"
	RbnBand15M  = "15"
	RbnBand12M  = "12"
	RbnBand10M  = "10"
	RbnBand6M   = "6"
	RbnBand2M   = "2"
)

//Some modes, to be used in Reverse Beacon Network queries
const (
	RbnModeCW    = "CW"
	RbnModeRTTY  = "RTTY"
	RbnModePSK31 = "PSK31"
	RbnModePSK63 = "PSK63"
)

//Reverse Beacon Network query results ordering
const (
	RbnOrderFrequency = 1
	RbnOrderCallsign  = 2
	RbnOrderSpotAge   = 3
)

//Continent codes, used in Reverse Beacon Network query
const (
	RbnContinentEurope       = "EU"
	RbnContinentAfrica       = "AF"
	RbnContinentAsia         = "AS"
	RbnContinentNorthAmerica = "NA"
	RbnContinentSouthAmerica = "SA"
	RbnContinentOceania      = "OC"
)

func NewHamQth(user string, pwd string) *HamQth {
	return &HamQth{
		User:     user,
		Password: pwd,
	}
}

func (hu *HamQth) DxccSearchJson(callsign string) (*datastructs.DxccResponse, error) {
	var url = "https://www.hamqth.com/dxcc_json.php?callsign=" + callsign
	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/json"
	if bvret, err := commons.DoGet(url, headerParams); err == nil {
		resp := new(datastructs.DxccResponse)
		err = json.Unmarshal(bvret, resp)
		return resp, err
	} else {
		return nil, err
	}
}

func (hu *HamQth) DxccSearchXml(callsign string) (*datastructs.DxccResponse, error) {
	var url = "https://www.hamqth.com/dxcc.php?callsign=" + callsign
	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/xml"
	if bvret, err := commons.DoGet(url, headerParams); err == nil {
		resp := new(datastructs.DxccXmlResponse)
		err = xml.Unmarshal(bvret, &resp)
		if err != nil {
			commons.InnerPrintln("DxccSearchXml response unmarshal error - ", err.Error())
		}
		return &(resp.Dxcc), err
	} else {
		commons.InnerPrintln("DxccSearchXml response error -  ", err.Error())
		return nil, err
	}
}

func (hu *HamQth) getSessionId(renew bool) (string, error) {
	if !renew && hu.SessionId != "" {
		return hu.SessionId, nil
	} else {
		url := "https://www.hamqth.com/xml.php?u=" + hu.User + "&p=" + hu.Password
		headerParams := make(map[string]string)
		headerParams["Content-Type"] = "application/json"
		if resp, err := commons.DoGet(url, headerParams); err == nil {
			var xmlSessResp datastructs.XmlSessionResponse
			xml.Unmarshal(resp, &xmlSessResp)
			if xmlSessResp.Session.Error != "" {
				return "", errors.New(xmlSessResp.Session.Error)
			} else {
				hu.SessionId = xmlSessResp.Session.SessionId
				return xmlSessResp.Session.SessionId, nil
			}
		} else {
			commons.InnerPrintln("getSessionId Error : ", err.Error())
			return "", err
		}
	}
}

func (hu *HamQth) CallbookSearch(callsign string) (*datastructs.CallbookSearchResponse, error) {
	var sessId string
	var err error
	var resp *datastructs.CallbookSearchResponse
	if sessId, err = hu.getSessionId(false); err != nil {
		return nil, errors.New("Errors during getSessionId")
	}
	if resp, err = hu.callbookSearch(sessId, callsign); err != nil {
		if sessId, err = hu.getSessionId(true); err != nil {
			return nil, errors.New("Errors during getSessionId")
		}
		if resp, err = hu.callbookSearch(sessId, callsign); err != nil {
			return nil, errors.New("Errors during callbookSearch")
		}
	}
	return resp, err
}

func (hu *HamQth) callbookSearch(sessionId string, callsign string) (*datastructs.CallbookSearchResponse, error) {
	var url = "https://www.hamqth.com/xml.php"
	url = url + "?id=" + sessionId
	url = url + "&callsign=" + callsign
	url = url + "&prg=TEST"
	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/xml"
	if bvret, err := commons.DoGet(url, headerParams); err == nil {
		resp := new(datastructs.CallbookSearchResponse)
		err = xml.Unmarshal(bvret, &resp)
		if err != nil {
			commons.InnerPrintln("callbookSearch response unmarshal error - ", err.Error())
		}
		return resp, err
	} else {
		commons.InnerPrintln("callbookSearch response  error - ", err.Error())
		return nil, err
	}
}

func (hu *HamQth) RecentActivitySearch(callsign string, recActivity bool, logActivity bool, logBook bool) (*datastructs.RecentActivitySearchResponse, error) {
	var sessId string
	var err error
	var resp *datastructs.RecentActivitySearchResponse
	if sessId, err = hu.getSessionId(false); err != nil {
		return nil, errors.New("Errors during getSessionId")
	}
	if resp, err = hu.recentActivitySearch(sessId, callsign, recActivity, logActivity, logBook); err != nil {
		if sessId, err = hu.getSessionId(true); err != nil {
			return nil, errors.New("Errors during getSessionId")
		}
		if resp, err = hu.recentActivitySearch(sessId, callsign, recActivity, logActivity, logBook); err != nil {
			return nil, errors.New("Errors during RecentActivitySearch")
		}
	}
	return resp, err
}

func (hu *HamQth) recentActivitySearch(sessionId string, callsign string, recActivity bool, logActivity bool, logBook bool) (*datastructs.RecentActivitySearchResponse, error) {
	var url = "https://www.hamqth.com/xml_recactivity.php"
	url = url + "?id=" + sessionId
	url = url + "&callsign=" + callsign

	if recActivity {
		url = url + "&rec_activity=1"
	} else {
		url = url + "&rec_activity=0"
	}

	if logActivity {
		url = url + "&log_activity=1"
	} else {
		url = url + "&log_activity=0"
	}

	if logBook {
		url = url + "&logBook=1"
	} else {
		url = url + "&logBook=0"
	}

	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/xml"
	if bvret, err := commons.DoGet(url, headerParams); err == nil {
		resp := new(datastructs.RecentActivitySearchResponse)
		err = xml.Unmarshal(bvret, &resp)
		if err != nil {
			commons.InnerPrintln("recentActivitySearch response unmarshal error - ", err.Error())
		}
		return resp, err
	} else {
		commons.InnerPrintln("recentActivitySearch response error - ", err.Error())
		return nil, err
	}
}

func (hu *HamQth) LogUpload() {
	//TODO
}

func (hu *HamQth) RealtimeQsoUpload() {
	//TODO
}

func (hu *HamQth) DxClusterSpots(limit int, band string) ([]string, error) {
	url := "https://www.hamqth.com/dxc_csv.php"
	if limit > 0 && limit <= 200 {
		url = url + "?limit=" + strconv.Itoa(limit)
	} else {
		url = url + "?limit=60"
	}

	ucBand := strings.ToUpper(band)
	switch ucBand {
	case Band160M, Band80M, Band40M, Band20M, Band15M, Band12M, Band10M, Band6M, Band2M, Band70CM:
		url = url + "&band=" + ucBand
	}

	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/json"

	if bvret, err := commons.DoGet(url, headerParams); err == nil {
		sv := strings.Split(string(bvret), "\n")
		return sv[:len(sv)-1], nil
	} else {
		commons.InnerPrintln("DxClusterSpots response error - ", err.Error())
		sv := make([]string, 0)
		return sv, err
	}
}

func (hu *HamQth) AwardVerification() {
	//TODO
}

func (hu *HamQth) AgregatedRBN(bands []string, fromContinents []string, mode []string,
	spottedContinents []string, waz []string, itu []string, ageSeconds int, order int) (*datastructs.ReverseBeaconResponse, error) {

	url := "https://www.hamqth.com/rbn_data.php?data=1"
	url = url + addArrayQueryStringParam("band", bands)
	url = url + addArrayQueryStringParam("fromcont", fromContinents)
	url = url + addArrayQueryStringParam("mode", mode)
	url = url + addArrayQueryStringParam("cont", spottedContinents)
	url = url + addArrayQueryStringParam("waz", waz)
	url = url + addArrayQueryStringParam("itu", itu)
	url = url + "&age=" + strconv.Itoa(ageSeconds)
	url = url + "&order=" + strconv.Itoa(order)

	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/json"
	if bvret, err := commons.DoGet(url, headerParams); err == nil {
		resp := new(datastructs.ReverseBeaconResponse)
		kvMap := make(map[string]map[string]interface{})
		err = json.Unmarshal(bvret, &kvMap)

		for _, v := range kvMap {
			bv, _ := json.Marshal(v)
			rbs := datastructs.ReverseBeaconSpot{}
			json.Unmarshal(bv, &rbs)
			resp.RbnSpots = append(resp.RbnSpots, rbs)
		}

		return resp, err
	} else {
		commons.InnerPrintln("AgregatedRBN response error - ", err.Error())
		return nil, err
	}
}

func addArrayQueryStringParam(paramName string, values []string) string {
	ret := ""
	if len(values) > 0 {
		ret = ret + "&" + paramName + "="
		vals := ""
		for _, currVal := range values {
			vals = vals + currVal + ","
		}
		ret = ret + vals[:len(vals)-1]
	}
	return ret
}
