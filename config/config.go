package config

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/strattonw/aug/input"
)

type Configuration struct {
	Group      string
	Start      string
	End        string
	Filter     string
}

func (c *Configuration) FilterLogEventsInput() *cloudwatchlogs.FilterLogEventsInput {
	return input.NewFilterLogEventsInput(c.Group, c.Start, c.End, c.Filter)
}
