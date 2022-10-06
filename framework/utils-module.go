package framework

import (
	"embed"
	"framework/app/structure"
	"github.com/golang-migrate/migrate/v4/source/iofs"
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

func (w *Utils) MigrateTools(fs embed.FS) {
	if !structure.SystemConf.UseMigration {
		return
	}

	list, _ := fs.ReadDir("migration")
	if len(list) == 0 {
		return // Error no folder migration
	}

	d, err := iofs.New(fs, "migration")
	if err != nil {
		log.Fatalln(err)
	}

	if structure.SystemConf.Database == "mysql" {
		dbMysql.MigrateDatabase(d)
	}
}
