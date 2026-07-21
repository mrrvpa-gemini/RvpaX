package banner

import (
	"fmt"

	"github.com/mrrvpa-gemini/rvpa/pkg/colors"
)

const BannerText = `
    ____ _    ______  ___
   / __ \ |  / / __ \/   |
  / /_/ / | / / /_/ / /| |
 / _, _/| |/ / ____/ ___ |
/_/ |_| |___/_/   /_/  |_|
`

func PrintBanner() {
	fmt.Println(colors.Cyan(BannerText))
	fmt.Println(colors.Green("Fast Web Scanning & Security Audit Framework"))
	fmt.Println(colors.Green("Written in Go - Compatible with Termux"))
	fmt.Println()
}
