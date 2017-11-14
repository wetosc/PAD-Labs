//Package eugddc stands for "EUGenius Distributed Data Collections"
// and contains the shared code of the application
package eugddc

import (
	"github.com/rs/zerolog/log"
)

//CheckError checks if error is not nil and if so prints the description in logs on level INFO
func CheckError(err error, info string) {
	if err != nil {
		log.Info().Msgf("%v : %v", info, err)
	}
}
