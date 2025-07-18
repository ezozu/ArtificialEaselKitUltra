// cmd/artificialeaselkitultra/main.go
package main

import (
"flag"
"log"
"os"

"artificialeaselkitultra/internal/artificialeaselkitultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeaselkitultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
