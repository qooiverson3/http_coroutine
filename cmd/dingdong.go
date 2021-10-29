/*
Copyright © 2021 wesley_lai

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"sync"

	r "go/angelVisit/pkg/requests"

	"github.com/spf13/cobra"
)

var amount int
var destination string

// dingdongCmd represents the dingdong command
var dingdongCmd = &cobra.Command{
	Use:   "dingdong",
	Short: "友善的叮咚拜訪",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		wg.Add(amount)

		visitor := r.Prepare{}
		for i := 0; i < amount; i++ {
			visitor.Amount = i
			visitor.Dest = destination

			go visitor.Travel(wg)
		}
		wg.Wait()
		// fmt.Println("Bye!!!")
	},
}

func init() {
	rootCmd.AddCommand(dingdongCmd)
	dingdongCmd.Flags().IntVarP(&amount, "amount", "m", amount, "開啟「線程」的數量")
	dingdongCmd.Flags().StringVarP(&destination, "destination", "d", destination, "欲訪問的目的地 [example: https://wesleylai.com]")
}
