package tcphttpclient

import "github.com/cyborg-client/Remap-server/config"

// Segment is a raw MEA segment from the MEA server
type Segment [60 * config.SegmentLength]int32

type startStopTcp int