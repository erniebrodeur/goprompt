package aggregator

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog/v2"
	"github.com/cucumber/godog/v2/colors"
)

var (
	testSegments []Segment
	testResult   string
	testErr      error
)

type fakeSegment struct {
	name      string
	fail      bool
	extraSlow bool
}

func (f fakeSegment) Output(ctx context.Context) (string, error) {
	if f.fail {
		return "", errors.New("segment error")
	}
	if f.extraSlow {
		select {
		case <-ctx.Done():
			return "", errors.New("segment timed out")
		case <-time.After(500 * time.Millisecond):
			return f.name, nil
		}
	}
	return f.name, nil
}

func iHaveSegments(segmentNames string, detail string) error {
	testSegments = nil
	clean := strings.ReplaceAll(segmentNames, "[", "")
	clean = strings.ReplaceAll(clean, "]", "")
	clean = strings.ReplaceAll(clean, `"`, "")
	names := strings.Split(clean, ",")

	for _, n := range names {
		n = strings.TrimSpace(n)
		switch detail {
		case "that succeed quickly":
			testSegments = append(testSegments, fakeSegment{name: n})
		case "where \"FailSegment\" fails":
			if n == "FailSegment" || strings.HasPrefix(n, "BadSegment") {
				testSegments = append(testSegments, fakeSegment{name: n, fail: true})
			} else {
				testSegments = append(testSegments, fakeSegment{name: n})
			}
		case "that all fail":
			testSegments = append(testSegments, fakeSegment{name: n, fail: true})
		case "where \"SlowSegment\" is very slow":
			if n == "SlowSegment" {
				testSegments = append(testSegments, fakeSegment{name: n, extraSlow: true})
			} else {
				testSegments = append(testSegments, fakeSegment{name: n})
			}
		default:
			testSegments = append(testSegments, fakeSegment{name: n})
		}
	}
	return nil
}

func iBuildThePromptWithTimeout(ms int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ms)*time.Millisecond)
	defer cancel()

	testResult, testErr = BuildPrompt(ctx, testSegments, time.Duration(ms)*time.Millisecond)
	return nil
}

func theResultShouldInclude(substr string) error {
	if !strings.Contains(testResult, substr) {
		return fmt.Errorf("expected result to include %q, got %q", substr, testResult)
	}
	return nil
}

func theResultShouldBe(expected string) error {
	if testResult != expected {
		return fmt.Errorf("expected result %q, got %q", expected, testResult)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have segments \[(.+)\] (that succeed quickly|where "FailSegment" fails|that all fail|where "SlowSegment" is very slow)$`, iHaveSegments)
	s.Step(`^I build the prompt with a (\d+)ms timeout$`, iBuildThePromptWithTimeout)
	s.Step(`^the result should include "([^"]+)"$`, theResultShouldInclude)
	s.Step(`^the result should be "([^"]+)"$`, theResultShouldBe)
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "pretty",
		Output: colors.Colored(os.Stdout),
		Paths:  []string{"aggregator_fallback.feature"},
	}
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opts)
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}