package decoder_vector

import (
	"github.com/koykov/bytebuf"
	"github.com/koykov/decoder"
	"github.com/koykov/vector"
)

func modCoalesce(ctx *decoder.Ctx, buf *any, val any, args []any) error {
	if len(args) < 1 {
		return decoder.ErrModPoorArgs
	}
	var root *vector.Node
	switch x := val.(type) {
	case vector.Interface:
		root = x.Root()
	case *vector.Node:
		root = x
	default:
		return vector.ErrIncompatType
	}
	if root == nil || root.Type() == vector.TypeNull {
		return nil
	}

	var node *vector.Node
	for i := 0; i < len(args); i++ {
		path := ctx.BufAcc.StakeOut().WriteX(args[i]).StakedString()
		if ctx.BufAcc.Error() != nil {
			continue
		}
		if node = root.Dot(path); node.Type() != vector.TypeNull {
			break
		}
	}
	if node == nil || node.Type() == vector.TypeNull {
		return nil
	}
	*buf = node
	return nil
}

func modMarshal(ctx *decoder.Ctx, buf *any, val any, args []any) error {
	var root *vector.Node
	var src any
	if val != nil {
		src = val
	} else if len(args) > 0 {
		src = args[0]
	}
	if src == nil {
		return nil
	}
	switch x := src.(type) {
	case vector.Interface:
		root = x.Root()
	case *vector.Node:
		root = x
	default:
		return vector.ErrIncompatType
	}
	if root == nil || root.Type() == vector.TypeNull {
		return nil
	}

	ctx.BufAcc.StakeOut()
	w := bytebuf.AccBufWriter{AccBuf: &ctx.BufAcc}
	_ = root.Marshal(&w)
	*buf = ctx.BufAcc.StakedBytes()
	return nil
}
