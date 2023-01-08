 // 'ksdas;dkas'd
package greeting

import "time"

// 1. 上午 0 下午
func IsAm() bool { 
	return  time.Now().Hour()  < 12
}

