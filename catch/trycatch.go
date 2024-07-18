package catch

import (
	"reflect"

	"github.com/gopi-frame/exception"
)

type catcher struct {
	fn           func()
	catches      map[reflect.Type]func(error)
	defaultCatch func(error)
	final        func()
}

// Catch set exception catcher
func (c *catcher) Catch(e error, fn func(error)) *catcher {
	typ := reflect.TypeOf(e)
	if _, ok := c.catches[typ]; !ok {
		c.catches[typ] = fn
	}
	return c
}

// CatchAll set default exception catcher
func (c *catcher) CatchAll(fn func(error)) *catcher {
	c.defaultCatch = fn
	return c
}

// Finally finally
func (c *catcher) Finally(fn func()) *catcher {
	c.final = fn
	return c
}

// Run run
func (c *catcher) Run() {
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

// Try try catch
func Try(fn func()) *catcher {
	t := new(catcher)
	t.fn = fn
	t.catches = make(map[reflect.Type]func(error))
	return t
}
