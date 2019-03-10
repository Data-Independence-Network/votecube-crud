// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTowns(t *testing.T) {
	t.Parallel()

	query := Towns()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTownsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTownsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Towns().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTownsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TownSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTownsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TownExists(ctx, tx, o.TownID)
	if err != nil {
		t.Errorf("Unable to check if Town exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TownExists to return true, but got false.")
	}
}

func testTownsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	townFound, err := FindTown(ctx, tx, o.TownID)
	if err != nil {
		t.Error(err)
	}

	if townFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTownsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Towns().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTownsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Towns().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTownsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	townOne := &Town{}
	townTwo := &Town{}
	if err = randomize.Struct(seed, townOne, townDBTypes, false, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}
	if err = randomize.Struct(seed, townTwo, townDBTypes, false, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = townOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = townTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Towns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTownsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	townOne := &Town{}
	townTwo := &Town{}
	if err = randomize.Struct(seed, townOne, townDBTypes, false, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}
	if err = randomize.Struct(seed, townTwo, townDBTypes, false, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = townOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = townTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func townBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func townAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func townAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func townBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func townAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func townBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func townAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func townBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func townAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Town) error {
	*o = Town{}
	return nil
}

func testTownsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Town{}
	o := &Town{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, townDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Town object: %s", err)
	}

	AddTownHook(boil.BeforeInsertHook, townBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	townBeforeInsertHooks = []TownHook{}

	AddTownHook(boil.AfterInsertHook, townAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	townAfterInsertHooks = []TownHook{}

	AddTownHook(boil.AfterSelectHook, townAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	townAfterSelectHooks = []TownHook{}

	AddTownHook(boil.BeforeUpdateHook, townBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	townBeforeUpdateHooks = []TownHook{}

	AddTownHook(boil.AfterUpdateHook, townAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	townAfterUpdateHooks = []TownHook{}

	AddTownHook(boil.BeforeDeleteHook, townBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	townBeforeDeleteHooks = []TownHook{}

	AddTownHook(boil.AfterDeleteHook, townAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	townAfterDeleteHooks = []TownHook{}

	AddTownHook(boil.BeforeUpsertHook, townBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	townBeforeUpsertHooks = []TownHook{}

	AddTownHook(boil.AfterUpsertHook, townAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	townAfterUpsertHooks = []TownHook{}
}

func testTownsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTownsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(townColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTownToManyPollsTowns(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Town
	var b, c PollsTown

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, pollsTownDBTypes, false, pollsTownColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pollsTownDBTypes, false, pollsTownColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.TownID = a.TownID
	c.TownID = a.TownID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	pollsTown, err := a.PollsTowns().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range pollsTown {
		if v.TownID == b.TownID {
			bFound = true
		}
		if v.TownID == c.TownID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TownSlice{&a}
	if err = a.L.LoadPollsTowns(ctx, tx, false, (*[]*Town)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsTowns); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PollsTowns = nil
	if err = a.L.LoadPollsTowns(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsTowns); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", pollsTown)
	}
}

func testTownToManySuburbs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Town
	var b, c Suburb

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, suburbDBTypes, false, suburbColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, suburbDBTypes, false, suburbColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.TownID = a.TownID
	c.TownID = a.TownID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	suburb, err := a.Suburbs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range suburb {
		if v.TownID == b.TownID {
			bFound = true
		}
		if v.TownID == c.TownID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TownSlice{&a}
	if err = a.L.LoadSuburbs(ctx, tx, false, (*[]*Town)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Suburbs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Suburbs = nil
	if err = a.L.LoadSuburbs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Suburbs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", suburb)
	}
}

func testTownToManyAddOpPollsTowns(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Town
	var b, c, d, e PollsTown

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, townDBTypes, false, strmangle.SetComplement(townPrimaryKeyColumns, townColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsTown{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsTownDBTypes, false, strmangle.SetComplement(pollsTownPrimaryKeyColumns, pollsTownColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*PollsTown{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPollsTowns(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.TownID != first.TownID {
			t.Error("foreign key was wrong value", a.TownID, first.TownID)
		}
		if a.TownID != second.TownID {
			t.Error("foreign key was wrong value", a.TownID, second.TownID)
		}

		if first.R.Town != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Town != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PollsTowns[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PollsTowns[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PollsTowns().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTownToManyAddOpSuburbs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Town
	var b, c, d, e Suburb

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, townDBTypes, false, strmangle.SetComplement(townPrimaryKeyColumns, townColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Suburb{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, suburbDBTypes, false, strmangle.SetComplement(suburbPrimaryKeyColumns, suburbColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Suburb{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSuburbs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.TownID != first.TownID {
			t.Error("foreign key was wrong value", a.TownID, first.TownID)
		}
		if a.TownID != second.TownID {
			t.Error("foreign key was wrong value", a.TownID, second.TownID)
		}

		if first.R.Town != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Town != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Suburbs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Suburbs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Suburbs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTownToOneStateUsingState(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Town
	var foreign State

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, townDBTypes, false, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stateDBTypes, false, stateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize State struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.StateID = foreign.StateID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.State().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.StateID != foreign.StateID {
		t.Errorf("want: %v, got %v", foreign.StateID, check.StateID)
	}

	slice := TownSlice{&local}
	if err = local.L.LoadState(ctx, tx, false, (*[]*Town)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.State == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.State = nil
	if err = local.L.LoadState(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.State == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTownToOneSetOpStateUsingState(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Town
	var b, c State

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, townDBTypes, false, strmangle.SetComplement(townPrimaryKeyColumns, townColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stateDBTypes, false, strmangle.SetComplement(statePrimaryKeyColumns, stateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stateDBTypes, false, strmangle.SetComplement(statePrimaryKeyColumns, stateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*State{&b, &c} {
		err = a.SetState(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.State != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Towns[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StateID != x.StateID {
			t.Error("foreign key was wrong value", a.StateID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StateID))
		reflect.Indirect(reflect.ValueOf(&a.StateID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StateID != x.StateID {
			t.Error("foreign key was wrong value", a.StateID, x.StateID)
		}
	}
}

func testTownsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTownsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TownSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTownsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Towns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	townDBTypes = map[string]string{`StateID`: `int8`, `TownCode`: `varchar`, `TownID`: `int8`, `TownName`: `varchar`}
	_           = bytes.MinRead
)

func testTownsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(townPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(townColumns) == len(townPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, townDBTypes, true, townPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTownsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(townColumns) == len(townPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Town{}
	if err = randomize.Struct(seed, o, townDBTypes, true, townColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, townDBTypes, true, townPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(townColumns, townPrimaryKeyColumns) {
		fields = townColumns
	} else {
		fields = strmangle.SetComplement(
			townColumns,
			townPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TownSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTownsUpsert(t *testing.T) {
	t.Parallel()

	if len(townColumns) == len(townPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Town{}
	if err = randomize.Struct(seed, &o, townDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Town: %s", err)
	}

	count, err := Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, townDBTypes, false, townPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Town struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Town: %s", err)
	}

	count, err = Towns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
