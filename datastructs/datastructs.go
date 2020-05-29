package datastructs

import "encoding/xml"

type DxccResponse struct {
	Callsign  string `json:"callsign" xml:"callsign"`
	Name      string `json:"name" xml:"name"`
	Details   string `json:"details" xml:"details"`
	Continent string `json:"continent" xml:"continent"`
	Utc       string `json:"utc" xml:"utc"`
	Waz       string `json:"waz" xml:"waz"`
	Itu       string `json:"itu" xml:"itu"`
	Lat       string `json:"lat" xml:"lat"`
	Lng       string `json:"lng" xml:"lng"`
	Adif      string `json:"adif" xml:"adif"`
}

type DxccXmlResponse struct {
	HamQTH xml.Name     `xml:"HamQTH"`
	Dxcc   DxccResponse `xml:"dxcc"`
}

type XmlSessionContent struct {
	SessionId string `xml:"session_id"`
	Error     string `xml:"error"`
}

type XmlSessionResponse struct {
	HamQTH  xml.Name          `xml:"HamQTH"`
	Session XmlSessionContent `xml:"session"`
}

type CallbookSearchResponse struct {
	HamQTH xml.Name              `xml:"HamQTH"`
	Search CallbookSearchContent `xml:"search"`
}

type CallbookSearchContent struct {
	Callsign   string `xml:"callsign"`
	Nick       string `xml:"nick"`
	Qth        string `xml:"qth"`
	Country    string `xml:"country"`
	Adif       string `xml:"adif"`
	Itu        string `xml:"itu"`
	Cq         string `xml:"cq"`
	Grid       string `xml:"grid"`
	AdrName    string `xml:"adr_name"`
	AdrStreet1 string `xml:"adr_street1"`
	AdrStreet2 string `xml:"adr_street2"`
	AdrStreet3 string `xml:"adr_street3"`
	AdrCity    string `xml:"adr_city"`
	AdrZip     string `xml:"adr_zip"`
	AdrCountry string `xml:"adr_country"`
	AdrAdif    string `xml:"adr_adif"`
	District   string `xml:"district"`
	UsState    string `xml:"us_state"`
	UsCountry  string `xml:"us_country"`
	Oblast     string `xml:"oblast"`
	Dok        string `xml:"dok"`
	Iota       string `xml:"iota"`
	QslVia     string `xml:"qsl_via"`
	Lotw       string `xml:"lotw"`
	Eqsl       string `xml:"eqsl"`
	Qsl        string `xml:"qsl"`
	QslDirect  string `xml:"qsl_direct"`
	Email      string `xml:"email"`
	Jabber     string `xml:"jabber"`
	Icq        string `xml:"icq"`
	Msn        string `xml:"msn"`
	Skype      string `xml:"skype"`
	BirthYear  string `xml:"birth_year"`
	LicYear    string `xml:"lic_year"`
	Picture    string `xml:"picture"`
	Latitude   string `xml:"latitude"`
	Longitude  string `xml:"longitude"`
	Continent  string `xml:"continent"`
	UtcOffset  string `xml:"utc_offset"`
	Facebook   string `xml:"facebook"`
	Twitter    string `xml:"twitter"`
	Gplus      string `xml:"gplus"`
	Youtube    string `xml:"youtube"`
	Linkedin   string `xml:"linkedin"`
	Flicker    string `xml:"flicker"`
	Vimeo      string `xml:"vimeo"`
}

type RecentActivitySearchResponse struct {
	HamQTH xml.Name                    `xml:"HamQTH"`
	Search RecentActivitySearchContent `xml:"search"`
}

type RecentActivitySearchContent struct {
	Activity    ActivityData `xml:"activity"`
	LogActivity LogData      `xml:"log_activity"`
	LogBook     LogData      `xml:"logbook"`
}

type ActivityData struct {
	Data []ActivityContent `xml:"data"`
}

type ActivityContent struct {
	Source   string `xml:"source"`
	Spotter  string `xml:"spotter"`
	Callsign string `xml:"callsign"`
	Note     string `xml:"note"`
	Freq     string `xml:"freq"`
	Date     string `xml:"date"`
	Time     string `xml:"time"`
}

type LogData struct {
	Data []LogContent `xml:"data"`
}

type LogContent struct {
	Callsign string `xml:"callsign"`
	Mode     string `xml:"mode"`
	Band     string `xml:"band"`
	Date     string `xml:"date"`
}

type ReverseBeaconResponse struct {
	RbnSpots []ReverseBeaconSpot
}

type ReverseBeaconSpot struct {
	DxCall string            `json:"dxcall"`
	Freq   string            `json:"freq"`
	Mode   string            `json:"mode"`
	Age    int               `json:"age"`
	Lsn    map[string]string `json:"lsn"`
}
