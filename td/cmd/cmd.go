package cmd

import (
	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {
	return []*cobra.Command{
		AddCmd,
		CalendarCmd,
		ConfigCmd,
		ContextCmd,
		ListCmd,
		MarkCmd,
		OpenCmd,
		RankCmd,
		RemoveCmd,
		RenameCmd,
		ScheduleCmd,
	}
}
