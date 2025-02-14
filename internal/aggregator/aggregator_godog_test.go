package aggregator

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var (
	testSegments []Segment
	testResult   string
	testErr      error
)

// fakeSegment is a simple struct implementing Segment for test purposes.
type fakeSegment struct {
	name       string
	fail       bool
	extraSlow  bool
}

func (f fakeSegment) Output(ctx context.Context) (string, error) {
	// If fail is true, return error immediately.
	if f.fail {
		return "", errors.New("segment error")
	}
	// If extraSlow is true, simulate a big delay.
	if f.extraSlow {
		select {
		case <-ctx.Done():
			return "", errors.New("segment timed out")
		case <-time.After(500 * time.Millisecond):
			return f.name, nil
		}
	}
	// Quick success scenario
	return f.name, nil
}

// Step definitions:

func iHaveSegments(segmentNames string, detail string) error {
	testSegments = nil
	// The scenario might pass JSON-like or bracketed names, e.g. ["DirSegment","GitSegment"]
	// For simplicity, strip brackets/quotes:
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
			// If the name is "FailSegment" or "BadSegmentN", we set fail=true:
			if n == "FailSegment" || strings.HasPrefix(n, "BadSegment") {
				testSegments = append(testSegments, fakeSegment{name: n, fail: true})
			} else {
				testSegments = append(testSegments, fakeSegment{name: n})
			}
		case "that all fail":
			// Mark them all as failing:
			testSegments = append(testSegments, fakeSegment{name: n, fail: true})
		case "where \"SlowSegment\" is very slow":
			if n == "SlowSegment" {
				testSegments = append(testSegments, fakeSegment{name: n, extraSlow: true})
			} else {
				testSegments = append(testSegments, fakeSegment{name: n})
			}
		default:
			// fallback
			testSegments = append(testSegments, fakeSegment{name: n})
		}
	}
	return nil
}

func iBuildThePromptWithTimeout(ms int) error {
	ctx := context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
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

// Godog integration:

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^I have segments \[(.+)\] (that succeed quickly|where "FailSegment" fails|that all fail|where "SlowSegment" is very slow)$`, iHaveSegments)
	s.Step(`^I build the prompt with a (\d+)ms timeout$`, iBuildThePromptWithTimeout)
	s.Step(`^the result should include "([^"]+)"$`, theResultShouldInclude)
	s.Step(`^the result should be "([^"]+)"$`, theResultShouldBe)
}

// If you want to run via 'go test', add a TestMain:

func TestMain(m *godog.TestingT) {
	// Optionally customize output color, paths, etc.
	godog.TestSuite{
		Name:                 "aggregator_fallback",
		TestSuiteInitializer: nil,
		ScenarioInitializer:  FeatureContext,
		Options: &godog.Options{
			Output: colors.Colored(os.Stdout),
			Paths:  []string{"aggregator_fallback.feature"},
		},
	}.Run()
}