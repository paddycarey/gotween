package gotween

import (
	"errors"
	"math"
)

var version string = "0.0.1"

// EasingFunc defines a common interface that most of the easing functions
// conform to. The few that don't can be easily wrapped as required so that
// they match the interface.
type EasingFunc func(float64) (float64, error)

// GetPointOnLine returns the (x, y) coordinates of the point that has
// progressed a proportion n along the line defined by a pair of coordinates.
func GetPointOnLine(x1, y1, x2, y2, n float64) (float64, float64) {
	x := ((x2 - x1) * n) + x1
	y := ((y2 - y1) * n) + y1
	return x, y
}

// Linear is a linear tween function
func Linear(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return n, nil
}

// EaseInQuad is a quadratic tween function that begins slow and then
// accelerates.
func EaseInQuad(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return math.Pow(n, 2.0), nil
}

// EaseOutQuad is a quadratic tween function that begins fast and then
// decelerates.
func EaseOutQuad(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return -n * (n - 2), nil
}

// EaseInOutQuad is a quadratic tween function that accelerates, reaches the
// midpoint, and then decelerates.
func EaseInOutQuad(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if n < 0.5 {
		return 2 * math.Pow(n, 2.0), nil
	} else {
		n = n*2 - 1
		return -0.5 * (n*(n-2) - 1), nil
	}
}

// EaseInCubic is a cubic tween function that begins slow and then accelerates.
func EaseInCubic(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return math.Pow(n, 3.0), nil
}

// EaseOutCubic is a cubic tween function that begins fast and then
// decelerates.
func EaseOutCubic(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	n = n - 1
	return math.Pow(n, 3.0) + 1, nil
}

// EaseInOutCubic is a cubic tween function that accelerates, reaches the
// midpoint, and then decelerates.
func EaseInOutCubic(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	n = 2 * n
	if n < 1 {
		return 0.5 * math.Pow(n, 3.0), nil
	} else {
		n = n - 2
		return 0.5 * (math.Pow(n, 3.0) + 2), nil
	}
}

// EaseInQuart is a quartic tween function that begins slow and then
// accelerates.
func EaseInQuart(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return math.Pow(n, 4.0), nil
}

// EaseOutQuart is a quartic tween function that begins fast and then
// decelerates.
func EaseOutQuart(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	n = n - 1
	return -(math.Pow(n, 4.0) - 1), nil
}

// EaseInOutQuart is a quartic tween function that accelerates, reaches the
// midpoint, and then decelerates.
func EaseInOutQuart(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	n = 2 * n
	if n < 1 {
		return 0.5 * math.Pow(n, 4.0), nil
	} else {
		n = n - 2
		return -0.5 * (math.Pow(n, 4.0) - 2), nil
	}
}

// EaseInQuint is a quintic tween function that begins slow and then
// accelerates.
func EaseInQuint(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return math.Pow(n, 5.0), nil
}

// EaseOutQuint is a quintic tween function that begins fast and then
// decelerates.
func EaseOutQuint(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	n = n - 1
	return math.Pow(n, 5.0) + 1, nil
}

// EaseInOutQuint is a quintic tween function that accelerates, reaches the
// midpoint, and then decelerates.
func EaseInOutQuint(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	n = 2 * n
	if n < 1 {
		return 0.5 * math.Pow(n, 5.0), nil
	} else {
		n = n - 2
		return 0.5 * (math.Pow(n, 5.0) + 2), nil
	}
}

// EaseInSine is a sinusoidal tween function that begins slow and then
// accelerates.
func EaseInSine(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return -1*math.Cos(n*math.Pi/2) + 1, nil
}

// EaseOutSine is a sinusoidal tween function that begins fast and then
// decelerates.
func EaseOutSine(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return math.Sin(n * math.Pi / 2), nil
}

// EaseInOutSine is a sinusoidal tween function that accelerates, reaches the
// midpoint, and then decelerates.
func EaseInOutSine(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return -0.5 * (math.Cos(math.Pi*n) - 1), nil
}

// EaseInExpo is an exponential tween function that begins slow and then
// accelerates.
func EaseInExpo(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if n == 0 {
		return 0, nil
	} else {
		return math.Pow(2.0, (10 * (n - 1))), nil
	}
}

// EaseOutExpo is an exponential tween function that begins fast and then
// decelerates.
func EaseOutExpo(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if n == 1 {
		return 1, nil
	} else {
		return -(math.Pow(2.0, (-10 * n))) + 1, nil
	}
}

// EaseInOutExpo is an exponential tween function that accelerates, reaches the
// midpoint, and then decelerates.
func EaseInOutExpo(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if n == 0 {
		return 0, nil
	} else if n == 1 {
		return 1, nil
	} else {
		n = n * 2
		if n < 1 {
			return 0.5 * math.Pow(2.0, (10*(n-1))), nil
		} else {
			n -= 1
			return 0.5 * (-1*(math.Pow(2.0, (-10*n))) + 2), nil
		}
	}
}

// EaseInCirc is a circular tween function that begins slow and then
// accelerates.
func EaseInCirc(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	return -1 * (math.Sqrt(1-n*n) - 1), nil
}

// EaseOutCirc is a circular tween function that begins fast and then
// decelerates.
func EaseOutCirc(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	n -= 1
	return math.Sqrt(1 - (n * n)), nil
}

// EaseInOutCirc is a circular tween function that accelerates, reaches the
// midpoint, and then decelerates.
func EaseInOutCirc(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	n = n * 2
	if n < 1 {
		return -0.5 * (math.Sqrt(1-math.Pow(n, 2.0)) - 1), nil
	} else {
		n = n - 2
		return 0.5 * (math.Sqrt(1-math.Pow(n, 2.0)) + 1), nil
	}
}

// EaseInElastic is an elastic tween function that begins with an increasing
// wobble and then snaps into the destination.
func EaseInElastic(n, amplitude, period float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if period == 0.0 {
		period = 0.3
	}
	if amplitude == 0.0 {
		amplitude = 1
	}
	var s float64
	if amplitude < 1 {
		amplitude = 1
		s = period / 4
	} else {
		s = period / (2 * math.Pi) * math.Asin(1/amplitude)
	}

	n -= 1
	return -1 * (amplitude * math.Pow(2.0, (10*n)) * math.Sin((n-s)*(2*math.Pi)/period)), nil
}

// EaseOutElastic is an elastic tween function that overshoots the destination
// and then "rubber bands" into the destination.
func EaseOutElastic(n, amplitude, period float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if period == 0.0 {
		period = 0.3
	}
	if amplitude == 0.0 {
		amplitude = 1
	}
	var s float64
	if amplitude < 1 {
		amplitude = 1
		s = period / 4
	} else {
		s = period / (2 * math.Pi) * math.Asin(1/amplitude)
	}

	return amplitude*math.Pow(2.0, (-10*n))*math.Sin((n-s)*(2*math.Pi/period)) + 1, nil
}

// EaseInOutElastic is an elastic tween function wobbles towards the midpoint.
func EaseInOutElastic(n, amplitude, period float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if period == 0.0 {
		period = 0.5
	}
	if amplitude == 0.0 {
		amplitude = 1
	}
	var s float64
	if amplitude < 1 {
		amplitude = 1
		s = period / 4
	} else {
		s = period / (2 * math.Pi) * math.Asin(1/amplitude)
	}

	n *= 2
	if n < 1 {
		n = n - 1
		return -0.5 * (amplitude * math.Pow(2, (10*n)) * math.Sin((n-s)*2*math.Pi/period)), nil
	} else {
		n = n - 1
		return amplitude*math.Pow(2, (-10*n))*math.Sin((n-s)*2*math.Pi/period)*0.5 + 1, nil
	}
}

// EaseInBack is a tween function that backs up first at the start and then
// goes to the destination.
func EaseInBack(n, s float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if s == 0 {
		s = 1.70158
	}
	return n * n * ((s+1)*n - s), nil
}

// EaseOutBack is a tween function that overshoots the destination a little and
// then backs into the destination.
func EaseOutBack(n, s float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if s == 0 {
		s = 1.70158
	}
	n = n - 1
	return n*n*((s+1)*n+s) + 1, nil
}

// EaseInOutBounce is a "back-in" tween function that overshoots both the start
// and destination.
func EaseInOutBack(n, s float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if s == 0 {
		s = 1.70158
	}
	n = n * 2
	if n < 1 {
		s *= 1.525
		return 0.5 * (n * n * ((s+1)*n - s)), nil
	} else {
		n -= 2
		s *= 1.525
		return 0.5 * (n*n*((s+1)*n+s) + 2), nil
	}
}

// EaseInBounce is a bouncing tween function that begins bouncing and then
// jumps to the destination.
func EaseInBounce(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	eob, err := EaseOutBounce(1 - n)
	if err != nil {
		return 0.0, err
	}
	return 1 - eob, nil
}

// EaseOutBounce is a bouncing tween function that hits the destination and
// then bounces to rest.
func EaseOutBounce(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if n < (1 / 2.75) {
		return 7.5625 * n * n, nil
	} else if n < (2 / 2.75) {
		n -= (1.5 / 2.75)
		return 7.5625*n*n + 0.75, nil
	} else if n < (2.5 / 2.75) {
		n -= (2.25 / 2.75)
		return 7.5625*n*n + 0.9375, nil
	} else {
		n -= (2.65 / 2.75)
		return 7.5625*n*n + 0.984375, nil
	}
}

// EaseInOutBounce is a bouncing tween function that bounces at the start and
// end.
func EaseInOutBounce(n float64) (float64, error) {
	if n < 0.0 || n > 1.0 {
		return 0.0, errors.New("`n` must be between 0.0 and 1.0.")
	}
	if n < 0.5 {
		eob, err := EaseOutBounce(n * 2)
		if err != nil {
			return 0.0, err
		}
		return eob * 0.5, nil
	} else {
		eob, err := EaseOutBounce(n*2 - 1)
		if err != nil {
			return 0.0, err
		}
		return eob*0.5 + 0.5, nil
	}
}
