package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetOutliningSpansDepthChainedCalls(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare var router: any;
router
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]
    .get[|("/", async(ctx) =>[|{
        ctx.body = "base";
    }|])|]
    .post[|("/a", async(ctx) =>[|{
        //a
    }|])|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOutliningSpans(t)
}
