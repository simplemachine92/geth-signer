/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"

	"github.com/spf13/cobra"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Username not specified")
		}
		address := args[0]

		fmt.Println("Signing to address:", address)
		fmt.Println("Message Input:", message)

		ctx := context.Background()
		conf := &firebase.Config{
			DatabaseURL: "https://proofofstake-91004-default-rtdb.firebaseio.com/",
		}
		// Fetch the service account key JSON file contents
		opt := option.WithCredentialsFile("service-account-key.json")

		// Initialize the app with a service account, granting admin privileges
		app, err := firebase.NewApp(ctx, conf, opt)
		if err != nil {
			log.Fatalln("Error initializing app:", err)
		}

		client, err := app.Database(ctx)
		if err != nil {
			log.Fatalln("Error initializing database client:", err)
		}

		type Signature struct {
			Message struct {
				Msg       string `json:"msg"`
				Pledge    string `json:"pledge"`
				Recipient string `json:"recipient"`
				Sender    string `json:"sender"`
				Timestamp string `json:"timestamp"`
			} `json:"message"`
			Signature string `json:"signature"`
			TypedData string `json:"typedData"`
		}

		/* type Copy struct {
			key       string
			signature Signature
		} */

		// Call our FB Realtime Database and return what matches the request query
		q := client.NewRef("PoS").OrderByKey()

		result, err := q.GetOrdered(ctx)
		if err != nil {
			log.Fatal(err)
		}

		s := make([]string, len(result))
		// Results will be logged in the increasing order of balance.
		for _, r := range result {
			var acc Signature

			if err := r.Unmarshal(&acc); err != nil {
				log.Fatal(err)
			}
			log.Printf("%s => %v\n", r.Key(), acc)

			// Put our address results in a slice, these are not comma separated like arrays
			s = append(s, r.Key())

		}

		// Print (later compare) after range function is completed and slice is populated
		log.Println("Slice", s)
	},
}

var message string

/* var creditAmount int64 */

func init() {
	rootCmd.AddCommand(keyCmd)
	keyCmd.Flags().StringVarP(&message, "message", "m", "", "Message to be signed")
	/* keyCmd.Flags().Int64VarP(&creditAmount, "amount", "a", 0, "Amount to be credited") */
	keyCmd.MarkFlagRequired("message")
	/* keyCmd.MarkFlagRequired("amount") */
}
