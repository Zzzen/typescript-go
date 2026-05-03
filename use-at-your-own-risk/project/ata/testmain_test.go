package ata_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil/baseline"
)

func TestMain(m *testing.M) {
	core.ApplyDebugStackLimit()
	defer baseline.Track()()
	m.Run()
}
