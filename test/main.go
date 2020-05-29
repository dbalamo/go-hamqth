package main

import (
	"log"

	hamqthpkg "github.com/dbalamo/go-hamqth"
	"github.com/dbalamo/go-hamqth/commons"
)

func main() {

	commons.SetHttpTransportParams()
	hamqth := hamqthpkg.NewHamQth("username", "password") //you must provide a VALID username & password HERE

	DxccSearchJson_test(hamqth)
	DxccSearchXml_test(hamqth)
	DxClusterSpots_test(hamqth)
	CallbookSearch_test(hamqth)
	RecentActivitySearch_test(hamqth)
	AgregatedRBN_test(hamqth)
}

func AgregatedRBN_test(hamqth *hamqthpkg.HamQth) {
	log.Println("reverse beacons test")
	bands := []string{"40", "20"}
	fromContinents := []string{"EU"}
	mode := []string{}
	spottedContinents := []string{"EU"}
	waz := []string{"*"}
	itu := []string{"*"}
	ageSeconds := 120
	order := 3
	rbr, _ := hamqth.AgregatedRBN(bands, fromContinents, mode, spottedContinents, waz, itu, ageSeconds, order)
	for _, currSpot := range rbr.RbnSpots {
		log.Println("currSpot = ", currSpot.DxCall, " ", currSpot.Freq)
	}
}

func RecentActivitySearch_test(hamqth *hamqthpkg.HamQth) {
	log.Println("Recent Activity Search test")
	ra, _ := hamqth.RecentActivitySearch("IQ0LT", true, true, true)
	if ra != nil {
		log.Println("Activity has ", len(ra.Search.Activity.Data), " elements")
		for _, ad := range ra.Search.Activity.Data {
			log.Println("RecentActivitySearch Activity => ", ad.Date, " ", ad.Time)
		}
	}
}

func CallbookSearch_test(hamqth *hamqthpkg.HamQth) {
	log.Println("callbook search test")
	cs, _ := hamqth.CallbookSearch("IQ0LT")
	log.Println("CallbookSearch => ", cs)
}

func DxClusterSpots_test(hamqth *hamqthpkg.HamQth) {
	log.Println("Dx Cluster Spots test")
	clust, _ := hamqth.DxClusterSpots(10, hamqthpkg.Band20M)
	for i, curr := range clust {
		log.Println(i, " = ", curr)
	}
}

func DxccSearchXml_test(hamqth *hamqthpkg.HamQth) {
	log.Println("DXCC search in XML format test")
	retX, _ := hamqth.DxccSearchXml("IQ0LT")
	log.Println("RETX => ", retX)
}

func DxccSearchJson_test(hamqth *hamqthpkg.HamQth) {
	log.Println("DXCC search in json format test")
	ret, _ := hamqth.DxccSearchJson("IQ0LT")
	log.Println("RET => ", ret)
}
