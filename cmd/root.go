package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/Sisylocke/clash-cli/controller"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)


var (
	ToAdd 		bool	
	ToDelete	bool
	ToSearch	bool

	ToSwitchNode bool
	ToChangeMode bool
	ToShowStatus bool

)

var rootCmd = &cobra.Command{
	Use: "clash",
	Run: func (cmd *cobra.Command, args []string) {
		ss := strings.Join(args, " ") 

		if ToAdd {
			controller.AddRule(ss)
			controller.ReloadConfig()
			color.Green("rule added")
			return 
		}

		if ToDelete {
			controller.DeleteRule(ss)
			controller.ReloadConfig()
			color.Green("rule deleted")
			return
		}

		if ToSearch {
			controller.SearchDomain(ss)
			return
		}

		if ToSwitchNode {
			controller.SwitchNode("", ss)
			return
		}

		if ToChangeMode {

		}

		if ToShowStatus {
			ss := controller.ShowStatus()	
			color.Green("current mode is %s\n", ss.Mode)
			color.Green("proxy node is %s\n", ss.Node)
			return
		}

		// handle the situation when there are no flags
		if len(args) > 0 && args[0] != "-" {
			color.HiWhite("Please assign a valid flag to continue...\n")
			if err := cmd.Usage(); err != nil {
				// fmt.Println(err)
				// os.Exit(1)
				log.Fatalln(err)
			}
			return
		}


	},
}

func init()  {
	reg := func (v *bool, name, shorthand string, default_value bool, description string)  {
		rootCmd.PersistentFlags().BoolVarP(v, name, shorthand, default_value, description)
	}
	reg(&ToAdd, "add", "a", false, "add a new piece of rule")
	reg(&ToDelete, "delete", "d", false, "delete an existed piece of rule")
	reg(&ToSearch, "find", "f", false, "find a specific piece of rule")
	reg(&ToSwitchNode, "node", "n", false, "switch to another proxy node")
	reg(&ToChangeMode, "mode", "m", false, "change the proxy mode: GLOBAL, Rule or Direct")
	reg(&ToShowStatus, "status", "s", false, "show the current clash status")
}

// Execute the cobra process
func Execute()  {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}