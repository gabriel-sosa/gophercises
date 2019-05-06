package cmd

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//RootCmd root of the cli
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task manager",
	Long:  `Application for task managing`,
}

//DataBase is the address to the data base
var DataBase = "./prob_7/tasks.db"

//Execute execute the root
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//Itob transform an int to an array o bytes
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
