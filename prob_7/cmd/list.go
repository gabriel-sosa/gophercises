package cmd

import (
	"encoding/binary"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show the list of tasks",
	Long:  `Show a list of all the tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open(DataBase, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			c := b.Cursor()
			log.Println("You have the following task(s)")
			for k, v := c.First(); k != nil; k, v = c.Next() {
				log.Println("key:", binary.BigEndian.Uint64(k), ", value:", string(v))
			}
			return nil
		})
		// if err != nil {
		// 	log.Fatalln(err)
		// }
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
