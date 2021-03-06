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

func testPollsStates(t *testing.T) {
	t.Parallel()

	query := PollsStates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPollsStatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
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

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPollsStatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PollsStates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPollsStatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PollsStateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPollsStatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PollsStateExists(ctx, tx, o.PollStateID)
	if err != nil {
		t.Errorf("Unable to check if PollsState exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PollsStateExists to return true, but got false.")
	}
}

func testPollsStatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	pollsStateFound, err := FindPollsState(ctx, tx, o.PollStateID)
	if err != nil {
		t.Error(err)
	}

	if pollsStateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPollsStatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PollsStates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPollsStatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PollsStates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPollsStatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pollsStateOne := &PollsState{}
	pollsStateTwo := &PollsState{}
	if err = randomize.Struct(seed, pollsStateOne, pollsStateDBTypes, false, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}
	if err = randomize.Struct(seed, pollsStateTwo, pollsStateDBTypes, false, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = pollsStateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = pollsStateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PollsStates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPollsStatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	pollsStateOne := &PollsState{}
	pollsStateTwo := &PollsState{}
	if err = randomize.Struct(seed, pollsStateOne, pollsStateDBTypes, false, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}
	if err = randomize.Struct(seed, pollsStateTwo, pollsStateDBTypes, false, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = pollsStateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = pollsStateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func pollsStateBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func pollsStateAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func pollsStateAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func pollsStateBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func pollsStateAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func pollsStateBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func pollsStateAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func pollsStateBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func pollsStateAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PollsState) error {
	*o = PollsState{}
	return nil
}

func testPollsStatesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PollsState{}
	o := &PollsState{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, pollsStateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PollsState object: %s", err)
	}

	AddPollsStateHook(boil.BeforeInsertHook, pollsStateBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	pollsStateBeforeInsertHooks = []PollsStateHook{}

	AddPollsStateHook(boil.AfterInsertHook, pollsStateAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	pollsStateAfterInsertHooks = []PollsStateHook{}

	AddPollsStateHook(boil.AfterSelectHook, pollsStateAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	pollsStateAfterSelectHooks = []PollsStateHook{}

	AddPollsStateHook(boil.BeforeUpdateHook, pollsStateBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	pollsStateBeforeUpdateHooks = []PollsStateHook{}

	AddPollsStateHook(boil.AfterUpdateHook, pollsStateAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	pollsStateAfterUpdateHooks = []PollsStateHook{}

	AddPollsStateHook(boil.BeforeDeleteHook, pollsStateBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	pollsStateBeforeDeleteHooks = []PollsStateHook{}

	AddPollsStateHook(boil.AfterDeleteHook, pollsStateAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	pollsStateAfterDeleteHooks = []PollsStateHook{}

	AddPollsStateHook(boil.BeforeUpsertHook, pollsStateBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	pollsStateBeforeUpsertHooks = []PollsStateHook{}

	AddPollsStateHook(boil.AfterUpsertHook, pollsStateAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	pollsStateAfterUpsertHooks = []PollsStateHook{}
}

func testPollsStatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPollsStatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(pollsStateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPollsStateToOneStateUsingState(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PollsState
	var foreign State

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pollsStateDBTypes, false, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
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

	slice := PollsStateSlice{&local}
	if err = local.L.LoadState(ctx, tx, false, (*[]*PollsState)(&slice), nil); err != nil {
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

func testPollsStateToOnePollUsingPoll(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PollsState
	var foreign Poll

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pollsStateDBTypes, false, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pollDBTypes, false, pollColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Poll struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PollID = foreign.PollID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Poll().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.PollID != foreign.PollID {
		t.Errorf("want: %v, got %v", foreign.PollID, check.PollID)
	}

	slice := PollsStateSlice{&local}
	if err = local.L.LoadPoll(ctx, tx, false, (*[]*PollsState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Poll == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Poll = nil
	if err = local.L.LoadPoll(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Poll == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPollsStateToOneSetOpStateUsingState(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PollsState
	var b, c State

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pollsStateDBTypes, false, strmangle.SetComplement(pollsStatePrimaryKeyColumns, pollsStateColumnsWithoutDefault)...); err != nil {
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

		if x.R.PollsStates[0] != &a {
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
func testPollsStateToOneSetOpPollUsingPoll(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PollsState
	var b, c Poll

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pollsStateDBTypes, false, strmangle.SetComplement(pollsStatePrimaryKeyColumns, pollsStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pollDBTypes, false, strmangle.SetComplement(pollPrimaryKeyColumns, pollColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pollDBTypes, false, strmangle.SetComplement(pollPrimaryKeyColumns, pollColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Poll{&b, &c} {
		err = a.SetPoll(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Poll != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PollsStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PollID != x.PollID {
			t.Error("foreign key was wrong value", a.PollID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PollID))
		reflect.Indirect(reflect.ValueOf(&a.PollID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PollID != x.PollID {
			t.Error("foreign key was wrong value", a.PollID, x.PollID)
		}
	}
}

func testPollsStatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
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

func testPollsStatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PollsStateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPollsStatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PollsStates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	pollsStateDBTypes = map[string]string{`PollID`: `int8`, `PollStateID`: `int8`, `StateID`: `int8`}
	_                 = bytes.MinRead
)

func testPollsStatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(pollsStatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(pollsStateColumns) == len(pollsStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPollsStatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(pollsStateColumns) == len(pollsStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PollsState{}
	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, pollsStateDBTypes, true, pollsStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(pollsStateColumns, pollsStatePrimaryKeyColumns) {
		fields = pollsStateColumns
	} else {
		fields = strmangle.SetComplement(
			pollsStateColumns,
			pollsStatePrimaryKeyColumns,
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

	slice := PollsStateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPollsStatesUpsert(t *testing.T) {
	t.Parallel()

	if len(pollsStateColumns) == len(pollsStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PollsState{}
	if err = randomize.Struct(seed, &o, pollsStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PollsState: %s", err)
	}

	count, err := PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, pollsStateDBTypes, false, pollsStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PollsState struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PollsState: %s", err)
	}

	count, err = PollsStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
