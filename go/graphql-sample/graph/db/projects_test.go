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

func testProjects(t *testing.T) {
	t.Parallel()

	query := Projects()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProjectsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
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

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProjectsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Projects().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProjectsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProjectSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProjectsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProjectExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Project exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProjectExists to return true, but got false.")
	}
}

func testProjectsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	projectFound, err := FindProject(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if projectFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProjectsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Projects().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProjectsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Projects().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProjectsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	projectOne := &Project{}
	projectTwo := &Project{}
	if err = randomize.Struct(seed, projectOne, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}
	if err = randomize.Struct(seed, projectTwo, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = projectOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = projectTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Projects().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProjectsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	projectOne := &Project{}
	projectTwo := &Project{}
	if err = randomize.Struct(seed, projectOne, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}
	if err = randomize.Struct(seed, projectTwo, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = projectOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = projectTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func projectBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func projectBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func projectBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func projectBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Project) error {
	*o = Project{}
	return nil
}

func testProjectsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Project{}
	o := &Project{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, projectDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Project object: %s", err)
	}

	AddProjectHook(boil.BeforeInsertHook, projectBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	projectBeforeInsertHooks = []ProjectHook{}

	AddProjectHook(boil.AfterInsertHook, projectAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	projectAfterInsertHooks = []ProjectHook{}

	AddProjectHook(boil.AfterSelectHook, projectAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	projectAfterSelectHooks = []ProjectHook{}

	AddProjectHook(boil.BeforeUpdateHook, projectBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	projectBeforeUpdateHooks = []ProjectHook{}

	AddProjectHook(boil.AfterUpdateHook, projectAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	projectAfterUpdateHooks = []ProjectHook{}

	AddProjectHook(boil.BeforeDeleteHook, projectBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	projectBeforeDeleteHooks = []ProjectHook{}

	AddProjectHook(boil.AfterDeleteHook, projectAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	projectAfterDeleteHooks = []ProjectHook{}

	AddProjectHook(boil.BeforeUpsertHook, projectBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	projectBeforeUpsertHooks = []ProjectHook{}

	AddProjectHook(boil.AfterUpsertHook, projectAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	projectAfterUpsertHooks = []ProjectHook{}
}

func testProjectsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProjectsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(projectColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProjectToManyProjectcards(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Project
	var b, c Projectcard

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, projectcardDBTypes, false, projectcardColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, projectcardDBTypes, false, projectcardColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.Project = a.ID
	c.Project = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Projectcards().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.Project == b.Project {
			bFound = true
		}
		if v.Project == c.Project {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProjectSlice{&a}
	if err = a.L.LoadProjectcards(ctx, tx, false, (*[]*Project)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Projectcards); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Projectcards = nil
	if err = a.L.LoadProjectcards(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Projectcards); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProjectToManyAddOpProjectcards(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Project
	var b, c, d, e Projectcard

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Projectcard{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, projectcardDBTypes, false, strmangle.SetComplement(projectcardPrimaryKeyColumns, projectcardColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Projectcard{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddProjectcards(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.Project {
			t.Error("foreign key was wrong value", a.ID, first.Project)
		}
		if a.ID != second.Project {
			t.Error("foreign key was wrong value", a.ID, second.Project)
		}

		if first.R.ProjectcardProject != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ProjectcardProject != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Projectcards[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Projectcards[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Projectcards().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testProjectToOneUserUsingOwnerUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Project
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
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

	slice := ProjectSlice{&local}
	if err = local.L.LoadOwnerUser(ctx, tx, false, (*[]*Project)(&slice), nil); err != nil {
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

func testProjectToOneSetOpUserUsingOwnerUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Project
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
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

		if x.R.OwnerProjects[0] != &a {
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

func testProjectsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
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

func testProjectsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProjectSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProjectsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Projects().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	projectDBTypes = map[string]string{`ID`: `text`, `Title`: `text`, `URL`: `text`, `Owner`: `text`}
	_              = bytes.MinRead
)

func testProjectsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(projectPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(projectAllColumns) == len(projectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, projectDBTypes, true, projectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProjectsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(projectAllColumns) == len(projectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Project{}
	if err = randomize.Struct(seed, o, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, projectDBTypes, true, projectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(projectAllColumns, projectPrimaryKeyColumns) {
		fields = projectAllColumns
	} else {
		fields = strmangle.SetComplement(
			projectAllColumns,
			projectPrimaryKeyColumns,
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

	slice := ProjectSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProjectsUpsert(t *testing.T) {
	t.Parallel()

	if len(projectAllColumns) == len(projectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Project{}
	if err = randomize.Struct(seed, &o, projectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Project: %s", err)
	}

	count, err := Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, projectDBTypes, false, projectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Project: %s", err)
	}

	count, err = Projects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
