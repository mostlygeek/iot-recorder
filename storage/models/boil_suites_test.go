// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Demands", testDemands)
	t.Run("Summations", testSummations)
}

func TestDelete(t *testing.T) {
	t.Run("Demands", testDemandsDelete)
	t.Run("Summations", testSummationsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Demands", testDemandsQueryDeleteAll)
	t.Run("Summations", testSummationsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Demands", testDemandsSliceDeleteAll)
	t.Run("Summations", testSummationsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Demands", testDemandsExists)
	t.Run("Summations", testSummationsExists)
}

func TestFind(t *testing.T) {
	t.Run("Demands", testDemandsFind)
	t.Run("Summations", testSummationsFind)
}

func TestBind(t *testing.T) {
	t.Run("Demands", testDemandsBind)
	t.Run("Summations", testSummationsBind)
}

func TestOne(t *testing.T) {
	t.Run("Demands", testDemandsOne)
	t.Run("Summations", testSummationsOne)
}

func TestAll(t *testing.T) {
	t.Run("Demands", testDemandsAll)
	t.Run("Summations", testSummationsAll)
}

func TestCount(t *testing.T) {
	t.Run("Demands", testDemandsCount)
	t.Run("Summations", testSummationsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Demands", testDemandsHooks)
	t.Run("Summations", testSummationsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Demands", testDemandsInsert)
	t.Run("Demands", testDemandsInsertWhitelist)
	t.Run("Summations", testSummationsInsert)
	t.Run("Summations", testSummationsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Demands", testDemandsReload)
	t.Run("Summations", testSummationsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Demands", testDemandsReloadAll)
	t.Run("Summations", testSummationsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Demands", testDemandsSelect)
	t.Run("Summations", testSummationsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Demands", testDemandsUpdate)
	t.Run("Summations", testSummationsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Demands", testDemandsSliceUpdateAll)
	t.Run("Summations", testSummationsSliceUpdateAll)
}
