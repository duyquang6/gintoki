package presenter

type Server interface {
	Run()
	Close()
}
