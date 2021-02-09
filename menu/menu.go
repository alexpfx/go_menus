package menu

type Menu interface {
	Run(input string) (string, error)
}



type Builder interface {
	Prompt(string) Builder
	AutoSelect(bool) Builder
	Build() Menu
}
