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

func testDesignPatterns(t *testing.T) {
	t.Parallel()

	query := DesignPatterns()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDesignPatternsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
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

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDesignPatternsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DesignPatterns().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDesignPatternsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DesignPatternSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDesignPatternsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DesignPatternExists(ctx, tx, o.DesignPatternID)
	if err != nil {
		t.Errorf("Unable to check if DesignPattern exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DesignPatternExists to return true, but got false.")
	}
}

func testDesignPatternsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	designPatternFound, err := FindDesignPattern(ctx, tx, o.DesignPatternID)
	if err != nil {
		t.Error(err)
	}

	if designPatternFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDesignPatternsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DesignPatterns().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDesignPatternsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DesignPatterns().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDesignPatternsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	designPatternOne := &DesignPattern{}
	designPatternTwo := &DesignPattern{}
	if err = randomize.Struct(seed, designPatternOne, designPatternDBTypes, false, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}
	if err = randomize.Struct(seed, designPatternTwo, designPatternDBTypes, false, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = designPatternOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = designPatternTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DesignPatterns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDesignPatternsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	designPatternOne := &DesignPattern{}
	designPatternTwo := &DesignPattern{}
	if err = randomize.Struct(seed, designPatternOne, designPatternDBTypes, false, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}
	if err = randomize.Struct(seed, designPatternTwo, designPatternDBTypes, false, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = designPatternOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = designPatternTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func designPatternBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func designPatternAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func designPatternAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func designPatternBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func designPatternAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func designPatternBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func designPatternAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func designPatternBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func designPatternAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DesignPattern) error {
	*o = DesignPattern{}
	return nil
}

func testDesignPatternsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DesignPattern{}
	o := &DesignPattern{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, designPatternDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DesignPattern object: %s", err)
	}

	AddDesignPatternHook(boil.BeforeInsertHook, designPatternBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	designPatternBeforeInsertHooks = []DesignPatternHook{}

	AddDesignPatternHook(boil.AfterInsertHook, designPatternAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	designPatternAfterInsertHooks = []DesignPatternHook{}

	AddDesignPatternHook(boil.AfterSelectHook, designPatternAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	designPatternAfterSelectHooks = []DesignPatternHook{}

	AddDesignPatternHook(boil.BeforeUpdateHook, designPatternBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	designPatternBeforeUpdateHooks = []DesignPatternHook{}

	AddDesignPatternHook(boil.AfterUpdateHook, designPatternAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	designPatternAfterUpdateHooks = []DesignPatternHook{}

	AddDesignPatternHook(boil.BeforeDeleteHook, designPatternBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	designPatternBeforeDeleteHooks = []DesignPatternHook{}

	AddDesignPatternHook(boil.AfterDeleteHook, designPatternAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	designPatternAfterDeleteHooks = []DesignPatternHook{}

	AddDesignPatternHook(boil.BeforeUpsertHook, designPatternBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	designPatternBeforeUpsertHooks = []DesignPatternHook{}

	AddDesignPatternHook(boil.AfterUpsertHook, designPatternAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	designPatternAfterUpsertHooks = []DesignPatternHook{}
}

func testDesignPatternsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDesignPatternsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(designPatternColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDesignPatternToManyPollsFactorsPositions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DesignPattern
	var b, c PollsFactorsPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, pollsFactorsPositionDBTypes, false, pollsFactorsPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pollsFactorsPositionDBTypes, false, pollsFactorsPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.DesignPatternID, a.DesignPatternID)
	queries.Assign(&c.DesignPatternID, a.DesignPatternID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	pollsFactorsPosition, err := a.PollsFactorsPositions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range pollsFactorsPosition {
		if queries.Equal(v.DesignPatternID, b.DesignPatternID) {
			bFound = true
		}
		if queries.Equal(v.DesignPatternID, c.DesignPatternID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DesignPatternSlice{&a}
	if err = a.L.LoadPollsFactorsPositions(ctx, tx, false, (*[]*DesignPattern)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsFactorsPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PollsFactorsPositions = nil
	if err = a.L.LoadPollsFactorsPositions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsFactorsPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", pollsFactorsPosition)
	}
}

func testDesignPatternToManyPositions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DesignPattern
	var b, c Position

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.DesignPatternID, a.DesignPatternID)
	queries.Assign(&c.DesignPatternID, a.DesignPatternID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	position, err := a.Positions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range position {
		if queries.Equal(v.DesignPatternID, b.DesignPatternID) {
			bFound = true
		}
		if queries.Equal(v.DesignPatternID, c.DesignPatternID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DesignPatternSlice{&a}
	if err = a.L.LoadPositions(ctx, tx, false, (*[]*DesignPattern)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Positions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Positions = nil
	if err = a.L.LoadPositions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Positions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", position)
	}
}

func testDesignPatternToManyAddOpPollsFactorsPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DesignPattern
	var b, c, d, e PollsFactorsPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsFactorsPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsFactorsPositionDBTypes, false, strmangle.SetComplement(pollsFactorsPositionPrimaryKeyColumns, pollsFactorsPositionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PollsFactorsPosition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPollsFactorsPositions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.DesignPatternID, first.DesignPatternID) {
			t.Error("foreign key was wrong value", a.DesignPatternID, first.DesignPatternID)
		}
		if !queries.Equal(a.DesignPatternID, second.DesignPatternID) {
			t.Error("foreign key was wrong value", a.DesignPatternID, second.DesignPatternID)
		}

		if first.R.DesignPattern != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.DesignPattern != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PollsFactorsPositions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PollsFactorsPositions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PollsFactorsPositions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDesignPatternToManySetOpPollsFactorsPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DesignPattern
	var b, c, d, e PollsFactorsPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsFactorsPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsFactorsPositionDBTypes, false, strmangle.SetComplement(pollsFactorsPositionPrimaryKeyColumns, pollsFactorsPositionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetPollsFactorsPositions(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.PollsFactorsPositions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetPollsFactorsPositions(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.PollsFactorsPositions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.DesignPatternID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.DesignPatternID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.DesignPatternID, d.DesignPatternID) {
		t.Error("foreign key was wrong value", a.DesignPatternID, d.DesignPatternID)
	}
	if !queries.Equal(a.DesignPatternID, e.DesignPatternID) {
		t.Error("foreign key was wrong value", a.DesignPatternID, e.DesignPatternID)
	}

	if b.R.DesignPattern != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.DesignPattern != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.DesignPattern != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.DesignPattern != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.PollsFactorsPositions[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.PollsFactorsPositions[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDesignPatternToManyRemoveOpPollsFactorsPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DesignPattern
	var b, c, d, e PollsFactorsPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsFactorsPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsFactorsPositionDBTypes, false, strmangle.SetComplement(pollsFactorsPositionPrimaryKeyColumns, pollsFactorsPositionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddPollsFactorsPositions(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.PollsFactorsPositions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemovePollsFactorsPositions(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.PollsFactorsPositions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.DesignPatternID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.DesignPatternID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.DesignPattern != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.DesignPattern != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.DesignPattern != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.DesignPattern != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.PollsFactorsPositions) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.PollsFactorsPositions[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.PollsFactorsPositions[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDesignPatternToManyAddOpPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DesignPattern
	var b, c, d, e Position

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Position{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Position{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPositions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.DesignPatternID, first.DesignPatternID) {
			t.Error("foreign key was wrong value", a.DesignPatternID, first.DesignPatternID)
		}
		if !queries.Equal(a.DesignPatternID, second.DesignPatternID) {
			t.Error("foreign key was wrong value", a.DesignPatternID, second.DesignPatternID)
		}

		if first.R.DesignPattern != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.DesignPattern != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Positions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Positions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Positions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDesignPatternToManySetOpPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DesignPattern
	var b, c, d, e Position

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Position{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetPositions(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Positions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetPositions(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Positions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.DesignPatternID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.DesignPatternID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.DesignPatternID, d.DesignPatternID) {
		t.Error("foreign key was wrong value", a.DesignPatternID, d.DesignPatternID)
	}
	if !queries.Equal(a.DesignPatternID, e.DesignPatternID) {
		t.Error("foreign key was wrong value", a.DesignPatternID, e.DesignPatternID)
	}

	if b.R.DesignPattern != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.DesignPattern != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.DesignPattern != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.DesignPattern != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Positions[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Positions[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDesignPatternToManyRemoveOpPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DesignPattern
	var b, c, d, e Position

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Position{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddPositions(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Positions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemovePositions(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Positions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.DesignPatternID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.DesignPatternID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.DesignPattern != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.DesignPattern != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.DesignPattern != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.DesignPattern != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Positions) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Positions[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Positions[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDesignPatternsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
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

func testDesignPatternsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DesignPatternSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDesignPatternsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DesignPatterns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	designPatternDBTypes = map[string]string{`DesignPatternID`: `int8`, `DesignPatternName`: `varchar`}
	_                    = bytes.MinRead
)

func testDesignPatternsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(designPatternPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(designPatternColumns) == len(designPatternPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDesignPatternsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(designPatternColumns) == len(designPatternPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DesignPattern{}
	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, designPatternDBTypes, true, designPatternPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(designPatternColumns, designPatternPrimaryKeyColumns) {
		fields = designPatternColumns
	} else {
		fields = strmangle.SetComplement(
			designPatternColumns,
			designPatternPrimaryKeyColumns,
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

	slice := DesignPatternSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDesignPatternsUpsert(t *testing.T) {
	t.Parallel()

	if len(designPatternColumns) == len(designPatternPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DesignPattern{}
	if err = randomize.Struct(seed, &o, designPatternDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DesignPattern: %s", err)
	}

	count, err := DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, designPatternDBTypes, false, designPatternPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DesignPattern: %s", err)
	}

	count, err = DesignPatterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
