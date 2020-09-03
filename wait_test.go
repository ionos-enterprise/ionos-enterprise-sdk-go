package profitbricks

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var TestTimeout = time.Second * 30

func TestRunConditionWithCrashProtection(t *testing.T) {
	t.Run("with string panic", func(t *testing.T) {
		ok, err := runConditionWithCrashProtection(func() (bool, error) {
			panic("oops, a string panic")
		})
		require.Errorf(t, err, "expected an error")
		require.Containsf(t, err.Error(), "string panic", "expected err to contain 'string panic', got %v", err)
		require.False(t, ok)
	})
	t.Run("with error panic", func(t *testing.T) {
		ok, err := runConditionWithCrashProtection(func() (bool, error) {
			panic(errors.New("oops, an error panic"))
		})
		require.Errorf(t, err, "expected an error")
		require.Containsf(t, err.Error(), "error panic", "expected err to contain 'error panic', got %v", err)
		require.False(t, ok)
	})
	t.Run("with unknown panic", func(t *testing.T) {
		ok, err := runConditionWithCrashProtection(func() (bool, error) {
			panic(1)
		})
		require.Errorf(t, err, "expected an error")
		require.Containsf(t, err.Error(), "unknown panic", "expected err to contain 'unknown panic', got %v", err)
		require.False(t, ok)

	})
}

func TestPoller(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	w := poller(time.Millisecond, 2*time.Millisecond)
	ch := w(done)
	count := 0
DRAIN:
	for {
		select {
		case _, open := <-ch:
			if !open {
				break DRAIN
			}
			count++
		case <-time.After(TestTimeout):
			t.Errorf("unexpected timeout after poll")
		}
	}
	assert.Less(t, count, 3, "expected up to three values, got %d", count)
}

type fakePoller struct {
	max  int
	used int32
	wg   sync.WaitGroup
}

func fakeTicker(max int, used *int32, doneFunc func()) WaitFunc {
	return func(done <-chan struct{}) <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			defer doneFunc()
			defer close(ch)
			for i := 0; i < max; i++ {
				select {
				case ch <- struct{}{}:
				case <-done:
					return
				}
				if used != nil {
					atomic.AddInt32(used, 1)
				}
			}
		}()
		return ch
	}
}

func (fp *fakePoller) GetWaitFunc() WaitFunc {
	fp.wg.Add(1)
	return fakeTicker(fp.max, &fp.used, fp.wg.Done)
}

func TestWaitFor(t *testing.T) {
	var invocations int
	tests := []struct {
		name    string
		F       ConditionFunc
		Ticks   int
		Invoked int
		Err     bool
	}{
		{
			name: "invoked once",
			F: ConditionFunc(func() (bool, error) {
				invocations++
				return true, nil
			}),
			Ticks:   2,
			Invoked: 1,
			Err:     false,
		}, {
			name: "invoked and returns a timeout",
			F: ConditionFunc(func() (bool, error) {
				invocations++
				return false, nil
			}),
			Ticks:   2,
			Invoked: 3,
			Err:     true,
		}, {
			name: "returns an error",
			F: ConditionFunc(func() (bool, error) {
				invocations++
				return false, errors.New("oops, something went wrong")
			}),
			Ticks:   2,
			Invoked: 1,
			Err:     true,
		},
	}
	for _, tt := range tests {
		invocations = 0
		ticker := fakeTicker(tt.Ticks, nil, func() {})
		err := func() error {
			done := make(chan struct{})
			defer close(done)
			return WaitFor(ticker, tt.F, done)
		}()
		if tt.Err {
			assert.Errorf(t, err, "Expected error")
		} else {
			assert.NoError(t, err, "Expected no error")
		}
		assert.Equal(t, tt.Invoked, invocations)
	}
}

// TestWaitForWithEarlyClosingWaitFunc tests WaitFor when the WaitFunc closes its channel. Should return ErrWaitTimeout
func TestWaitForWithEarlyClosingWaitFunc(t *testing.T) {
	stopCh := make(chan struct{})
	defer close(stopCh)

	start := time.Now()
	err := WaitFor(func(done <-chan struct{}) <-chan struct{} {
		c := make(chan struct{})
		close(c)
		return c
	}, func() (bool, error) {
		return false, nil
	}, stopCh)
	duration := time.Since(start)
	require.LessOrEqual(t, int64(duration), int64(TestTimeout/2), "expected short timeout duration")
	require.EqualError(t, err, ErrWaitTimeout.Error(), "expected ErrWaitTimeout for WaitFunc")
}

// TestWaitForWithCloseChannel test WaitFor receiving a close channel, should return ErrWaitTimeout
func TestWaitForWithCloseChannel(t *testing.T) {
	stopCh := make(chan struct{})
	close(stopCh)
	c := make(chan struct{})
	defer close(c)
	start := time.Now()
	err := WaitFor(func(done <-chan struct{}) <-chan struct{} {
		return c
	}, func() (bool, error) {
		return false, nil
	}, stopCh)
	duration := time.Since(start)
	assert.LessOrEqual(t, int64(duration), int64(TestTimeout/2), "expected short timeout duration")
	assert.EqualError(t, err, ErrWaitTimeout.Error(), "expected ErrWaitTimeout for WaitFunc")
}

// TestWaitForClosesStopCh verifies that after the condition func returns true, WaitFor closes the stop channel
func TestWaitForClosesStopCh(t *testing.T) {
	stopCh := make(chan struct{})
	defer close(stopCh)
	waitFunc := poller(time.Millisecond, TestTimeout)
	var doneCh <-chan struct{}

	WaitFor(func(done <-chan struct{}) <-chan struct{} {
		doneCh = done
		return waitFunc(done)
	}, func() (bool, error) {
		time.Sleep(time.Millisecond * 10)
		return true, nil
	}, stopCh)

	select {
	case _, ok := <-doneCh:
		assert.False(t, ok, "expected closed channel after WaitFunc returning")
	default:
		t.Errorf("expected an ack of the done signal")
	}
}

func TestPoll(t *testing.T) {
	invocations := 0
	f := ConditionFunc(func() (bool, error) {
		invocations++
		return true, nil
	})
	fp := fakePoller{max: 1}
	err := pollInternal(fp.GetWaitFunc(), f)
	fp.wg.Wait()
	require.NoErrorf(t, err, "expected no error, got %v", err)
	require.Equalf(t, 1, invocations, "expected exactly one invocation, got %d", invocations)
	used := atomic.LoadInt32(&fp.used)
	require.Equalf(t, int32(1), used, "expected exactly one ticks, got", used)
}

func TestPollError(t *testing.T) {
	expectedError := errors.New("oops, something went wrong")
	f := ConditionFunc(func() (bool, error) {
		return false, expectedError
	})
	fp := fakePoller{max: 1}
	err := pollInternal(fp.GetWaitFunc(), f)
	require.Errorf(t, err, "error expected")
	require.EqualErrorf(t, err, expectedError.Error(), "expected %v, got %v", expectedError, err)
	fp.wg.Wait()
	used := atomic.LoadInt32(&fp.used)
	require.Equal(t, int32(1), used, "expected exactly one tick, got %d", used)
}

func TestPollImmediate(t *testing.T) {
	invocations := 0
	f := ConditionFunc(func() (bool, error) {
		invocations++
		return true, nil
	})
	fp := fakePoller{max: 0}
	err := pollImmediateInternal(fp.GetWaitFunc(), f)
	require.NoErrorf(t, err, "expected no error, got %v", err)
	require.Equalf(t, 1, invocations, "expected exactly one invocation, got %d", invocations)
	used := atomic.LoadInt32(&fp.used)
	require.Equalf(t, int32(0), used, "expected exactly zero ticks, got", used)
}
