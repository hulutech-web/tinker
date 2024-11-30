package contracts

type Tinker interface {
	Call()
	GetQueryFileContent() string
	GetDBFileContent() string
}
