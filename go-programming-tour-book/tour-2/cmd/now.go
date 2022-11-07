/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tool/internal/timer"

	"github.com/spf13/cobra"
)

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		fmt.Printf("现在时间是: %s, 时间戳为: %d\n", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
