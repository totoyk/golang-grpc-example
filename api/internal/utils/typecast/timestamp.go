package typecast

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func TimestampToTime(ts *timestamppb.Timestamp) time.Time {
	return ts.AsTime()
}
