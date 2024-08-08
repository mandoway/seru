package plugin

const (
	MainFunction  = "Main"
	CountFunction = "CountTokens"
)

type MainFunctionType = func(fileContent []byte) ([][]byte, error)

type CountFunctionType = func(fileContent []byte) int

type Functions struct {
	MainFunction  MainFunctionType
	CountFunction CountFunctionType
}
