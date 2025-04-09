package ip

import (
	"fmt"
	"testing"
)

func TestGetRegion(t *testing.T) {
	Init("C:/Workspace/erase-game/etc/data/ip2region.xdb")
	var region *Region
	// region, _ = GetRegion("127.0.0.1")
	// fmt.Printf("%#v\n", region)
	// region, _ = GetRegion("8.8.8.8")
	// fmt.Printf("%#v\n", region)
	// region, _ = GetRegion("192.168.0.1")
	// fmt.Printf("%#v\n", region)
	// region, _ = GetRegion("1.2.3.4")
	// fmt.Printf("%#v\n", region)
	region, _ = GetRegion("45.149.92.73")
	fmt.Printf("%#v\n", region)
}
