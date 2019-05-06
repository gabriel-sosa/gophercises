package cmd

import (
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	Long:  `Add a new task to the task manager`,
	Run: func(cmd *cobra.Command, args []string) {
		var arg string
		db, err := bolt.Open(DataBase, 0600, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer db.Close()
		if len(args) >= 1 {
			arg = strings.Join(args, " ")
		} else {
			log.Fatalln("wrong number of arguments")
			return
		}
		err = db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("tasks"))
			if err != nil {
				return err
			}
			id, _ := b.NextSequence()
			return b.Put(Itob(int(id)), []byte(arg))
		})
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("added \"" + arg + "\" to your task list")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
