package menu

type Menu interface {
	Run(input string) (string, error)
}
