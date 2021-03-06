package read_test

import (
	"context"
	"github.com/modern-go/parse"
	"github.com/modern-go/parse/read"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"testing"
	"unicode"
)

func TestUnicodeRange(t *testing.T) {
	t.Run("complete rune", test.Case(func(ctx context.Context) {
		src := parse.NewSourceString("中文c,")
		must.Equal("中文", string(read.UnicodeRange(
			src, nil, unicode.Han)))
	}))
}

func TestUnicodeRanges(t *testing.T) {
	t.Run("complex range", test.Case(func(ctx context.Context) {
		src := parse.NewSourceString("ab中文c,")
		id := read.UnicodeRanges(src, nil, nil, []*unicode.RangeTable{
			unicode.Pattern_Syntax,
			unicode.Pattern_White_Space,
		})
		must.Equal("ab中文c", string(id))
	}))
}
