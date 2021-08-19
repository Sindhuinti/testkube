package scripts

import (
	"fmt"
	"time"

	"github.com/kubeshop/kubtest/pkg/ui"
	"github.com/spf13/cobra"
)

const WatchInterval = 2 * time.Second

var params map[string]string

func init() {
	StartScriptCmd.Flags().StringP("name", "n", "", "execution name, if empty will be autogenerated")
	StartScriptCmd.Flags().StringToStringVarP(&params, "param", "p", map[string]string{}, "execution envs passed to executor")

	StartScriptCmd.Flags().StringP("namespace", "s", "default", "scripts kubernetes namespace")
}

var StartScriptCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts new script",
	Long:  `Starts new script based on Script Custom Resource name, returns results to console`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Logo()

		// TODO move it to some Validator
		if len(args) == 0 {
			ui.ExitOnError("Invalid arguments", fmt.Errorf("please pass script name to run"))
		}

		scriptID := args[0]
		name := cmd.Flag("name").Value.String()
		namespace := cmd.Flag("namespace").Value.String()
		namespacedName := fmt.Sprintf("%s/%s", namespace, scriptID)

		client := GetClient(cmd)

		scriptExecution, err := client.ExecuteScript(scriptID, namespace, name, params)
		ui.ExitOnError("starting script execution "+namespacedName, err)

		scriptExecution, err = client.GetExecution(scriptID, scriptExecution.Id)
		ui.ExitOnError("watching API for script completion", err)

		result := scriptExecution.Execution.Result
		switch true {

		case scriptExecution.Execution.IsSuccesful():
			fmt.Println(result.RawOutput)
			ui.Success("Script execution completed with sucess")

		case scriptExecution.Execution.IsFailed():
			fmt.Println(result.ErrorMessage)
			ui.Errf("Script execution failed")

		}

		ui.BR()
		ui.ShellCommand(
			"Use following command to get script execution details",
			"kubectl kubtest scripts execution test "+scriptExecution.Id,
		)

	},
}
