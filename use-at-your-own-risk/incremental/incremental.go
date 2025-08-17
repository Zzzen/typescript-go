package incremental

import (
	"github.com/go-json-experiment/json"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/compiler"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tsoptions"
)

type BuildInfoReader interface {
	ReadBuildInfo(buildInfoFileName string) *BuildInfo
}

var _ BuildInfoReader = (*buildInfoReader)(nil)

type buildInfoReader struct {
	host compiler.CompilerHost
}

func (r *buildInfoReader) ReadBuildInfo(buildInfoFileName string) *BuildInfo {
	// Read build info file
	data, ok := r.host.FS().ReadFile(buildInfoFileName)
	if !ok {
		return nil
	}
	var buildInfo BuildInfo
	err := json.Unmarshal([]byte(data), &buildInfo)
	if err != nil {
		return nil
	}
	return &buildInfo
}

func NewBuildInfoReader(
	host compiler.CompilerHost,
) BuildInfoReader {
	return &buildInfoReader{host: host}
}

func ReadBuildInfoProgram(config *tsoptions.ParsedCommandLine, reader BuildInfoReader, host compiler.CompilerHost) *Program {
	buildInfoFileName := config.GetBuildInfoFileName()
	if buildInfoFileName == "" {
		return nil
	}

	// Read buildinFo file
	buildInfo := reader.ReadBuildInfo(buildInfoFileName)
	if buildInfo == nil || !buildInfo.IsValidVersion() || !buildInfo.IsIncremental() {
		return nil
	}

	// Convert to information that can be used to create incremental program
	incrementalProgram := &Program{
		snapshot: buildInfoToSnapshot(buildInfo, buildInfoFileName, config, host),
	}
	return incrementalProgram
}
