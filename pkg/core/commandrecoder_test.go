package core

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/trypophob1a/fileschecker/pkg/core/testdata"
)

func TestCommandRecorder_getCommandName(t *testing.T) {
	recorder := NewCommandRecorder()
	require.Equal(t, "check", recorder.getCommandName("commands.commandCheck"))
	require.Equal(t, "check", recorder.getCommandName("check"))
	require.NotEqual(t, "check", recorder.getCommandName("notCheck"))
}

func TestCommandRecorder_Add(t *testing.T) {
	recorder := NewCommandRecorder()
	require.Empty(t, recorder.getCommands())
	recorder.Add("Test", func() {
	})
	require.NotEmpty(t, recorder.getCommands())
	require.Contains(t, recorder.getCommands(), "Test")
}

func TestCommandRecorder_AddExecutor(t *testing.T) {
	recorder := NewCommandRecorder()
	require.Empty(t, recorder.getCommands())
	recorder.AddExecutor(testdata.NewExecutor("Test"))
	require.NotEmpty(t, recorder.getCommands())
	require.Contains(t, recorder.getCommands(), "executor")
}

func TestCommandRecorder_Listener(t *testing.T) {
	recorder := NewCommandRecorder()
	testSlice := make([]int, 0)

	recorder.Add("1", func() {
		testSlice = append(testSlice, 1)
	})
	recorder.Add("2", func() {
		testSlice = append(testSlice, 2)
	})
	recorder.Add("3", func() {
		testSlice = append(testSlice, 3)
	})

	require.Empty(t, testSlice)
	for i := 1; i < 4; i++ {
		recorder.Listener(strconv.Itoa(i))
	}
	require.Equal(t, []int{1, 2, 3}, testSlice)
}
