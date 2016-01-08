package logdemo_test

import (
	"./"
	"github.com/stretchr/testify/assert"
	"testing"
)

// We are going to log, so we need to catch the logs in an io.Writer.
type LogCatcher struct {
	Logs []string
	Last string
}

// Write satisfies the io.Writer interface while catching useful data.
func (lc *LogCatcher) Write(p []byte) (n int, err error) {
	s := string(p)
	lc.Logs = append(lc.Logs, s)
	lc.Last = s
	return len(p), nil
}

func TestStructLogging(t *testing.T) {

	assert := assert.New(t)

	thing := logdemo.New("one")

	// Now we can set the logger's output to something useful for testing:
	catcher := &LogCatcher{}
	thing.Logger.SetOutput(catcher)

	// If you want to test the prefix and output format you can of course
	// do that separately.  For testing the written logs it's asking a
	// lot, but we can override!
	thing.Logger.SetFlags(0)
	thing.Logger.SetPrefix("")

	// Here we log, and catch it!
	thing.Log("here")
	assert.Equal("here\n", catcher.Last, "caught first log")

	// If we have more logs we can catch them together.
	thing.Log("there")
	thing.Log("everywhere")
	expected := []string{
		"here\n",
		"there\n",
		"everywhere\n",
	}
	assert.Equal(expected, catcher.Logs, "caught all logs")

	// OK, that was fun.  Now let's test two things with the same logger;
	// this could be useful for testing that one thing happens before another,
	// without having to account for times.
	catcher = &LogCatcher{}
	t1 := logdemo.New("one")
	t2 := logdemo.New("two")
	t1.Logger.SetOutput(catcher)
	t2.Logger.SetOutput(catcher)
	t1.Logger.SetFlags(0) // NOTE: not overriding the prefix!
	t2.Logger.SetFlags(0)

	t1.Log("first")
	t2.Log("second")
	t1.Log("third")
	t2.Log("fourth")

	// They should be as:
	expected = []string{
		"thing-one first\n",
		"thing-two second\n",
		"thing-one third\n",
		"thing-two fourth\n",
	}
	assert.Equal(expected, catcher.Logs, "caught all logs for both")

}

// Not everyone gets to have a struct...
func TestPackageLogging(t *testing.T) {

	assert := assert.New(t)

	catcher := &LogCatcher{}
	logdemo.Logger.SetOutput(catcher)

	logdemo.Logger.SetFlags(0)
	logdemo.Logger.SetPrefix("")

	logdemo.Log("here")
	assert.Equal("here\n", catcher.Last, "caught first log")

	// If we have more logs we can catch them together.
	logdemo.Log("there")
	logdemo.Log("everywhere")
	expected := []string{
		"here\n",
		"there\n",
		"everywhere\n",
	}
	assert.Equal(expected, catcher.Logs, "caught all logs")

	// ...and so on.  But please, use a struct if you possibly can.

}
