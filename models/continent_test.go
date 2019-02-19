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

func testContinents(t *testing.T) {
	t.Parallel()

	query := Continents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testContinentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
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

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContinentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Continents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContinentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContinentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContinentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ContinentExists(ctx, tx, o.ContinentID2)
	if err != nil {
		t.Errorf("Unable to check if Continent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContinentExists to return true, but got false.")
	}
}

func testContinentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	continentFound, err := FindContinent(ctx, tx, o.ContinentID2)
	if err != nil {
		t.Error(err)
	}

	if continentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testContinentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Continents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testContinentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Continents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContinentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	continentOne := &Continent{}
	continentTwo := &Continent{}
	if err = randomize.Struct(seed, continentOne, continentDBTypes, false, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}
	if err = randomize.Struct(seed, continentTwo, continentDBTypes, false, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = continentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = continentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Continents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContinentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	continentOne := &Continent{}
	continentTwo := &Continent{}
	if err = randomize.Struct(seed, continentOne, continentDBTypes, false, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}
	if err = randomize.Struct(seed, continentTwo, continentDBTypes, false, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = continentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = continentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func continentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func continentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func continentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func continentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func continentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func continentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func continentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func continentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func continentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Continent) error {
	*o = Continent{}
	return nil
}

func testContinentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Continent{}
	o := &Continent{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, continentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Continent object: %s", err)
	}

	AddContinentHook(boil.BeforeInsertHook, continentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	continentBeforeInsertHooks = []ContinentHook{}

	AddContinentHook(boil.AfterInsertHook, continentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	continentAfterInsertHooks = []ContinentHook{}

	AddContinentHook(boil.AfterSelectHook, continentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	continentAfterSelectHooks = []ContinentHook{}

	AddContinentHook(boil.BeforeUpdateHook, continentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	continentBeforeUpdateHooks = []ContinentHook{}

	AddContinentHook(boil.AfterUpdateHook, continentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	continentAfterUpdateHooks = []ContinentHook{}

	AddContinentHook(boil.BeforeDeleteHook, continentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	continentBeforeDeleteHooks = []ContinentHook{}

	AddContinentHook(boil.AfterDeleteHook, continentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	continentAfterDeleteHooks = []ContinentHook{}

	AddContinentHook(boil.BeforeUpsertHook, continentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	continentBeforeUpsertHooks = []ContinentHook{}

	AddContinentHook(boil.AfterUpsertHook, continentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	continentAfterUpsertHooks = []ContinentHook{}
}

func testContinentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContinentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(continentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContinentToManyCountries(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Continent
	var b, c Country

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, countryDBTypes, false, countryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, countryDBTypes, false, countryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ContinentID = a.ContinentID2
	c.ContinentID = a.ContinentID2

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	country, err := a.Countries().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range country {
		if v.ContinentID == b.ContinentID {
			bFound = true
		}
		if v.ContinentID == c.ContinentID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ContinentSlice{&a}
	if err = a.L.LoadCountries(ctx, tx, false, (*[]*Continent)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Countries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Countries = nil
	if err = a.L.LoadCountries(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Countries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", country)
	}
}

func testContinentToManyPollsContinents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Continent
	var b, c PollsContinent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, pollsContinentDBTypes, false, pollsContinentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pollsContinentDBTypes, false, pollsContinentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ContinentID = a.ContinentID2
	c.ContinentID = a.ContinentID2

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	pollsContinent, err := a.PollsContinents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range pollsContinent {
		if v.ContinentID == b.ContinentID {
			bFound = true
		}
		if v.ContinentID == c.ContinentID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ContinentSlice{&a}
	if err = a.L.LoadPollsContinents(ctx, tx, false, (*[]*Continent)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsContinents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PollsContinents = nil
	if err = a.L.LoadPollsContinents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsContinents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", pollsContinent)
	}
}

func testContinentToManyAddOpCountries(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Continent
	var b, c, d, e Country

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, continentDBTypes, false, strmangle.SetComplement(continentPrimaryKeyColumns, continentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Country{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, countryDBTypes, false, strmangle.SetComplement(countryPrimaryKeyColumns, countryColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Country{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCountries(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ContinentID2 != first.ContinentID {
			t.Error("foreign key was wrong value", a.ContinentID2, first.ContinentID)
		}
		if a.ContinentID2 != second.ContinentID {
			t.Error("foreign key was wrong value", a.ContinentID2, second.ContinentID)
		}

		if first.R.Continent != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Continent != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Countries[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Countries[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Countries().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testContinentToManyAddOpPollsContinents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Continent
	var b, c, d, e PollsContinent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, continentDBTypes, false, strmangle.SetComplement(continentPrimaryKeyColumns, continentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsContinent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsContinentDBTypes, false, strmangle.SetComplement(pollsContinentPrimaryKeyColumns, pollsContinentColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PollsContinent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPollsContinents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ContinentID2 != first.ContinentID {
			t.Error("foreign key was wrong value", a.ContinentID2, first.ContinentID)
		}
		if a.ContinentID2 != second.ContinentID {
			t.Error("foreign key was wrong value", a.ContinentID2, second.ContinentID)
		}

		if first.R.Continent != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Continent != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PollsContinents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PollsContinents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PollsContinents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testContinentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
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

func testContinentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContinentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testContinentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Continents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	continentDBTypes = map[string]string{`ContinentCode`: `varchar`, `ContinentFullName`: `varchar`, `ContinentID`: `int8`, `ContinentID2`: `int8`, `ContinentName`: `varchar`, `CreatedDate`: `timestamptz`}
	_                = bytes.MinRead
)

func testContinentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(continentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(continentColumns) == len(continentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, continentDBTypes, true, continentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testContinentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(continentColumns) == len(continentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Continent{}
	if err = randomize.Struct(seed, o, continentDBTypes, true, continentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, continentDBTypes, true, continentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(continentColumns, continentPrimaryKeyColumns) {
		fields = continentColumns
	} else {
		fields = strmangle.SetComplement(
			continentColumns,
			continentPrimaryKeyColumns,
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

	slice := ContinentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testContinentsUpsert(t *testing.T) {
	t.Parallel()

	if len(continentColumns) == len(continentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Continent{}
	if err = randomize.Struct(seed, &o, continentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Continent: %s", err)
	}

	count, err := Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, continentDBTypes, false, continentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Continent struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Continent: %s", err)
	}

	count, err = Continents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
