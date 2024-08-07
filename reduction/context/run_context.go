package context

import "path"

type RunContext struct {
	Language     string
	InputFile    string
	TestScript   string
	ReductionDir string
	BestSize     int // todo determine a comparable metric
}

func (ctx RunContext) InputFilePath() string {
	return path.Join(ctx.ReductionDir, ctx.InputFile)
}

func (ctx RunContext) TestScriptPath() string {
	return path.Join(ctx.ReductionDir, ctx.TestScript)
}
