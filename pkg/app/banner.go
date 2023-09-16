package app

import (
	"fmt"
	"strings"

	"ashish.com/m/internal/utils"
)

const banner = `
_______        _   
|__   __|      | |  
   | | ___  ___| |_ 
   | |/ _ \/ __| __|
   | |  __/\__ \ |_ 
   |_|\___||___/\__| `

// CreateBanner generates a display banner for application terminal. Use the web
// tool at http://patorjk.com/software/taag/#p=display&f=Big&t=Banner for reference.
func CreateBanner(footerText string) string {
	const footerWidth = 85
	const Blue = "\033[34m"
	const Reset = "\033[0m"

	footer := fmt.Sprintf("/%s/%s", footerText, strings.Repeat("-", footerWidth-(len(footerText)+2)))
	blueFooter := fmt.Sprintf("%s%s%s%s", strings.Repeat(utils.Space, 5), Blue, footer, Reset)

	return fmt.Sprintf("%s\n%s\n", banner, blueFooter)
}
