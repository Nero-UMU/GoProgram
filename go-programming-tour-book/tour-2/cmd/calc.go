/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"tool/internal/timer"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			fmt.Printf("timer.GetCalculateTime err: %v", err)
		}
		fmt.Printf("计算结果: %s, 时间戳: %d\n", t.Format(layout), t.Unix())
	},
}

var calculateTime string
var duration string

func init() {
	timeCmd.AddCommand(calcCmd)
	calcCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "计算需要的时间,有效单位为时间戳或者已格式化后的时间")
	calcCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间,有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}
