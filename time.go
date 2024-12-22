package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	// Format time
	timeStr := now.Format("02 January 2006 15:04 MST")
	fmt.Println(timeStr)

	// Parse time
	timeStr = "02 January 2006 15:04 MST"
	parseTime, err := time.Parse(timeStr, "02 January 2006 15:04 MST")
	if err != nil {
		fmt.Println("Failed to parse time")
		return
	}
	fmt.Println(parseTime)

	// Duration
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()

	duration := end.Sub(start)
	fmt.Println("durasi: ", duration)

	//timer
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Menunggu 5 detik")

	<-timer.C
	fmt.Println("Timer Selesai")

	// Selisih waktu
	time.Sleep(2 * time.Second)

	elapsed := time.Since(start)
	fmt.Println("Waktu yang dibutuhkan: ", elapsed)

	// Countdown timer
	countdown := 10

	for i := countdown; i > 0; i-- {
		fmt.Printf("Countdown: %d\n\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Waktu Habis!")
}
