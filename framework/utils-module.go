package framework

import (
	"embed"
	"framework/app/structure"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"log"
)

func (w *Utils) checkModuleVersion(moduleName string, targetVersion int, moduleTarget structure.ModularStruct) {
	if moduleTarget.MinVersion > targetVersion && moduleTarget.MinVersion > 0 {
		log.Panicf(errorModule1, moduleName, moduleTarget.MinVersion, targetVersion)
	}
	if moduleTarget.MaxVersion < targetVersion && moduleTarget.MaxVersion > 0 {
		log.Panicf(errorModule1, moduleName, moduleTarget.MaxVersion, targetVersion)
	}
}

func (w *Utils) UseModule(module structure.ModularStruct) {
	for _, moduleTarget := range module.Depending {
		isDetect := false
		for _, existModule := range listModule {
			if existModule.Name == moduleTarget.Name {
				w.checkModuleVersion(existModule.Name, existModule.Version, moduleTarget)
				isDetect = true
				break
			}
		}
		if !isDetect {
			log.Panicf(errorModule2, module.Name, moduleTarget.Name)
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
		log.Panic(err)
	}

	if structure.SystemConf.Database == "mysql" {
		dbMysql.MigrateDatabase(d)
	}
}
