package zip

type Zipper interface {
	Zip(source string, target string) error
	Unzip(source, target string) error
}
