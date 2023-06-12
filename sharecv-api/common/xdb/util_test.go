package xdb

import (
	"encoding/binary"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net"
	"testing"
	"time"
)

func TestCheckIP(t *testing.T) {
	var str = "29.34.191.255"
	ip, err := CheckIP(str)
	if err != nil {
		t.Errorf("check ip `%s`: %s\n", str, err)
	}

	netIP := net.ParseIP(str).To4()
	if netIP == nil {
		t.Fatalf("parse ip `%s` failed", str)
	}

	u32 := binary.BigEndian.Uint32(netIP)
	fmt.Printf("checkip: %d, parseip: %d, isEqual: %v\n", ip, u32, ip == u32)
}

func TestLong2IP(t *testing.T) {
	var str = "29.34.191.255"
	netIP := net.ParseIP(str).To4()
	if netIP == nil {
		t.Fatalf("parse ip `%s` failed", str)
	}

	u32 := binary.BigEndian.Uint32(netIP)
	ipStr := Long2IP(u32)
	fmt.Printf("originIP: %s, Long2IP: %s, isEqual: %v\n", str, ipStr, ipStr == str)
}

func TestLoadVectorIndex(t *testing.T) {
	vIndex, err := LoadVectorIndexFromFile("./ip2region.xdb")
	if err != nil {
		fmt.Printf("failed to load vector index: %s\n", err)
		return
	}

	fmt.Printf("vIndex length: %d\n", len(vIndex))
}

func TestLoadContent(t *testing.T) {
	buff, err := LoadContentFromFile("../../deploy/xdb/ip2region.xdb")
	if err != nil {
		fmt.Printf("failed to load xdb content: %s\n", err)
		return
	}

	fmt.Printf("buff length: %d\n", len(buff))

	searcher, err := xdb.NewWithBuffer(buff)
	if err != nil {
		fmt.Printf("failed to create searcher with content: %s\n", err)
		return
	}
	defer searcher.Close()

	// do the search
	var ip = "2.88.0.0"
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return
	}

	fmt.Printf("{region: %s, took: %s}\n", region, time.Since(tStart))
}

func TestLoadHeader(t *testing.T) {
	header, err := LoadHeaderFromFile("./ip2region.xdb")
	if err != nil {
		fmt.Printf("failed to load xdb header info: %s\n", err)
		return
	}

	fmt.Printf("Version        : %d\n", header.Version)
	fmt.Printf("IndexPolicy    : %s\n", header.IndexPolicy.String())
	fmt.Printf("CreatedAt      : %d(%s)\n", header.CreatedAt, time.Unix(int64(header.CreatedAt), 0).Format(time.RFC3339))
	fmt.Printf("StartIndexPtr  : %d\n", header.StartIndexPtr)
	fmt.Printf("EndIndexPtr    : %d\n", header.EndIndexPtr)
}
