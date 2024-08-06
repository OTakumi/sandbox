// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Issues", testIssues)
	t.Run("Projectcards", testProjectcards)
	t.Run("Projects", testProjects)
	t.Run("Pullrequests", testPullrequests)
	t.Run("Repositories", testRepositories)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Issues", testIssuesDelete)
	t.Run("Projectcards", testProjectcardsDelete)
	t.Run("Projects", testProjectsDelete)
	t.Run("Pullrequests", testPullrequestsDelete)
	t.Run("Repositories", testRepositoriesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Issues", testIssuesQueryDeleteAll)
	t.Run("Projectcards", testProjectcardsQueryDeleteAll)
	t.Run("Projects", testProjectsQueryDeleteAll)
	t.Run("Pullrequests", testPullrequestsQueryDeleteAll)
	t.Run("Repositories", testRepositoriesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Issues", testIssuesSliceDeleteAll)
	t.Run("Projectcards", testProjectcardsSliceDeleteAll)
	t.Run("Projects", testProjectsSliceDeleteAll)
	t.Run("Pullrequests", testPullrequestsSliceDeleteAll)
	t.Run("Repositories", testRepositoriesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Issues", testIssuesExists)
	t.Run("Projectcards", testProjectcardsExists)
	t.Run("Projects", testProjectsExists)
	t.Run("Pullrequests", testPullrequestsExists)
	t.Run("Repositories", testRepositoriesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Issues", testIssuesFind)
	t.Run("Projectcards", testProjectcardsFind)
	t.Run("Projects", testProjectsFind)
	t.Run("Pullrequests", testPullrequestsFind)
	t.Run("Repositories", testRepositoriesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Issues", testIssuesBind)
	t.Run("Projectcards", testProjectcardsBind)
	t.Run("Projects", testProjectsBind)
	t.Run("Pullrequests", testPullrequestsBind)
	t.Run("Repositories", testRepositoriesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Issues", testIssuesOne)
	t.Run("Projectcards", testProjectcardsOne)
	t.Run("Projects", testProjectsOne)
	t.Run("Pullrequests", testPullrequestsOne)
	t.Run("Repositories", testRepositoriesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Issues", testIssuesAll)
	t.Run("Projectcards", testProjectcardsAll)
	t.Run("Projects", testProjectsAll)
	t.Run("Pullrequests", testPullrequestsAll)
	t.Run("Repositories", testRepositoriesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Issues", testIssuesCount)
	t.Run("Projectcards", testProjectcardsCount)
	t.Run("Projects", testProjectsCount)
	t.Run("Pullrequests", testPullrequestsCount)
	t.Run("Repositories", testRepositoriesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Issues", testIssuesHooks)
	t.Run("Projectcards", testProjectcardsHooks)
	t.Run("Projects", testProjectsHooks)
	t.Run("Pullrequests", testPullrequestsHooks)
	t.Run("Repositories", testRepositoriesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Issues", testIssuesInsert)
	t.Run("Issues", testIssuesInsertWhitelist)
	t.Run("Projectcards", testProjectcardsInsert)
	t.Run("Projectcards", testProjectcardsInsertWhitelist)
	t.Run("Projects", testProjectsInsert)
	t.Run("Projects", testProjectsInsertWhitelist)
	t.Run("Pullrequests", testPullrequestsInsert)
	t.Run("Pullrequests", testPullrequestsInsertWhitelist)
	t.Run("Repositories", testRepositoriesInsert)
	t.Run("Repositories", testRepositoriesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Issues", testIssuesReload)
	t.Run("Projectcards", testProjectcardsReload)
	t.Run("Projects", testProjectsReload)
	t.Run("Pullrequests", testPullrequestsReload)
	t.Run("Repositories", testRepositoriesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Issues", testIssuesReloadAll)
	t.Run("Projectcards", testProjectcardsReloadAll)
	t.Run("Projects", testProjectsReloadAll)
	t.Run("Pullrequests", testPullrequestsReloadAll)
	t.Run("Repositories", testRepositoriesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Issues", testIssuesSelect)
	t.Run("Projectcards", testProjectcardsSelect)
	t.Run("Projects", testProjectsSelect)
	t.Run("Pullrequests", testPullrequestsSelect)
	t.Run("Repositories", testRepositoriesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Issues", testIssuesUpdate)
	t.Run("Projectcards", testProjectcardsUpdate)
	t.Run("Projects", testProjectsUpdate)
	t.Run("Pullrequests", testPullrequestsUpdate)
	t.Run("Repositories", testRepositoriesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Issues", testIssuesSliceUpdateAll)
	t.Run("Projectcards", testProjectcardsSliceUpdateAll)
	t.Run("Projects", testProjectsSliceUpdateAll)
	t.Run("Pullrequests", testPullrequestsSliceUpdateAll)
	t.Run("Repositories", testRepositoriesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}