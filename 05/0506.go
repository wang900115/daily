package kata

import "fmt"

func Int32ToIp(n uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", (n>>24)&0xFF, (n>>16)&0xFF, (n>>8)&0xFF, (n)&0xFF)
}
