package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Sisylocke/clash-cli/controller"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)


var (
	// alter Rules ...
	ToAdd 		bool	
	ToDelete	bool
	ToModify	bool
	ToSearch	bool

	//switch proxy node
	ToSwitchNode bool
)

var rootCmd = &cobra.Command{
	Use: "clash",
	Run: func (cmd *cobra.Command, args []string) {
		ss := strings.Join(args, " ") 

		if ToAdd {
			err := controller.AddRule(ss)
			if err != nil {
				log.Fatalln(err)
			}
			color.Green("rule added")
			return 
		}

		if ToDelete {

		}

		if ToModify {

		}

		if ToSearch {

		}

		if ToSwitchNode {

		}
		// handle the situation when there are no flags
		if len(args) > 0 && args[0] != "-" {
			color.Red("Please assign a specific flag to continue...\n")
			if err := cmd.Usage(); err != nil {
				fmt.Println(err)
				os.Exit(1)
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
	reg(&ToModify, "modify", "m", false, "modify an existed piece of rule")
	reg(&ToSearch, "find", "f", false, "find a specific piece of rule")
	reg(&ToSwitchNode, "switch", "s", false, "switch to another proxy node")
}

// Execute the cobra process
func Execute()  {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}