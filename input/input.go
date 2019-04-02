package input

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"time"
)

func NewFilterLogEventsInput(group string, start string, end string, filter string) *cloudwatchlogs.FilterLogEventsInput {
	input := cloudwatchlogs.FilterLogEventsInput{}
	input.SetInterleaved(true)
	input.SetLogGroupName(group)

	currentTime := time.Now()
	absoluteStartTime := currentTime

	if start != "" {
		st, err := getTime(start, currentTime)
		if err == nil {
			absoluteStartTime = st
		}
	}

	input.SetStartTime(aws.TimeUnixMilli(absoluteStartTime))

	if end != "" {
		et, err := getTime(end, currentTime)

		if err == nil {
			input.SetEndTime(aws.TimeUnixMilli(et))
		}
	}

	if filter != "" {
		input.SetFilterPattern(filter)
	}

	return &input
}

func getTime(timeStr string, currentTime time.Time) (time.Time, error) {
	relative, err := time.ParseDuration(timeStr)
	if err == nil {
		return currentTime.Add(relative), nil
	}

	absolute, err := time.Parse(time.RFC3339, timeStr)

	if err == nil {
		return absolute, err
	}

	return time.Time{}, errors.New("could not parse relative or absolute time")
}
