package menu

type Menu interface {
	Run(input interface{}) (string, error)
}


