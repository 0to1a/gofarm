package framework

import (
	"framework/app/structure"
	"log"
)

func (w *Utils) UseModule(module structure.ModularStruct) {
	for _, moduleTarget := range module.Depending {
		isDetect := false
		for _, existModule := range listModule {
			if existModule.Name == moduleTarget.Name {
				if moduleTarget.MinVersion > existModule.Version && moduleTarget.MinVersion > 0 {
					log.Fatalln("Module '"+module.Name+"' incompatible,", "target version:", moduleTarget.MinVersion, "exist version:", existModule.Version)
				}
				if moduleTarget.MaxVersion < existModule.Version && moduleTarget.MaxVersion > 0 {
					log.Fatalln("Module '"+module.Name+"' incompatible,", "target version:", moduleTarget.MaxVersion, "exist version:", existModule.Version)
				}
				isDetect = true
				break
			}
		}
		if !isDetect {
			log.Fatalln("Module '"+module.Name+"' incompatible,", "no depending '"+moduleTarget.Name+"' included")
		}
	}
	listModule = append(listModule, &module)
}
