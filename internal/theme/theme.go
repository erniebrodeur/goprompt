package theme

import (
	"fmt"
	"strconv"
)

func parseHexColor(hex string) (int, int, int) {
	rVal, _ := strconv.ParseInt(hex[1:3], 16, 32)
	gVal, _ := strconv.ParseInt(hex[3:5], 16, 32)
	bVal, _ := strconv.ParseInt(hex[5:7], 16, 32)
	return int(rVal), int(gVal), int(bVal)
}

func HexToANSI(hex string) string {
	r, g, b := parseHexColor(hex)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// Example theme map, if desired
var MonokaiDark = map[string]string{
	"dir.normal": "#FD971F",
	"git.clean":  "#A6E22E",
	"git.dirty":  "#F92672",
}

