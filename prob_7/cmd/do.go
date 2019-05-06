package cmd

import (
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "complete a task",
	Long:  `Mark a task as done`,
	Run: func(cmd *cobra.Command, args []string) {
		var id int
		db, err := bolt.Open(DataBase, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		if len(args) >= 1 {
			id, _ = strconv.Atoi(args[0])
		} else {
			log.Fatalln("wrong number of arguments")
			return
		}
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			return b.Delete(Itob(id))
		})
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Println("you have completed the task")
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
