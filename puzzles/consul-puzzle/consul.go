package consulpuzzle

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/go-puzzles/cores"
	"github.com/go-puzzles/cores/discover"
	"github.com/go-puzzles/plog"
	"github.com/pkg/errors"
)

type consulPuzzle struct {
}

func (cp *consulPuzzle) Name() string {
	return "ConsulHandler"
}

func (cp *consulPuzzle) StartPuzzle(ctx context.Context, opt *cores.Options) error {
	if opt.ListenerAddr == "" {
		return errors.New("Consul Handler can only be used when the service is listening on a port")
	}

	if len(opt.Tags) == 0 {
		opt.Tags = append(opt.Tags, "dev")
	}

	tags := opt.Tags[:]
	sort.Strings(tags)
	if err := discover.GetServiceFinder().RegisterServiceWithTags(opt.ServiceName, opt.ListenerAddr, tags); err != nil {
		return errors.Wrap(err, "registerConsul")
	}

	var logArgs []any
	logText := "Registered into consul success. Service=%v"
	logArgs = append(logArgs, opt.ServiceName)
	if len(tags) > 0 {
		logText = fmt.Sprintf("%v %v", logText, "Tag=%v")
		logArgs = append(logArgs, strings.Join(tags, ","))
	}

	plog.Infoc(ctx, logText, logArgs...)

	<-ctx.Done()

	discover.GetServiceFinder().Close()
	return nil
}

func (cp *consulPuzzle) Stop() error {
	return nil
}
