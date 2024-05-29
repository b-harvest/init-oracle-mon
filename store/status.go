package store

import "fmt"

// If the oracle miss count exceeds this threshold, will return false else true
func UpdateStatus() bool {
	GlobalState.Status.OracleMissed = fmt.Sprintf(
		"%d / %d",
		GlobalState.Status.OracleMissCnt, GlobalState.Status.WindowSize,
	)

	GlobalState.Status.Uptime = fmt.Sprintf(
		"%d%%",
		(GlobalState.Status.WindowSize-GlobalState.Status.OracleMissCnt)*100/GlobalState.Status.WindowSize,
	)

	if GlobalState.Status.OracleMissCnt > Threshold {
		// Reset oracle miss count
		// For preventing continuous alert
		GlobalState.Status.OracleMissCnt = 0

		GlobalState.Status.Status = false
		return false
	}

	GlobalState.Status.Status = true
	return true
}
