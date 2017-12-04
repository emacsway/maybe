package maybe_test

import (
	"errors"
	"testing"

	"github.com/xdg/maybe"
	"github.com/xdg/testy"
)

func TestMaybeString(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var good, bad, got maybe.S
	var just string
	var err error

	good = maybe.NewS("Hello")
	just, err = good.Unbox()
	is.Equal(just, "Hello")
	is.Nil(err)

	bad = maybe.ErrS(errors.New("bad string"))
	just, err = bad.Unbox()
	is.Equal(just, "")
	is.NotNil(err)
	is.Equal(err.Error(), "bad string")

	// Map S to S
	got = good.Bind(func(s string) maybe.S { return maybe.NewS(s + " World") })
	just, err = got.Unbox()
	is.Equal(just, "Hello World")
	is.Nil(err)

	// Map S to AoS
	maos := good.BindAoS(func(s string) maybe.AoS { return maybe.NewAoS([]string{s}) })
	aos, err := maos.Unbox()
	is.Equal(aos, []string{"Hello"})
	is.Nil(err)
}