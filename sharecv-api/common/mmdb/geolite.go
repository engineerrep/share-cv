package mmdb

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

var db *geoip2.Reader

func Init(ipDataPath string) {
	var err error
	db, err = geoip2.Open(ipDataPath)
	if err != nil {
		fmt.Printf("open country data error data path: %s error: %v", ipDataPath, err)
	}
}

func SearchContinent(ip string) (string, string, error) {
	//db, err := geoip2.Open(ipDataPath)
	//if err != nil {
	//	fmt.Printf("open country data error data path: %s error: %v", ipDataPath, err)
	//}
	//defer db.Close()

	parseIp := net.ParseIP(ip)
	country, err := db.Country(parseIp)
	if err != nil {
		fmt.Printf("ip get country error: %v", err)
		return "", "", err
	}
	return country.Continent.Names["zh-CN"], country.Country.Names["zh-CN"], nil
}

func Close() {
	defer db.Close()
}
