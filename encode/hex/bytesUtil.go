package hex

import (
	"encoding/hex"
	"strings"
)

func BytesToHumanHexString(data []byte, splitStr string) string {
	var st strings.Builder
	st.Grow(3 * len(data))
	hexStr := strings.ToUpper(hex.EncodeToString(data))
	for i,value := range hexStr {
		st.WriteString(string(value))
		if (i+1) % 2 == 0 {
			st.WriteString(splitStr)
		}
	}
	/*
	for _, value := range data {
		st.WriteString(hex.EncodeToString([]byte{value}))
		//_, _ = fmt.Fprintf(&st, "%2X", val)
		st.WriteString(splitStr)
	}
	 */
	str := st.String()
	return str[:len(str)-1]
}
