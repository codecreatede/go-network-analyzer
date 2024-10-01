package main

/*

implementing the tcp, icp, uid connection check for the systemd to see the
system.d ctrl start service for the login and the host port check for
containers.

*/

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var (
	tcp  string
	udp  string
	tcp4 string
	tcp6 string
	// port string
)

var rootCmd = &cobra.Command{
	Use:  "options",
	Long: "Provide the tcp, udp, 4tcp, 6tcp or the port for the check for the running containers ip address",
	Run:  flagFunc,
}

func init() {
	rootCmd.Flags().StringVarP(&tcp, "tcp", "t", "provide the tcp", "tcp as a flag")
	rootCmd.Flags().StringVarP(&udp, "udp", "u", "provide the udp", "udp as a flag")
	rootCmd.Flags().StringVarP(&tcp6, "tcp6", "s", "provide the tcp", "tcp6 as a flag")
	rootCmd.Flags().StringVarP(&tcp4, "tcp4", "f", "provide the tcp4", "tcp4 as a flag")
}

func flagFunc(cmd *cobra.Command, args []string) {
	if tcp == "tcp" {
		out, err := exec.Command("lsof", "-i", "tcp").Output()
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create("tcpfile.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(string(out))
	}
	if udp == "udp" {
		out, err := exec.Command("lsof", "-i", "udp").Output()
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create("udpfile.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(string(out))
	}
	if tcp4 == "tcp4" {
		out, err := exec.Command("lsof", "-i", "4tcp").Output()
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create("tcp4file.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(string(out))
	}
	if tcp6 == "tcp6" {
		out, err := exec.Command("lsof", "-i", "6tcp").Output()
		if err != nil {
			log.Fatal(err)
		}
		t := time.Now()
		_, month, _ := t.Date()
		file, err := os.Create("tcp6file.txt" + string(month))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(string(out))
	}
}
