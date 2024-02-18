package log

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoggerDynamicLvl(t *testing.T) {
	rw := bytes.NewBuffer(make([]byte, 0))
	l := New(rw, false)
	ChangeLvl(LvlDebug)
	l.Debug("foo")
	require.NotEmpty(t, rw.String())
	rw.Reset()
	ChangeLvl(LvlInfo)
	l.Debug("foo")
	require.Empty(t, rw.String())
}

func TestLoggerParse(t *testing.T) {
	for i := FirstLvl; i <= LastLvl; i++ {
		var l Lvl
		require.NoError(t, l.FromString(Lvl(i).ToString()))
		require.Equal(t, i, l)
	}
	var l Lvl
	for _, lvl := range l.Lvls() {
		var b Lvl
		require.NoError(t, b.FromString(lvl))
		require.Equal(t, lvl, b.ToString())
	}
}

func TestFatal(t *testing.T) {
	rw := bytes.NewBuffer(make([]byte, 0))
	l := New(rw, false)
	require.Equal(t, 255, l.fatal("foo"))
	require.NotEmpty(t, rw.String())
}
