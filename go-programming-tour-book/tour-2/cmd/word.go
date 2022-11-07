/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"tool/internal/word"

	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var describe = strings.Join([]string{
	"该子命令可进行如下单词格式转换:",
	"1:单词全部转换为大写",
	"1:单词全部转换为大写",
	"3:下划线单词转换为首字母大写的驼峰单词",
	"3:下划线单词转换为首字母大写的驼峰单词",
	"5:驼峰单词转换为下划线单词",
}, "\n")

// wordCmd represents the word command
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  describe,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelClass(str)
		case ModeCamelCaseToUnderscore:
			content = word.CampelCaseToUnderscore(str)
		default:
			content = "暂不支持的模式,请执行 help word 查看帮助文档\n"
		}
		fmt.Printf("结果如下: %v\n", content)
	},
}

var str string
var mode int

func init() {
	rootCmd.AddCommand(wordCmd)
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().IntVarP(&mode, "mode", "m", 0, "请输入单词的转换类型")
}
