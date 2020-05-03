package rc

type RCGenerator interface {
	Regenerate(filePath string) error
}
