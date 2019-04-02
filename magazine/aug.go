package magazine

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type CloudWatchLogsPageHandler func(*cloudwatchlogs.FilterLogEventsOutput, bool) bool

type Aug struct {
	cwl *cloudwatchlogs.CloudWatchLogs
}

func New(cwl *cloudwatchlogs.CloudWatchLogs) *Aug {
	return &Aug{
		cwl: cwl,
	}
}

func NewDefault() (*Aug, error) {
	if sess, err := session.NewSession(); err == nil {
		return New(cloudwatchlogs.New(sess)), nil
	} else {
		return &Aug{}, err
	}
}

func (aug *Aug) GetEvents(input *cloudwatchlogs.FilterLogEventsInput) error {
	c := make(chan *cloudwatchlogs.FilteredLogEvent, 10000)

	go func(c chan *cloudwatchlogs.FilteredLogEvent) {
		for e := range c {
			fmt.Println(e)
		}
	}(c)

	err := aug.GetEventsChannel(c, input)

	close(c)

	return err
}

func (aug *Aug) GetEventsChannel(c chan *cloudwatchlogs.FilteredLogEvent, input *cloudwatchlogs.FilterLogEventsInput) error {
	hp := func(page *cloudwatchlogs.FilterLogEventsOutput, lastPage bool) bool {
		for _, e := range page.Events {
			c <- e
		}

		return !lastPage
	}

	return aug.GetEventsHandledPage(hp, input)
}

func (aug *Aug) GetEventsRoutine(f func(*cloudwatchlogs.FilteredLogEvent), input *cloudwatchlogs.FilterLogEventsInput) error {
	fn := func(page *cloudwatchlogs.FilterLogEventsOutput, lastPage bool) bool {
		for _, e := range page.Events {
			go f(e)
		}

		return !lastPage
	}

	return aug.GetEventsHandledPage(fn, input)
}

func (aug *Aug) GetEventsHandledPage(fn CloudWatchLogsPageHandler, input *cloudwatchlogs.FilterLogEventsInput) error {
	return aug.cwl.FilterLogEventsPages(input, fn)
}
