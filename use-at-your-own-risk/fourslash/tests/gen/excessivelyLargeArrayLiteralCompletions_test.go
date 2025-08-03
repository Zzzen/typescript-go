package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestExcessivelyLargeArrayLiteralCompletions(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*
   Route exported from CloudMade;

   Map data used is Copyright (c) OpenStreetMap contributors 2010
   OpenStreetMap® is open data, licensed under the 
   Open Data Commons Open Database License (ODbL) by 
   the OpenStreetMap Foundation (OSMF).

   See full license: https://www.openstreetmap.org/copyright
*/
