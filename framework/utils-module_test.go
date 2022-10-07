package framework

import (
	"framework/app/structure"
	"github.com/bmizerany/assert"
	"testing"
	"testing/fstest"
)

func TestCheckModuleVersion(t *testing.T) {
	target := structure.ModularStruct{
		Name:       "module name",
		MinVersion: 10,
		MaxVersion: 11,
	}

	t.Run("Module on same version", func(t *testing.T) {
		assert.Panic(t, nil, func() {
			utils.checkModuleVersion("module name", 10, target)
		})
	})
	t.Run("Module on maximum version", func(t *testing.T) {
		assert.Panic(t, nil, func() {
			utils.checkModuleVersion("module name", 11, target)
		})
	})
	t.Run("Module is older", func(t *testing.T) {
		assert.Panic(t, "Module 'module name' incompatible, target version: 10 exist version: 1", func() {
			utils.checkModuleVersion("module name", 1, target)
		})
	})
	t.Run("Module is newest", func(t *testing.T) {
		assert.Panic(t, "Module 'module name' incompatible, target version: 11 exist version: 100", func() {
			utils.checkModuleVersion("module name", 100, target)
		})
	})
}

func TestUseModule(t *testing.T) {
	module1 := structure.ModularStruct{
		Name:    "module1",
		Version: 1,
	}
	module2 := structure.ModularStruct{
		Name:    "module2",
		Version: 1,
	}

	t.Run("Module single load", func(t *testing.T) {
		listModule = []*structure.ModularStruct{}
		var expected []*structure.ModularStruct
		expected = append(expected, &module1)
		utils.UseModule(module1)

		assert.Equal(t, listModule, expected)
	})
	t.Run("Module multiple load", func(t *testing.T) {
		listModule = []*structure.ModularStruct{}
		var expected []*structure.ModularStruct
		module2.Depending = append(module2.Depending, module1)
		expected = append(expected, &module1)
		expected = append(expected, &module2)
		utils.UseModule(module1)
		utils.UseModule(module2)

		assert.Equal(t, listModule, expected)
	})
	t.Run("Module depending success", func(t *testing.T) {
		listModule = []*structure.ModularStruct{}
		var expected []*structure.ModularStruct
		module1.MinVersion = 1
		module1.MaxVersion = 0
		module2.Depending = append(module2.Depending, module1)
		expected = append(expected, &module1)
		expected = append(expected, &module2)
		utils.UseModule(module1)
		utils.UseModule(module2)

		assert.Equal(t, listModule, expected)
	})
	t.Run("No module find", func(t *testing.T) {
		listModule = []*structure.ModularStruct{}
		var expected []*structure.ModularStruct
		module1.MinVersion = 1
		module1.MaxVersion = 0
		module2.Depending = append(module2.Depending, module1)
		expected = append(expected, &module2)

		assert.Panic(t, "Module 'module2' incompatible, no depending 'module1' included", func() {
			utils.UseModule(module2)
		})
	})
}

func TestMigrateTools(t *testing.T) {
	fsSample := fstest.MapFS{
		"migration/1_test.up.sql": {
			Data: []byte("SELECT 1;"),
		},
		"migration/1_test.down.sql": {
			Data: []byte("SELECT 1;"),
		},
	}

	t.Run("Not use migration", func(t *testing.T) {
		fs := fstest.MapFS{}
		structure.SystemConf.UseMigration = false
		assert.Panic(t, nil, func() {
			utils.MigrateTools(fs)
		})
	})
	t.Run("Using migration and no folder", func(t *testing.T) {
		fs := fstest.MapFS{}
		structure.SystemConf.UseMigration = true
		assert.Panic(t, nil, func() {
			utils.MigrateTools(fs)
		})
	})
	t.Run("Using migration", func(t *testing.T) {
		structure.SystemConf.UseMigration = true
		assert.Panic(t, nil, func() {
			utils.MigrateTools(fsSample)
		})
	})
	t.Run("Using migration and mock", func(t *testing.T) {
		structure.SystemConf.UseMigration = true
		structure.SystemConf.Database = "mysql"
		dbMysql.Host = "127.0.0.1:13306"
		assert.Panic(t, "dial tcp 127.0.0.1:13306: connect: connection refused", func() {
			utils.MigrateTools(fsSample)
		})
	})
}
