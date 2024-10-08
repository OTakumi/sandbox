// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRepositories(t *testing.T) {
	t.Parallel()

	query := Repositories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRepositoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
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

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRepositoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Repositories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRepositoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RepositorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRepositoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RepositoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Repository exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RepositoryExists to return true, but got false.")
	}
}

func testRepositoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	repositoryFound, err := FindRepository(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if repositoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRepositoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Repositories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRepositoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Repositories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRepositoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	repositoryOne := &Repository{}
	repositoryTwo := &Repository{}
	if err = randomize.Struct(seed, repositoryOne, repositoryDBTypes, false, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}
	if err = randomize.Struct(seed, repositoryTwo, repositoryDBTypes, false, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = repositoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = repositoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Repositories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRepositoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	repositoryOne := &Repository{}
	repositoryTwo := &Repository{}
	if err = randomize.Struct(seed, repositoryOne, repositoryDBTypes, false, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}
	if err = randomize.Struct(seed, repositoryTwo, repositoryDBTypes, false, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = repositoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = repositoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func repositoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func repositoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func repositoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func repositoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func repositoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func repositoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func repositoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func repositoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func repositoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Repository) error {
	*o = Repository{}
	return nil
}

func testRepositoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Repository{}
	o := &Repository{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, repositoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Repository object: %s", err)
	}

	AddRepositoryHook(boil.BeforeInsertHook, repositoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	repositoryBeforeInsertHooks = []RepositoryHook{}

	AddRepositoryHook(boil.AfterInsertHook, repositoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	repositoryAfterInsertHooks = []RepositoryHook{}

	AddRepositoryHook(boil.AfterSelectHook, repositoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	repositoryAfterSelectHooks = []RepositoryHook{}

	AddRepositoryHook(boil.BeforeUpdateHook, repositoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	repositoryBeforeUpdateHooks = []RepositoryHook{}

	AddRepositoryHook(boil.AfterUpdateHook, repositoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	repositoryAfterUpdateHooks = []RepositoryHook{}

	AddRepositoryHook(boil.BeforeDeleteHook, repositoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	repositoryBeforeDeleteHooks = []RepositoryHook{}

	AddRepositoryHook(boil.AfterDeleteHook, repositoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	repositoryAfterDeleteHooks = []RepositoryHook{}

	AddRepositoryHook(boil.BeforeUpsertHook, repositoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	repositoryBeforeUpsertHooks = []RepositoryHook{}

	AddRepositoryHook(boil.AfterUpsertHook, repositoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	repositoryAfterUpsertHooks = []RepositoryHook{}
}

func testRepositoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRepositoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(repositoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRepositoryToManyIssues(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Repository
	var b, c Issue

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, issueDBTypes, false, issueColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, issueDBTypes, false, issueColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.Repository = a.ID
	c.Repository = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Issues().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.Repository == b.Repository {
			bFound = true
		}
		if v.Repository == c.Repository {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RepositorySlice{&a}
	if err = a.L.LoadIssues(ctx, tx, false, (*[]*Repository)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Issues); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Issues = nil
	if err = a.L.LoadIssues(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Issues); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testRepositoryToManyPullrequests(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Repository
	var b, c Pullrequest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, pullrequestDBTypes, false, pullrequestColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pullrequestDBTypes, false, pullrequestColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.Repository = a.ID
	c.Repository = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Pullrequests().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.Repository == b.Repository {
			bFound = true
		}
		if v.Repository == c.Repository {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RepositorySlice{&a}
	if err = a.L.LoadPullrequests(ctx, tx, false, (*[]*Repository)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Pullrequests); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Pullrequests = nil
	if err = a.L.LoadPullrequests(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Pullrequests); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testRepositoryToManyAddOpIssues(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Repository
	var b, c, d, e Issue

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, repositoryDBTypes, false, strmangle.SetComplement(repositoryPrimaryKeyColumns, repositoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Issue{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Issue{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddIssues(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.Repository {
			t.Error("foreign key was wrong value", a.ID, first.Repository)
		}
		if a.ID != second.Repository {
			t.Error("foreign key was wrong value", a.ID, second.Repository)
		}

		if first.R.IssueRepository != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.IssueRepository != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Issues[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Issues[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Issues().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testRepositoryToManyAddOpPullrequests(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Repository
	var b, c, d, e Pullrequest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, repositoryDBTypes, false, strmangle.SetComplement(repositoryPrimaryKeyColumns, repositoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Pullrequest{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pullrequestDBTypes, false, strmangle.SetComplement(pullrequestPrimaryKeyColumns, pullrequestColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Pullrequest{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPullrequests(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.Repository {
			t.Error("foreign key was wrong value", a.ID, first.Repository)
		}
		if a.ID != second.Repository {
			t.Error("foreign key was wrong value", a.ID, second.Repository)
		}

		if first.R.PullrequestRepository != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.PullrequestRepository != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Pullrequests[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Pullrequests[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Pullrequests().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testRepositoryToOneUserUsingOwnerUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Repository
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, repositoryDBTypes, false, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.Owner = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.OwnerUser().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := RepositorySlice{&local}
	if err = local.L.LoadOwnerUser(ctx, tx, false, (*[]*Repository)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OwnerUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OwnerUser = nil
	if err = local.L.LoadOwnerUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OwnerUser == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testRepositoryToOneSetOpUserUsingOwnerUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Repository
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, repositoryDBTypes, false, strmangle.SetComplement(repositoryPrimaryKeyColumns, repositoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetOwnerUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OwnerUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OwnerRepositories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Owner != x.ID {
			t.Error("foreign key was wrong value", a.Owner)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Owner))
		reflect.Indirect(reflect.ValueOf(&a.Owner)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Owner != x.ID {
			t.Error("foreign key was wrong value", a.Owner, x.ID)
		}
	}
}

func testRepositoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
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

func testRepositoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RepositorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRepositoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Repositories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	repositoryDBTypes = map[string]string{`ID`: `text`, `Owner`: `text`, `Name`: `text`, `CreatedAt`: `timestamp with time zone`}
	_                 = bytes.MinRead
)

func testRepositoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(repositoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(repositoryAllColumns) == len(repositoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRepositoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(repositoryAllColumns) == len(repositoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Repository{}
	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, repositoryDBTypes, true, repositoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(repositoryAllColumns, repositoryPrimaryKeyColumns) {
		fields = repositoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			repositoryAllColumns,
			repositoryPrimaryKeyColumns,
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

	slice := RepositorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRepositoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(repositoryAllColumns) == len(repositoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Repository{}
	if err = randomize.Struct(seed, &o, repositoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Repository: %s", err)
	}

	count, err := Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, repositoryDBTypes, false, repositoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Repository struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Repository: %s", err)
	}

	count, err = Repositories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
