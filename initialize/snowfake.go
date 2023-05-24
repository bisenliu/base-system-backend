package initialize

import (
	"base-system-backend/global"
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

func SnowFlake() (node *sf.Node) {
	var st time.Time
	//st, err := time.Parse("2006-01-02", global.CONFIG.SnowFlake.StartTime)
	st, err := time.Parse("2006-01-02", "2023-05-01")
	if err != nil {
		panic(fmt.Errorf("snowfake time parse failed: %s", err))
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(global.CONFIG.SnowFlake.MachineId)
	if err != nil {
		panic(fmt.Errorf("snowfake newnode failed: %s", err))
	}
	return
}
