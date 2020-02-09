package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"syscall"
)

// store disk status
type DiskStatus struct {
	All            uint64  `json:"all"`
	Used           uint64  `json:"used"`
	Free           uint64  `json:"free"`
	PercentageUsed float64 `json:"Percentage Used of Disk"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{} // used to get filesystem statistics
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	// disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Free = fs.Bavail * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	second_pair := disk.All + disk.Free
	disk.PercentageUsed = (((float64(disk.Used) / float64(GB)) / (float64(second_pair) / float64(GB) / 2)) * 100)

	if disk.PercentageUsed <= 80 {
		hostname, error := os.Hostname()
		if error != nil {
			panic(error)
		}
		// Sending  email to admin
		body := `warning: ` + hostname + ` is over 80% of disk`
		from := "somebody@gmail.com"
		pass := ""
		to := "somebody@gmail.com"
		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Disk Storage Alert\n\n" +
			body
		err := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))

		if err != nil {
			log.Printf("smtp error: %s", err)
			return
		}
		log.Print("sent, visit ", to)
		return
	} else {
		fmt.Printf("you have %.4f GB available\n", float64(disk.Free)/float64(GB))
		return
	}
	return
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	disk := DiskUsage("/boot")
	fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
	fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	fmt.Printf("Avail: %.2f GB\n", float64(disk.Free)/float64(GB))
	fmt.Printf("Percentage used of disk : %.2f\n", disk.PercentageUsed)
}
