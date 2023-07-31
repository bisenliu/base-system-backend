package initialize

import (
	"base-system-backend/global"
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

func SnowFlake() (node *sf.Node) {
	var st time.Time
	st, err := time.Parse(time.DateOnly, global.CONFIG.SnowFlake.StartTime)
	if err != nil {
		panic(fmt.Errorf("snowfake time parse failed: %w", err))
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(global.CONFIG.SnowFlake.MachineId)
	if err != nil {
		panic(fmt.Errorf("snowfake newnode failed: %w", err))
	}
	return
}
