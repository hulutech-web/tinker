package contracts

type Tinker interface {
	Call()
	GetDBFileContent() string
	GetQueryFileContent() string
}
