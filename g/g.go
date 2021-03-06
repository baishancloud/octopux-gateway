package g

import (
	"log"
	"runtime"
)

// changelog:
// 0.0.1: init project
// 0.0.3: change conn pools, enlarge sending queue
// 0.0.4: use relative paths in 'import', mv conn_pool to central libs
// 0.0.5: use absolute paths in 'import'
// 0.0.6: support load balance between transfer instances
// 0.0.7: substitute common pkg for the old model pkg
// 0.0.8: do not retry in send, change send concurrent
// 0.0.9: add proc for send failure, rm git version
// 0.0.10: control sending concurrent of slow transfers
// 0.0.11: use pfc
// 0.0.12: use baishancloud pfc
// 0.0.13: 添加优雅重启支持
// 1.0.0: 修改打包方式

const (
	VERSION      = "1.0.0"
	GAUGE        = "GAUGE"
	COUNTER      = "COUNTER"
	DERIVE       = "DERIVE"
	DEFAULT_STEP = 60
	MIN_STEP     = 30
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
