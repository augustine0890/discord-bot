package monitor

import (
	"discordbot/internal/discord"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/shirou/gopsutil/v3/mem"
)

func StartMonitoring(s *discordgo.Session) {
	monitoringInterval := 2 * time.Minute
	printInterval := 24 * time.Hour

	printTicker := time.NewTicker(printInterval)
	defer printTicker.Stop()

	for {
		memInfo := getMemoryInfo()

		select {
		case <-time.After(monitoringInterval):
			if memInfo.UsedPercent > 85 {
				discord.SendAlertEmbedMessage(s, memInfo)
			}
			// continue monitoring
		case <-printTicker.C:
			// It's time to print the results
			printMemoryUsage(memInfo)
			discord.SendInfotEmbedMessage(s, memInfo)
		}
	}
}

func getMemoryInfo() *mem.VirtualMemoryStat {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error getting memory info:", err)
		return nil
	}
	return v
}

func printMemoryUsage(v *mem.VirtualMemoryStat) {
	if v == nil {
		return
	}

	fmt.Printf("Total memory: %v bytes (%.2f GB)\n", v.Total, bytesToGB(v.Total))
	fmt.Printf("Used memory: %v bytes (%.2f GB)\n", v.Used, bytesToGB(v.Used))
	fmt.Printf("Free memory: %v bytes (%.2f GB)\n", v.Free, bytesToGB(v.Free))
	fmt.Printf("Available memory: %v bytes (%.2f GB)\n", v.Free, bytesToGB(v.Available))
	fmt.Printf("Used memory percentage: %.2f%%\n", v.UsedPercent)
	fmt.Println("------")
}

func bytesToGB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}
