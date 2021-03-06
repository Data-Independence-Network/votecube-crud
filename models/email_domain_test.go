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

func testEmailDomains(t *testing.T) {
	t.Parallel()

	query := EmailDomains()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEmailDomainsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
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

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailDomainsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EmailDomains().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailDomainsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmailDomainSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailDomainsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EmailDomainExists(ctx, tx, o.EmailDomainID)
	if err != nil {
		t.Errorf("Unable to check if EmailDomain exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EmailDomainExists to return true, but got false.")
	}
}

func testEmailDomainsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	emailDomainFound, err := FindEmailDomain(ctx, tx, o.EmailDomainID)
	if err != nil {
		t.Error(err)
	}

	if emailDomainFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEmailDomainsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EmailDomains().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEmailDomainsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EmailDomains().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEmailDomainsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	emailDomainOne := &EmailDomain{}
	emailDomainTwo := &EmailDomain{}
	if err = randomize.Struct(seed, emailDomainOne, emailDomainDBTypes, false, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}
	if err = randomize.Struct(seed, emailDomainTwo, emailDomainDBTypes, false, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = emailDomainOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emailDomainTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmailDomains().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEmailDomainsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	emailDomainOne := &EmailDomain{}
	emailDomainTwo := &EmailDomain{}
	if err = randomize.Struct(seed, emailDomainOne, emailDomainDBTypes, false, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}
	if err = randomize.Struct(seed, emailDomainTwo, emailDomainDBTypes, false, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = emailDomainOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emailDomainTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func emailDomainBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func emailDomainAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func emailDomainAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func emailDomainBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func emailDomainAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func emailDomainBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func emailDomainAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func emailDomainBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func emailDomainAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EmailDomain) error {
	*o = EmailDomain{}
	return nil
}

func testEmailDomainsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EmailDomain{}
	o := &EmailDomain{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, emailDomainDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EmailDomain object: %s", err)
	}

	AddEmailDomainHook(boil.BeforeInsertHook, emailDomainBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	emailDomainBeforeInsertHooks = []EmailDomainHook{}

	AddEmailDomainHook(boil.AfterInsertHook, emailDomainAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	emailDomainAfterInsertHooks = []EmailDomainHook{}

	AddEmailDomainHook(boil.AfterSelectHook, emailDomainAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	emailDomainAfterSelectHooks = []EmailDomainHook{}

	AddEmailDomainHook(boil.BeforeUpdateHook, emailDomainBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	emailDomainBeforeUpdateHooks = []EmailDomainHook{}

	AddEmailDomainHook(boil.AfterUpdateHook, emailDomainAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	emailDomainAfterUpdateHooks = []EmailDomainHook{}

	AddEmailDomainHook(boil.BeforeDeleteHook, emailDomainBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	emailDomainBeforeDeleteHooks = []EmailDomainHook{}

	AddEmailDomainHook(boil.AfterDeleteHook, emailDomainAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	emailDomainAfterDeleteHooks = []EmailDomainHook{}

	AddEmailDomainHook(boil.BeforeUpsertHook, emailDomainBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	emailDomainBeforeUpsertHooks = []EmailDomainHook{}

	AddEmailDomainHook(boil.AfterUpsertHook, emailDomainAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	emailDomainAfterUpsertHooks = []EmailDomainHook{}
}

func testEmailDomainsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmailDomainsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(emailDomainColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmailDomainToManyEmailAddresses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EmailDomain
	var b, c EmailAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, emailAddressDBTypes, false, emailAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, emailAddressDBTypes, false, emailAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.EmailDomainID = a.EmailDomainID
	c.EmailDomainID = a.EmailDomainID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	emailAddress, err := a.EmailAddresses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range emailAddress {
		if v.EmailDomainID == b.EmailDomainID {
			bFound = true
		}
		if v.EmailDomainID == c.EmailDomainID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := EmailDomainSlice{&a}
	if err = a.L.LoadEmailAddresses(ctx, tx, false, (*[]*EmailDomain)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.EmailAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.EmailAddresses = nil
	if err = a.L.LoadEmailAddresses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.EmailAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", emailAddress)
	}
}

func testEmailDomainToManyAddOpEmailAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EmailDomain
	var b, c, d, e EmailAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emailDomainDBTypes, false, strmangle.SetComplement(emailDomainPrimaryKeyColumns, emailDomainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*EmailAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, emailAddressDBTypes, false, strmangle.SetComplement(emailAddressPrimaryKeyColumns, emailAddressColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*EmailAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddEmailAddresses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.EmailDomainID != first.EmailDomainID {
			t.Error("foreign key was wrong value", a.EmailDomainID, first.EmailDomainID)
		}
		if a.EmailDomainID != second.EmailDomainID {
			t.Error("foreign key was wrong value", a.EmailDomainID, second.EmailDomainID)
		}

		if first.R.EmailDomain != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.EmailDomain != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.EmailAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.EmailAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.EmailAddresses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testEmailDomainToOneUserAccountUsingUserAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local EmailDomain
	var foreign UserAccount

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, emailDomainDBTypes, false, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userAccountDBTypes, false, userAccountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserAccount struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserAccountID = foreign.UserAccountID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.UserAccount().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UserAccountID != foreign.UserAccountID {
		t.Errorf("want: %v, got %v", foreign.UserAccountID, check.UserAccountID)
	}

	slice := EmailDomainSlice{&local}
	if err = local.L.LoadUserAccount(ctx, tx, false, (*[]*EmailDomain)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserAccount == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.UserAccount = nil
	if err = local.L.LoadUserAccount(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserAccount == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEmailDomainToOneSetOpUserAccountUsingUserAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EmailDomain
	var b, c UserAccount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emailDomainDBTypes, false, strmangle.SetComplement(emailDomainPrimaryKeyColumns, emailDomainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userAccountDBTypes, false, strmangle.SetComplement(userAccountPrimaryKeyColumns, userAccountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userAccountDBTypes, false, strmangle.SetComplement(userAccountPrimaryKeyColumns, userAccountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UserAccount{&b, &c} {
		err = a.SetUserAccount(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.UserAccount != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EmailDomains[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserAccountID != x.UserAccountID {
			t.Error("foreign key was wrong value", a.UserAccountID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserAccountID))
		reflect.Indirect(reflect.ValueOf(&a.UserAccountID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserAccountID != x.UserAccountID {
			t.Error("foreign key was wrong value", a.UserAccountID, x.UserAccountID)
		}
	}
}

func testEmailDomainsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
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

func testEmailDomainsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmailDomainSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEmailDomainsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmailDomains().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	emailDomainDBTypes = map[string]string{`CreatedAt`: `timestamptz`, `EmailDomainID`: `int8`, `EmailDomainName`: `varchar`, `UserAccountID`: `int8`}
	_                  = bytes.MinRead
)

func testEmailDomainsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(emailDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(emailDomainColumns) == len(emailDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEmailDomainsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(emailDomainColumns) == len(emailDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmailDomain{}
	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emailDomainDBTypes, true, emailDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(emailDomainColumns, emailDomainPrimaryKeyColumns) {
		fields = emailDomainColumns
	} else {
		fields = strmangle.SetComplement(
			emailDomainColumns,
			emailDomainPrimaryKeyColumns,
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

	slice := EmailDomainSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEmailDomainsUpsert(t *testing.T) {
	t.Parallel()

	if len(emailDomainColumns) == len(emailDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EmailDomain{}
	if err = randomize.Struct(seed, &o, emailDomainDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmailDomain: %s", err)
	}

	count, err := EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, emailDomainDBTypes, false, emailDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailDomain struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmailDomain: %s", err)
	}

	count, err = EmailDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
