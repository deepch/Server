/*
	Autor Andrey Semochkin
*/

package pkg

import (
	"github.com/deepch/Server/internal/constant"
	"github.com/deepch/Server/internal/logging"
	"github.com/deepch/Server/internal/memory"
)

/*
	init modules
*/

//New - load all modules
func New(pkg []memory.Service) error {

	//range all modules
	for i, i2 := range pkg {

		logging.Log.Debug(constant.ConstRunLevel, i, i2.Name())
		//call New all pkg
		err := i2.New()

		if err != nil {

			return err

		}

	}

	return nil
}

//Close - unload all modules
func Close(pkg []memory.Service) error {

	//reverse all module
	for i := len(pkg) - 1; i >= 0; i-- {

		logging.Log.Debug(constant.ConstStopLevel, i, pkg[i].Name())
		//call all pkg Close
		err := pkg[i].Close()

		if err != nil {

			return err

		}

	}

	return nil
}
