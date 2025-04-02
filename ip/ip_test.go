package ip

import (
	"fmt"
	"testing"
)

func TestGetRegion(t *testing.T) {
	var region *Region
	var err error
	region, _ = GetRegion("127.0.0.1")
	fmt.Printf("%#v\n", region)
	region, _ = GetRegion("8.8.8.8")
	fmt.Printf("%#v\n", region)
	region, _ = GetRegion("192.168.0.1")
	fmt.Printf("%#v\n", region)
	region, _ = GetRegion("1.2.3.4")
	fmt.Printf("%#v\n", region)
	region, _ = GetRegion("120.234.15.245")
	fmt.Printf("%#v\n", region)
	region, _ = GetRegion("58.60.186.1")
	fmt.Printf("%#v\n", region)
	region, _ = GetRegion("58.6.49.1")
	fmt.Printf("%#v\n", region)
	region, err = GetRegion("256.6.49.1")
	fmt.Printf("%#v\n", region)
	fmt.Printf("%#v\n", err)
}
