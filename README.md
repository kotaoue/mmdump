# mmdump
Dump tools for mermaid CLI
## Description
Extract mermaid block from md file to create mmd file and image file.
## Usage
```Go
import (
	"fmt"
	"os"

	"github.com/kotaoue/mmdump"
)

func main() {
	if err := mmdump.Dump("input.md", "output.svg"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

## Links
* [mermaidjs/mermaid.cli](https://github.com/mermaidjs/mermaid.cli)