package utils

import (
	"fmt"
	"hyperbot/controllers"

	"github.com/robfig/cron/v3"
)

func RunProfitJob() {
	cronJob := cron.New()
	id, errr := cronJob.AddFunc("CRON_TZ=Africa/Kigali @daily", func() {
		controllers.CronjobToCalculatodayProfit()
	})
	if errr != nil {
		fmt.Printf("Error adding profit job: %v\n", errr)
	}
	fmt.Printf("Profit job ID: %v\n", id)
	cronJob.Start()

}

func RunJobToLookupPendingPaypackTransaction() {
	cronJob := cron.New()
	id, errr := cronJob.AddFunc("CRON_TZ=Africa/Kigali @every 5m", func() {
		err := controllers.LookupAllPendingPaypackTransactions()
		if err != nil {
			fmt.Printf("Error looking up pending paypack transactions: %v\n", err)
		}
	})
	if errr != nil {
		fmt.Printf("Error adding pending transaction's job: %v\n", errr)
	}
	fmt.Printf("pending transactions job ID: %v\n", id)
	cronJob.Start()
	select {}
}

func RunCronJobs() {
	RunJobToLookupPendingPaypackTransaction()
	RunProfitJob()
}
