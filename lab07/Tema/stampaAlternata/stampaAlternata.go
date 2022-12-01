package main 
import ("fmt"
			"os"
			"sort "
)

func main() {
	var sett []string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}
	input := os.Args[1:]
	x := sort.SearchStrings(sett, input[0])
	
}
