// Code generated by ent, DO NOT EDIT.

package pullrequests

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldID, id))
}

// PrID applies equality check predicate on the "pr_id" field. It's identical to PrIDEQ.
func PrID(v uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldPrID, v))
}

// RepositoryName applies equality check predicate on the "repository_name" field. It's identical to RepositoryNameEQ.
func RepositoryName(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldRepositoryName, v))
}

// RepositoryOrganization applies equality check predicate on the "repository_organization" field. It's identical to RepositoryOrganizationEQ.
func RepositoryOrganization(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldRepositoryOrganization, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldNumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldCreatedAt, v))
}

// ClosedAt applies equality check predicate on the "closed_at" field. It's identical to ClosedAtEQ.
func ClosedAt(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldClosedAt, v))
}

// MergedAt applies equality check predicate on the "merged_at" field. It's identical to MergedAtEQ.
func MergedAt(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldMergedAt, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldState, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldAuthor, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldTitle, v))
}

// MergeCommit applies equality check predicate on the "merge_commit" field. It's identical to MergeCommitEQ.
func MergeCommit(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldMergeCommit, v))
}

// RetestCount applies equality check predicate on the "retest_count" field. It's identical to RetestCountEQ.
func RetestCount(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldRetestCount, v))
}

// RetestBeforeMergeCount applies equality check predicate on the "retest_before_merge_count" field. It's identical to RetestBeforeMergeCountEQ.
func RetestBeforeMergeCount(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldRetestBeforeMergeCount, v))
}

// PrIDEQ applies the EQ predicate on the "pr_id" field.
func PrIDEQ(v uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldPrID, v))
}

// PrIDNEQ applies the NEQ predicate on the "pr_id" field.
func PrIDNEQ(v uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldPrID, v))
}

// PrIDIn applies the In predicate on the "pr_id" field.
func PrIDIn(vs ...uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldPrID, vs...))
}

// PrIDNotIn applies the NotIn predicate on the "pr_id" field.
func PrIDNotIn(vs ...uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldPrID, vs...))
}

// PrIDGT applies the GT predicate on the "pr_id" field.
func PrIDGT(v uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldPrID, v))
}

// PrIDGTE applies the GTE predicate on the "pr_id" field.
func PrIDGTE(v uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldPrID, v))
}

// PrIDLT applies the LT predicate on the "pr_id" field.
func PrIDLT(v uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldPrID, v))
}

// PrIDLTE applies the LTE predicate on the "pr_id" field.
func PrIDLTE(v uuid.UUID) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldPrID, v))
}

// RepositoryNameEQ applies the EQ predicate on the "repository_name" field.
func RepositoryNameEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldRepositoryName, v))
}

// RepositoryNameNEQ applies the NEQ predicate on the "repository_name" field.
func RepositoryNameNEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldRepositoryName, v))
}

// RepositoryNameIn applies the In predicate on the "repository_name" field.
func RepositoryNameIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldRepositoryName, vs...))
}

// RepositoryNameNotIn applies the NotIn predicate on the "repository_name" field.
func RepositoryNameNotIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldRepositoryName, vs...))
}

// RepositoryNameGT applies the GT predicate on the "repository_name" field.
func RepositoryNameGT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldRepositoryName, v))
}

// RepositoryNameGTE applies the GTE predicate on the "repository_name" field.
func RepositoryNameGTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldRepositoryName, v))
}

// RepositoryNameLT applies the LT predicate on the "repository_name" field.
func RepositoryNameLT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldRepositoryName, v))
}

// RepositoryNameLTE applies the LTE predicate on the "repository_name" field.
func RepositoryNameLTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldRepositoryName, v))
}

// RepositoryNameContains applies the Contains predicate on the "repository_name" field.
func RepositoryNameContains(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContains(FieldRepositoryName, v))
}

// RepositoryNameHasPrefix applies the HasPrefix predicate on the "repository_name" field.
func RepositoryNameHasPrefix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasPrefix(FieldRepositoryName, v))
}

// RepositoryNameHasSuffix applies the HasSuffix predicate on the "repository_name" field.
func RepositoryNameHasSuffix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasSuffix(FieldRepositoryName, v))
}

// RepositoryNameEqualFold applies the EqualFold predicate on the "repository_name" field.
func RepositoryNameEqualFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEqualFold(FieldRepositoryName, v))
}

// RepositoryNameContainsFold applies the ContainsFold predicate on the "repository_name" field.
func RepositoryNameContainsFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContainsFold(FieldRepositoryName, v))
}

// RepositoryOrganizationEQ applies the EQ predicate on the "repository_organization" field.
func RepositoryOrganizationEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationNEQ applies the NEQ predicate on the "repository_organization" field.
func RepositoryOrganizationNEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationIn applies the In predicate on the "repository_organization" field.
func RepositoryOrganizationIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldRepositoryOrganization, vs...))
}

// RepositoryOrganizationNotIn applies the NotIn predicate on the "repository_organization" field.
func RepositoryOrganizationNotIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldRepositoryOrganization, vs...))
}

// RepositoryOrganizationGT applies the GT predicate on the "repository_organization" field.
func RepositoryOrganizationGT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationGTE applies the GTE predicate on the "repository_organization" field.
func RepositoryOrganizationGTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationLT applies the LT predicate on the "repository_organization" field.
func RepositoryOrganizationLT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationLTE applies the LTE predicate on the "repository_organization" field.
func RepositoryOrganizationLTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationContains applies the Contains predicate on the "repository_organization" field.
func RepositoryOrganizationContains(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContains(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationHasPrefix applies the HasPrefix predicate on the "repository_organization" field.
func RepositoryOrganizationHasPrefix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasPrefix(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationHasSuffix applies the HasSuffix predicate on the "repository_organization" field.
func RepositoryOrganizationHasSuffix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasSuffix(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationEqualFold applies the EqualFold predicate on the "repository_organization" field.
func RepositoryOrganizationEqualFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEqualFold(FieldRepositoryOrganization, v))
}

// RepositoryOrganizationContainsFold applies the ContainsFold predicate on the "repository_organization" field.
func RepositoryOrganizationContainsFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContainsFold(FieldRepositoryOrganization, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldCreatedAt, v))
}

// ClosedAtEQ applies the EQ predicate on the "closed_at" field.
func ClosedAtEQ(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldClosedAt, v))
}

// ClosedAtNEQ applies the NEQ predicate on the "closed_at" field.
func ClosedAtNEQ(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldClosedAt, v))
}

// ClosedAtIn applies the In predicate on the "closed_at" field.
func ClosedAtIn(vs ...time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldClosedAt, vs...))
}

// ClosedAtNotIn applies the NotIn predicate on the "closed_at" field.
func ClosedAtNotIn(vs ...time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldClosedAt, vs...))
}

// ClosedAtGT applies the GT predicate on the "closed_at" field.
func ClosedAtGT(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldClosedAt, v))
}

// ClosedAtGTE applies the GTE predicate on the "closed_at" field.
func ClosedAtGTE(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldClosedAt, v))
}

// ClosedAtLT applies the LT predicate on the "closed_at" field.
func ClosedAtLT(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldClosedAt, v))
}

// ClosedAtLTE applies the LTE predicate on the "closed_at" field.
func ClosedAtLTE(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldClosedAt, v))
}

// MergedAtEQ applies the EQ predicate on the "merged_at" field.
func MergedAtEQ(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldMergedAt, v))
}

// MergedAtNEQ applies the NEQ predicate on the "merged_at" field.
func MergedAtNEQ(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldMergedAt, v))
}

// MergedAtIn applies the In predicate on the "merged_at" field.
func MergedAtIn(vs ...time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldMergedAt, vs...))
}

// MergedAtNotIn applies the NotIn predicate on the "merged_at" field.
func MergedAtNotIn(vs ...time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldMergedAt, vs...))
}

// MergedAtGT applies the GT predicate on the "merged_at" field.
func MergedAtGT(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldMergedAt, v))
}

// MergedAtGTE applies the GTE predicate on the "merged_at" field.
func MergedAtGTE(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldMergedAt, v))
}

// MergedAtLT applies the LT predicate on the "merged_at" field.
func MergedAtLT(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldMergedAt, v))
}

// MergedAtLTE applies the LTE predicate on the "merged_at" field.
func MergedAtLTE(v time.Time) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldMergedAt, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasSuffix(FieldState, v))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContainsFold(FieldState, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContainsFold(FieldAuthor, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContainsFold(FieldTitle, v))
}

// MergeCommitEQ applies the EQ predicate on the "merge_commit" field.
func MergeCommitEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldMergeCommit, v))
}

// MergeCommitNEQ applies the NEQ predicate on the "merge_commit" field.
func MergeCommitNEQ(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldMergeCommit, v))
}

// MergeCommitIn applies the In predicate on the "merge_commit" field.
func MergeCommitIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldMergeCommit, vs...))
}

// MergeCommitNotIn applies the NotIn predicate on the "merge_commit" field.
func MergeCommitNotIn(vs ...string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldMergeCommit, vs...))
}

// MergeCommitGT applies the GT predicate on the "merge_commit" field.
func MergeCommitGT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldMergeCommit, v))
}

// MergeCommitGTE applies the GTE predicate on the "merge_commit" field.
func MergeCommitGTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldMergeCommit, v))
}

// MergeCommitLT applies the LT predicate on the "merge_commit" field.
func MergeCommitLT(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldMergeCommit, v))
}

// MergeCommitLTE applies the LTE predicate on the "merge_commit" field.
func MergeCommitLTE(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldMergeCommit, v))
}

// MergeCommitContains applies the Contains predicate on the "merge_commit" field.
func MergeCommitContains(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContains(FieldMergeCommit, v))
}

// MergeCommitHasPrefix applies the HasPrefix predicate on the "merge_commit" field.
func MergeCommitHasPrefix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasPrefix(FieldMergeCommit, v))
}

// MergeCommitHasSuffix applies the HasSuffix predicate on the "merge_commit" field.
func MergeCommitHasSuffix(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldHasSuffix(FieldMergeCommit, v))
}

// MergeCommitIsNil applies the IsNil predicate on the "merge_commit" field.
func MergeCommitIsNil() predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIsNull(FieldMergeCommit))
}

// MergeCommitNotNil applies the NotNil predicate on the "merge_commit" field.
func MergeCommitNotNil() predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotNull(FieldMergeCommit))
}

// MergeCommitEqualFold applies the EqualFold predicate on the "merge_commit" field.
func MergeCommitEqualFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEqualFold(FieldMergeCommit, v))
}

// MergeCommitContainsFold applies the ContainsFold predicate on the "merge_commit" field.
func MergeCommitContainsFold(v string) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldContainsFold(FieldMergeCommit, v))
}

// RetestCountEQ applies the EQ predicate on the "retest_count" field.
func RetestCountEQ(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldRetestCount, v))
}

// RetestCountNEQ applies the NEQ predicate on the "retest_count" field.
func RetestCountNEQ(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldRetestCount, v))
}

// RetestCountIn applies the In predicate on the "retest_count" field.
func RetestCountIn(vs ...float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldRetestCount, vs...))
}

// RetestCountNotIn applies the NotIn predicate on the "retest_count" field.
func RetestCountNotIn(vs ...float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldRetestCount, vs...))
}

// RetestCountGT applies the GT predicate on the "retest_count" field.
func RetestCountGT(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldRetestCount, v))
}

// RetestCountGTE applies the GTE predicate on the "retest_count" field.
func RetestCountGTE(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldRetestCount, v))
}

// RetestCountLT applies the LT predicate on the "retest_count" field.
func RetestCountLT(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldRetestCount, v))
}

// RetestCountLTE applies the LTE predicate on the "retest_count" field.
func RetestCountLTE(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldRetestCount, v))
}

// RetestCountIsNil applies the IsNil predicate on the "retest_count" field.
func RetestCountIsNil() predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIsNull(FieldRetestCount))
}

// RetestCountNotNil applies the NotNil predicate on the "retest_count" field.
func RetestCountNotNil() predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotNull(FieldRetestCount))
}

// RetestBeforeMergeCountEQ applies the EQ predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountEQ(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldEQ(FieldRetestBeforeMergeCount, v))
}

// RetestBeforeMergeCountNEQ applies the NEQ predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountNEQ(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNEQ(FieldRetestBeforeMergeCount, v))
}

// RetestBeforeMergeCountIn applies the In predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountIn(vs ...float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIn(FieldRetestBeforeMergeCount, vs...))
}

// RetestBeforeMergeCountNotIn applies the NotIn predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountNotIn(vs ...float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotIn(FieldRetestBeforeMergeCount, vs...))
}

// RetestBeforeMergeCountGT applies the GT predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountGT(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGT(FieldRetestBeforeMergeCount, v))
}

// RetestBeforeMergeCountGTE applies the GTE predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountGTE(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldGTE(FieldRetestBeforeMergeCount, v))
}

// RetestBeforeMergeCountLT applies the LT predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountLT(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLT(FieldRetestBeforeMergeCount, v))
}

// RetestBeforeMergeCountLTE applies the LTE predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountLTE(v float64) predicate.PullRequests {
	return predicate.PullRequests(sql.FieldLTE(FieldRetestBeforeMergeCount, v))
}

// RetestBeforeMergeCountIsNil applies the IsNil predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountIsNil() predicate.PullRequests {
	return predicate.PullRequests(sql.FieldIsNull(FieldRetestBeforeMergeCount))
}

// RetestBeforeMergeCountNotNil applies the NotNil predicate on the "retest_before_merge_count" field.
func RetestBeforeMergeCountNotNil() predicate.PullRequests {
	return predicate.PullRequests(sql.FieldNotNull(FieldRetestBeforeMergeCount))
}

// HasPrs applies the HasEdge predicate on the "prs" edge.
func HasPrs() predicate.PullRequests {
	return predicate.PullRequests(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrsTable, PrsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrsWith applies the HasEdge predicate on the "prs" edge with a given conditions (other predicates).
func HasPrsWith(preds ...predicate.Repository) predicate.PullRequests {
	return predicate.PullRequests(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrsTable, PrsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PullRequests) predicate.PullRequests {
	return predicate.PullRequests(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PullRequests) predicate.PullRequests {
	return predicate.PullRequests(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PullRequests) predicate.PullRequests {
	return predicate.PullRequests(func(s *sql.Selector) {
		p(s.Not())
	})
}
