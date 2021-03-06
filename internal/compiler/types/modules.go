package types

type Module struct {
	Name string
	Body []StmtIface
}

type Importer interface {
	Resolve(string) *Module
}

type Import struct {
	Module *Module
	Alias  string
}

type PartialImport struct {
	Module *Module
	Part   string
	Alias  string
}

func (mod *Module) Add(stmt StmtIface) {
	mod.Body = append(mod.Body, stmt)
}

func (imp *Import) Resolve(path string) *Module {
	return nil
}

func (pimp *PartialImport) Resolve(path string) *Module {
	return nil
}
