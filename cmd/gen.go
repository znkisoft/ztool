package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/znkisoft/ztool/crypto"
	"github.com/znkisoft/ztool/gen"
	"log"
)

type genType string

const (
	shortID genType = "shortId"
	UUID    genType = "uuid"
	sha     genType = "sha"
)

var genStr genType

var GenCommand = &cobra.Command{
	Use:   "gen",
	Short: "util functions like crypto, time, etc.",
	Long:  "generate results of specific functions",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 1 {
			switch genStr {
			case shortID:
				output, err := gen.ShortId()
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println(fmt.Sprintf("[shortId] is: %s", output))
			case sha:
				output, err := crypto.HashPassword(args[0])
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println(fmt.Sprintf("[SHA] code is: %s", output))
			default:
				panic("not implement yet.")
			}
		} else {
			fmt.Println("args input is less than or equal to 1.")
		}
	},
}

func init() {
	GenCommand.Flags().StringVarP((*string)(&genStr), "type", "t", "", "util function choose")
	GenCommand.MarkFlagRequired("type")
	RootCmd.AddCommand(GenCommand)
}
