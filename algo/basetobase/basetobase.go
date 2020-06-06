package basetobase

import (
	"github.com/eloylp/go-training/algo/basetodec"
	"github.com/eloylp/go-training/algo/dectobase"
)

func BaseToBase(input string, basea, baseb int) string {
	dec := basetodec.BaseToDec(input, basea)
	return dectobase.DecToBase(dec, baseb)
}
