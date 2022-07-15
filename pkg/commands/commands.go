package commands

type BaseCommand struct {
	Name        string
	Description string
	Run         func(commands *[]string, flags *map[string]interface{}) error
}
