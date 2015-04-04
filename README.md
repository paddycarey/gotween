gotween
=======

[![Build Status](https://travis-ci.org/paddycarey/gotween.svg?branch=master)](https://travis-ci.org/paddycarey/gotween)
[![GoDoc](https://godoc.org/github.com/paddycarey/gotween?status.svg)](https://godoc.org/github.com/paddycarey/gotween)

A collection of tweening/easing functions implemented in Go.

gotween is a port of the [pytweening](https://github.com/asweigart/pytweening)
library to Go. It is mostly a straightforward line-for-line port and the code
should look very familiar to users of
[pytweening](https://github.com/asweigart/pytweening).


## Documentation

See [GoDoc](https://godoc.org/github.com/paddycarey/gotween) for full library
documentation.


## Example Usage

All tweening/easing functions are passed an argument of a float from 0.0 (for
the beginning) to 1.0 (for the end) of the tween:

```go
package main

import (
    "github.com/paddycarey/gotween"
)

def main() {

    // returns 0.5
    o, _ := gotween.Linear(0.5)

    // returns 0.75
    o, _ := gotween.Linear(0.75)

    // returns 1.0
    o, _ := gotween.Linear(1.0)

    // returns 0.25
    o, _ := gotween.EaseInQuad(0.5)

    // returns 0.5625
    o, _ := gotween.EaseInQuad(0.75)

    // returns 1.0
    o, _ := gotween.EaseInQuad(1.0)

    // returns 0.49999999999999994
    o, _ := gotween.EaseInOutSine(0.5)

    // returns 0.8535533905932737
    o, _ := gotween.EaseInOutSine(0.75)

    // returns 1.0
    o, _ := gotween.EaseInOutSine(1.0)

}
```


### Copyright & License

- Copyright © 2015 Patrick Carey (https://github.com/paddycarey)
- Licensed under the **MIT** license.
