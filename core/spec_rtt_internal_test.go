package core

import (
	"testing"
	"time"
)

func TestRTT_first_value(t *testing.T) {
	t.Parallel()

	var (
		startingAverageRTT = time.Duration(0*float64(time.Millisecond))
		newRTT = time.Duration(10*float64(time.Millisecond))
		newAverageRTT = time.Duration(10*float64(time.Millisecond))
	)

	mon := ServerMonitor{
		averageRTT: startingAverageRTT,
		averageRTTSet: false,
	}

	expected := newAverageRTT
	actual := mon.updateAverageRTT(newRTT)

	if expected != actual {
		t.Errorf("\nStarting Avg RTT: %v\nNew RTT: %v\nExpected New Avg RTT: %v\nActual New Avg RTT: %v",
			startingAverageRTT.Seconds()*1000,
			newRTT.Seconds()*1000,
			expected.Seconds()*1000,
			actual.Seconds()*1000,
		)
	}
}

func TestRTT_first_value_zero(t *testing.T) {
	t.Parallel()

	var (
		startingAverageRTT = time.Duration(0*float64(time.Millisecond))
		newRTT = time.Duration(0*float64(time.Millisecond))
		newAverageRTT = time.Duration(0*float64(time.Millisecond))
	)

	mon := ServerMonitor{
		averageRTT: startingAverageRTT,
		averageRTTSet: false,
	}

	expected := newAverageRTT
	actual := mon.updateAverageRTT(newRTT)

	if expected != actual {
		t.Errorf("\nStarting Avg RTT: %v\nNew RTT: %v\nExpected New Avg RTT: %v\nActual New Avg RTT: %v",
			startingAverageRTT.Seconds()*1000,
			newRTT.Seconds()*1000,
			expected.Seconds()*1000,
			actual.Seconds()*1000,
		)
	}
}

func TestRTT_value_test_1(t *testing.T) {
	t.Parallel()

	var (
		startingAverageRTT = time.Duration(0*float64(time.Millisecond))
		newRTT = time.Duration(5*float64(time.Millisecond))
		newAverageRTT = time.Duration(1*float64(time.Millisecond))
	)

	mon := ServerMonitor{
		averageRTT: startingAverageRTT,
		averageRTTSet: true,
	}

	expected := newAverageRTT
	actual := mon.updateAverageRTT(newRTT)

	if expected != actual {
		t.Errorf("\nStarting Avg RTT: %v\nNew RTT: %v\nExpected New Avg RTT: %v\nActual New Avg RTT: %v",
			startingAverageRTT.Seconds()*1000,
			newRTT.Seconds()*1000,
			expected.Seconds()*1000,
			actual.Seconds()*1000,
		)
	}
}

func TestRTT_value_test_2(t *testing.T) {
	t.Parallel()

	var (
		startingAverageRTT = time.Duration(3.1*float64(time.Millisecond))
		newRTT = time.Duration(36*float64(time.Millisecond))
		newAverageRTT = time.Duration(9.68*float64(time.Millisecond))
	)

	mon := ServerMonitor{
		averageRTT: startingAverageRTT,
		averageRTTSet: true,
	}

	expected := newAverageRTT
	actual := mon.updateAverageRTT(newRTT)

	if expected != actual {
		t.Errorf("\nStarting Avg RTT: %v\nNew RTT: %v\nExpected New Avg RTT: %v\nActual New Avg RTT: %v",
			startingAverageRTT.Seconds()*1000,
			newRTT.Seconds()*1000,
			expected.Seconds()*1000,
			actual.Seconds()*1000,
		)
	}
}

func TestRTT_value_test_3(t *testing.T) {
	t.Parallel()

	var (
		startingAverageRTT = time.Duration(9.12*float64(time.Millisecond))
		newRTT = time.Duration(9.12*float64(time.Millisecond))
		newAverageRTT = time.Duration(9.12*float64(time.Millisecond))
	)

	mon := ServerMonitor{
		averageRTT: startingAverageRTT,
		averageRTTSet: true,
	}

	expected := newAverageRTT
	actual := mon.updateAverageRTT(newRTT)

	if expected != actual {
		t.Errorf("\nStarting Avg RTT: %v\nNew RTT: %v\nExpected New Avg RTT: %v\nActual New Avg RTT: %v",
			startingAverageRTT.Seconds()*1000,
			newRTT.Seconds()*1000,
			expected.Seconds()*1000,
			actual.Seconds()*1000,
		)
	}
}

func TestRTT_value_test_4(t *testing.T) {
	t.Parallel()

	var (
		startingAverageRTT = time.Duration(1*float64(time.Millisecond))
		newRTT = time.Duration(1000*float64(time.Millisecond))
		newAverageRTT = time.Duration(200.8*float64(time.Millisecond))
	)

	mon := ServerMonitor{
		averageRTT: startingAverageRTT,
		averageRTTSet: true,
	}

	expected := newAverageRTT
	actual := mon.updateAverageRTT(newRTT)

	if expected != actual {
		t.Errorf("\nStarting Avg RTT: %v\nNew RTT: %v\nExpected New Avg RTT: %v\nActual New Avg RTT: %v",
			startingAverageRTT.Seconds()*1000,
			newRTT.Seconds()*1000,
			expected.Seconds()*1000,
			actual.Seconds()*1000,
		)
	}
}

func TestRTT_value_test_5(t *testing.T) {
	t.Parallel()

	var (
		startingAverageRTT = time.Duration(0*float64(time.Millisecond))
		newRTT = time.Duration(0.25*float64(time.Millisecond))
		newAverageRTT = time.Duration(0.05*float64(time.Millisecond))
	)

	mon := ServerMonitor{
		averageRTT: startingAverageRTT,
		averageRTTSet: true,
	}

	expected := newAverageRTT
	actual := mon.updateAverageRTT(newRTT)

	if expected != actual {
		t.Errorf("\nStarting Avg RTT: %v\nNew RTT: %v\nExpected New Avg RTT: %v\nActual New Avg RTT: %v",
			startingAverageRTT.Seconds()*1000,
			newRTT.Seconds()*1000,
			expected.Seconds()*1000,
			actual.Seconds()*1000,
		)
	}
}

