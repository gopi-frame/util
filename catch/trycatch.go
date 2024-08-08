package catch

import (
	"reflect"

	"github.com/gopi-frame/exception"
)

// Catcher is used to run function and catch exceptions.
type Catcher struct {
	fn           func()
	catches      map[reflect.Type]func(error)
	defaultCatch func(error)
	final        func()
}

// Catch sets a function to catch a specific error.
func (c *Catcher) Catch(e error, fn func(error)) *Catcher {
	typ := reflect.TypeOf(e)
	if _, ok := c.catches[typ]; !ok {
		c.catches[typ] = fn
	}
	return c
}

// CatchAll sets a function to catch all errors.
func (c *Catcher) CatchAll(fn func(error)) *Catcher {
	c.defaultCatch = fn
	return c
}

// Finally sets final function which will always be called at the end.
func (c *Catcher) Finally(fn func()) *Catcher {
	c.final = fn
	return c
}

// Run runs the function and catches exceptions.
func (c *Catcher) Run() {
	defer func() {
		if c.final != nil {
			defer c.final()
		}
		if err := recover(); err != nil {
			switch exp := err.(type) {
			case error:
				if catch, ok := c.catches[reflect.TypeOf(err)]; ok {
					catch(exp)
				} else if c.defaultCatch != nil {
					c.defaultCatch(exp)
				} else {
					panic(exp)
				}
			default:
				if c.defaultCatch != nil {
					c.defaultCatch(exception.NewValueException(exp))
				} else {
					panic(exp)
				}
			}
		}
	}()
	c.fn()
}

// Try creates a new [Catcher] to hold the given function.
func Try(fn func()) *Catcher {
	t := new(Catcher)
	t.fn = fn
	t.catches = make(map[reflect.Type]func(error))
	return t
}
