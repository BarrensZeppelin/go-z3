package z3

import (
	"log"
	"time"
)

var (
	cCnt int32
	sCnt int32
	oCnt int32
)

func init() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			log.Printf("C: %d, S: %d, O: %d", cCnt, sCnt, oCnt)
		}
	}()
}
