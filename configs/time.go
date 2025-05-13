package configs

import "time"

func NewAppInitTime() {
	// # INITIAL TIME ZONE IN APPLICATION --------------------------
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}
