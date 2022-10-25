package commands

var StandardCommandHandlers = []Command{
	&EchoCommand{},
	&HelpCommand{},
	&TimeCommand{},
	&StatsCommand{},
	&LeetCommand{},
	&MentionedCommand{},
	//&GlobalCommand{},
}
