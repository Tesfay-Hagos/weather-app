package support

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func TimeStampToTime(t *timestamp.Timestamp) time.Time {
	return t.AsTime()
}
func TimeToTimeStamp(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
