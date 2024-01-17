/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/evanraisul/book_api/api"
	"github.com/evanraisul/book_api/utils"
	"github.com/spf13/cobra"
	"net/http"
)

var port string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  `A longer description that spans `,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server Run on Port %s\n", port)

		router, _ := utils.Server()
		api.RoutesAddress(router)

		if err := http.ListenAndServe(":"+port, router); err != nil {

			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	startCmd.Flags().StringVarP(&port, "port", "p", "8080", "Run on Specific Port")
}
