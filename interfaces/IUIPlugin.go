package interfaces

// IUIPlugin ...
type IUIPlugin interface {
	ClearScreen(i ...interface{}) interface{}
	PrintPercentOfScreen(startPercent int, endPercent int, i ...interface{}) interface{}
	PrintBanner(i ...interface{}) interface{}
	PrintTable(heads []string, rows [][]string, i ...interface{}) interface{}
	Println(i ...interface{}) (int, error)
	Printf(format string, i ...interface{}) (int, error)
	YesNoQuestion(question string, i ...interface{}) bool
	YesNoQuestionf(questionf string, i ...interface{}) bool
	Question(question string, i ...interface{}) error
	Questionf(questionf string, answer *string, i ...interface{}) error
}
