 package freemem

 import (
         "os"
         "fmt"
         "github.com/shirou/gopsutil/mem"
         "golang.org/x/text/language"
         "golang.org/x/text/message"
 )

func testBranch(err error) {
         if err != nil {
                 fmt.Println(err)
                 //os.Exit(-1)
         }
 }

 func dealwithErr(err error) {
         if err != nil {
                 fmt.Println(err)
                 //os.Exit(-1)
         }
 }

 func GetTotalMem() uint64 {
        vmStat, err := mem.VirtualMemory()
        dealwithErr(err)
        return vmStat.Total
 }

 func GetFreeMem() uint64 {
        vmStat, err := mem.VirtualMemory()
        dealwithErr(err)
        return vmStat.Free
 }


 func PrintHumanSize(msg string, size uint64) {
        K := size/1024
        M := size/1024/1024
        G := size/1024/1024/1024

        p := message.NewPrinter(language.English)
        if K > 0 {
                if M > 0 {
                        if G > 0 {
                                p.Printf("%s: %dB (%dKB, %dMB, %dGB)\n", msg, size, K, M, G)
                                return
                        } else {
                                p.Printf("%s: %dB (%dKB, %dMB)\n", msg, size, K, M)
                                return
                        }
                } else {
                        p.Printf("%s: %dB (%dKB)\n", msg, size, K)
                        return
                }
        }
        p.Printf("%s: %dB\n", msg, size)
 }

 func main() {
        // get Host Name
        hostname, _ := os.Hostname()

        totalMem := GetTotalMem()
        msg := "[" + hostname + "] Total Memory"
        PrintHumanSize(msg, totalMem)

        freeMem := GetFreeMem()
        msg = "[" + hostname + "]  Free Memory"
        PrintHumanSize(msg, freeMem)

//      fmt.Scanln()

 }
