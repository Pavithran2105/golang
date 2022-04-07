// package main

// import (
// 	"log"

// 	ps "github.com/Pavithran2105/goPackages"
// )

// func main() {
// 	processList, err := ps.Processes()
// 	if err != nil {
// 		log.Println("ps.Processes() Failed, are you using windows?")
// 		return
// 	}

// 	// map ages
// 	for x := range processList {
// 		var process ps.Process
// 		process = processList[x]
// 		log.Printf("%d\t%s\n", process.Pid(), process.Executable())

// 		// do os.* stuff on the pid
// 	}
// }
package main

import (
	"fmt"
	"os"
)

func main() {
	f := os.Args
	fmt.Println(len(f), f)
}
