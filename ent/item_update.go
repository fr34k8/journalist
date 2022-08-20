// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/item"
	"github.com/mrusme/journalist/ent/predicate"
	"github.com/mrusme/journalist/ent/read"
	"github.com/mrusme/journalist/ent/user"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetItemGUID sets the "item_guid" field.
func (iu *ItemUpdate) SetItemGUID(s string) *ItemUpdate {
	iu.mutation.SetItemGUID(s)
	return iu
}

// SetItemTitle sets the "item_title" field.
func (iu *ItemUpdate) SetItemTitle(s string) *ItemUpdate {
	iu.mutation.SetItemTitle(s)
	return iu
}

// SetItemDescription sets the "item_description" field.
func (iu *ItemUpdate) SetItemDescription(s string) *ItemUpdate {
	iu.mutation.SetItemDescription(s)
	return iu
}

// SetItemContent sets the "item_content" field.
func (iu *ItemUpdate) SetItemContent(s string) *ItemUpdate {
	iu.mutation.SetItemContent(s)
	return iu
}

// SetItemLink sets the "item_link" field.
func (iu *ItemUpdate) SetItemLink(s string) *ItemUpdate {
	iu.mutation.SetItemLink(s)
	return iu
}

// SetItemUpdated sets the "item_updated" field.
func (iu *ItemUpdate) SetItemUpdated(s string) *ItemUpdate {
	iu.mutation.SetItemUpdated(s)
	return iu
}

// SetItemPublished sets the "item_published" field.
func (iu *ItemUpdate) SetItemPublished(s string) *ItemUpdate {
	iu.mutation.SetItemPublished(s)
	return iu
}

// SetItemAuthorName sets the "item_author_name" field.
func (iu *ItemUpdate) SetItemAuthorName(s string) *ItemUpdate {
	iu.mutation.SetItemAuthorName(s)
	return iu
}

// SetItemAuthorEmail sets the "item_author_email" field.
func (iu *ItemUpdate) SetItemAuthorEmail(s string) *ItemUpdate {
	iu.mutation.SetItemAuthorEmail(s)
	return iu
}

// SetItemImageTitle sets the "item_image_title" field.
func (iu *ItemUpdate) SetItemImageTitle(s string) *ItemUpdate {
	iu.mutation.SetItemImageTitle(s)
	return iu
}

// SetItemImageURL sets the "item_image_url" field.
func (iu *ItemUpdate) SetItemImageURL(s string) *ItemUpdate {
	iu.mutation.SetItemImageURL(s)
	return iu
}

// SetItemCategories sets the "item_categories" field.
func (iu *ItemUpdate) SetItemCategories(s string) *ItemUpdate {
	iu.mutation.SetItemCategories(s)
	return iu
}

// SetItemEnclosures sets the "item_enclosures" field.
func (iu *ItemUpdate) SetItemEnclosures(s string) *ItemUpdate {
	iu.mutation.SetItemEnclosures(s)
	return iu
}

// SetCrawlerTitle sets the "crawler_title" field.
func (iu *ItemUpdate) SetCrawlerTitle(s string) *ItemUpdate {
	iu.mutation.SetCrawlerTitle(s)
	return iu
}

// SetCrawlerAuthor sets the "crawler_author" field.
func (iu *ItemUpdate) SetCrawlerAuthor(s string) *ItemUpdate {
	iu.mutation.SetCrawlerAuthor(s)
	return iu
}

// SetCrawlerExcerpt sets the "crawler_excerpt" field.
func (iu *ItemUpdate) SetCrawlerExcerpt(s string) *ItemUpdate {
	iu.mutation.SetCrawlerExcerpt(s)
	return iu
}

// SetCrawlerSiteName sets the "crawler_site_name" field.
func (iu *ItemUpdate) SetCrawlerSiteName(s string) *ItemUpdate {
	iu.mutation.SetCrawlerSiteName(s)
	return iu
}

// SetCrawlerImage sets the "crawler_image" field.
func (iu *ItemUpdate) SetCrawlerImage(s string) *ItemUpdate {
	iu.mutation.SetCrawlerImage(s)
	return iu
}

// SetCrawlerContentHTML sets the "crawler_content_html" field.
func (iu *ItemUpdate) SetCrawlerContentHTML(s string) *ItemUpdate {
	iu.mutation.SetCrawlerContentHTML(s)
	return iu
}

// SetCrawlerContentText sets the "crawler_content_text" field.
func (iu *ItemUpdate) SetCrawlerContentText(s string) *ItemUpdate {
	iu.mutation.SetCrawlerContentText(s)
	return iu
}

// SetCreatedAt sets the "created_at" field.
func (iu *ItemUpdate) SetCreatedAt(t time.Time) *ItemUpdate {
	iu.mutation.SetCreatedAt(t)
	return iu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableCreatedAt(t *time.Time) *ItemUpdate {
	if t != nil {
		iu.SetCreatedAt(*t)
	}
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *ItemUpdate) SetUpdatedAt(t time.Time) *ItemUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (iu *ItemUpdate) SetFeedID(id uuid.UUID) *ItemUpdate {
	iu.mutation.SetFeedID(id)
	return iu
}

// SetNillableFeedID sets the "feed" edge to the Feed entity by ID if the given value is not nil.
func (iu *ItemUpdate) SetNillableFeedID(id *uuid.UUID) *ItemUpdate {
	if id != nil {
		iu = iu.SetFeedID(*id)
	}
	return iu
}

// SetFeed sets the "feed" edge to the Feed entity.
func (iu *ItemUpdate) SetFeed(f *Feed) *ItemUpdate {
	return iu.SetFeedID(f.ID)
}

// AddReadByUserIDs adds the "read_by_users" edge to the User entity by IDs.
func (iu *ItemUpdate) AddReadByUserIDs(ids ...uuid.UUID) *ItemUpdate {
	iu.mutation.AddReadByUserIDs(ids...)
	return iu
}

// AddReadByUsers adds the "read_by_users" edges to the User entity.
func (iu *ItemUpdate) AddReadByUsers(u ...*User) *ItemUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return iu.AddReadByUserIDs(ids...)
}

// AddReadIDs adds the "reads" edge to the Read entity by IDs.
func (iu *ItemUpdate) AddReadIDs(ids ...uuid.UUID) *ItemUpdate {
	iu.mutation.AddReadIDs(ids...)
	return iu
}

// AddReads adds the "reads" edges to the Read entity.
func (iu *ItemUpdate) AddReads(r ...*Read) *ItemUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return iu.AddReadIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearFeed clears the "feed" edge to the Feed entity.
func (iu *ItemUpdate) ClearFeed() *ItemUpdate {
	iu.mutation.ClearFeed()
	return iu
}

// ClearReadByUsers clears all "read_by_users" edges to the User entity.
func (iu *ItemUpdate) ClearReadByUsers() *ItemUpdate {
	iu.mutation.ClearReadByUsers()
	return iu
}

// RemoveReadByUserIDs removes the "read_by_users" edge to User entities by IDs.
func (iu *ItemUpdate) RemoveReadByUserIDs(ids ...uuid.UUID) *ItemUpdate {
	iu.mutation.RemoveReadByUserIDs(ids...)
	return iu
}

// RemoveReadByUsers removes "read_by_users" edges to User entities.
func (iu *ItemUpdate) RemoveReadByUsers(u ...*User) *ItemUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return iu.RemoveReadByUserIDs(ids...)
}

// ClearReads clears all "reads" edges to the Read entity.
func (iu *ItemUpdate) ClearReads() *ItemUpdate {
	iu.mutation.ClearReads()
	return iu
}

// RemoveReadIDs removes the "reads" edge to Read entities by IDs.
func (iu *ItemUpdate) RemoveReadIDs(ids ...uuid.UUID) *ItemUpdate {
	iu.mutation.RemoveReadIDs(ids...)
	return iu
}

// RemoveReads removes "reads" edges to Read entities.
func (iu *ItemUpdate) RemoveReads(r ...*Read) *ItemUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return iu.RemoveReadIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ItemUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := item.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if v, ok := iu.mutation.ItemLink(); ok {
		if err := item.ItemLinkValidator(v); err != nil {
			return &ValidationError{Name: "item_link", err: fmt.Errorf(`ent: validator failed for field "Item.item_link": %w`, err)}
		}
	}
	return nil
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.ItemGUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemGUID,
		})
	}
	if value, ok := iu.mutation.ItemTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemTitle,
		})
	}
	if value, ok := iu.mutation.ItemDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemDescription,
		})
	}
	if value, ok := iu.mutation.ItemContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemContent,
		})
	}
	if value, ok := iu.mutation.ItemLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemLink,
		})
	}
	if value, ok := iu.mutation.ItemUpdated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemUpdated,
		})
	}
	if value, ok := iu.mutation.ItemPublished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemPublished,
		})
	}
	if value, ok := iu.mutation.ItemAuthorName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemAuthorName,
		})
	}
	if value, ok := iu.mutation.ItemAuthorEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemAuthorEmail,
		})
	}
	if value, ok := iu.mutation.ItemImageTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemImageTitle,
		})
	}
	if value, ok := iu.mutation.ItemImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemImageURL,
		})
	}
	if value, ok := iu.mutation.ItemCategories(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemCategories,
		})
	}
	if value, ok := iu.mutation.ItemEnclosures(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemEnclosures,
		})
	}
	if value, ok := iu.mutation.CrawlerTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerTitle,
		})
	}
	if value, ok := iu.mutation.CrawlerAuthor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerAuthor,
		})
	}
	if value, ok := iu.mutation.CrawlerExcerpt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerExcerpt,
		})
	}
	if value, ok := iu.mutation.CrawlerSiteName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerSiteName,
		})
	}
	if value, ok := iu.mutation.CrawlerImage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerImage,
		})
	}
	if value, ok := iu.mutation.CrawlerContentHTML(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerContentHTML,
		})
	}
	if value, ok := iu.mutation.CrawlerContentText(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerContentText,
		})
	}
	if value, ok := iu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldCreatedAt,
		})
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdatedAt,
		})
	}
	if iu.mutation.FeedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.FeedTable,
			Columns: []string{item.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.FeedTable,
			Columns: []string{item.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.ReadByUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.ReadByUsersTable,
			Columns: item.ReadByUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		createE := &ReadCreate{config: iu.config, mutation: newReadMutation(iu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedReadByUsersIDs(); len(nodes) > 0 && !iu.mutation.ReadByUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.ReadByUsersTable,
			Columns: item.ReadByUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReadCreate{config: iu.config, mutation: newReadMutation(iu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ReadByUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.ReadByUsersTable,
			Columns: item.ReadByUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReadCreate{config: iu.config, mutation: newReadMutation(iu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.ReadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ReadsTable,
			Columns: []string{item.ReadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: read.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedReadsIDs(); len(nodes) > 0 && !iu.mutation.ReadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ReadsTable,
			Columns: []string{item.ReadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: read.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ReadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ReadsTable,
			Columns: []string{item.ReadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: read.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetItemGUID sets the "item_guid" field.
func (iuo *ItemUpdateOne) SetItemGUID(s string) *ItemUpdateOne {
	iuo.mutation.SetItemGUID(s)
	return iuo
}

// SetItemTitle sets the "item_title" field.
func (iuo *ItemUpdateOne) SetItemTitle(s string) *ItemUpdateOne {
	iuo.mutation.SetItemTitle(s)
	return iuo
}

// SetItemDescription sets the "item_description" field.
func (iuo *ItemUpdateOne) SetItemDescription(s string) *ItemUpdateOne {
	iuo.mutation.SetItemDescription(s)
	return iuo
}

// SetItemContent sets the "item_content" field.
func (iuo *ItemUpdateOne) SetItemContent(s string) *ItemUpdateOne {
	iuo.mutation.SetItemContent(s)
	return iuo
}

// SetItemLink sets the "item_link" field.
func (iuo *ItemUpdateOne) SetItemLink(s string) *ItemUpdateOne {
	iuo.mutation.SetItemLink(s)
	return iuo
}

// SetItemUpdated sets the "item_updated" field.
func (iuo *ItemUpdateOne) SetItemUpdated(s string) *ItemUpdateOne {
	iuo.mutation.SetItemUpdated(s)
	return iuo
}

// SetItemPublished sets the "item_published" field.
func (iuo *ItemUpdateOne) SetItemPublished(s string) *ItemUpdateOne {
	iuo.mutation.SetItemPublished(s)
	return iuo
}

// SetItemAuthorName sets the "item_author_name" field.
func (iuo *ItemUpdateOne) SetItemAuthorName(s string) *ItemUpdateOne {
	iuo.mutation.SetItemAuthorName(s)
	return iuo
}

// SetItemAuthorEmail sets the "item_author_email" field.
func (iuo *ItemUpdateOne) SetItemAuthorEmail(s string) *ItemUpdateOne {
	iuo.mutation.SetItemAuthorEmail(s)
	return iuo
}

// SetItemImageTitle sets the "item_image_title" field.
func (iuo *ItemUpdateOne) SetItemImageTitle(s string) *ItemUpdateOne {
	iuo.mutation.SetItemImageTitle(s)
	return iuo
}

// SetItemImageURL sets the "item_image_url" field.
func (iuo *ItemUpdateOne) SetItemImageURL(s string) *ItemUpdateOne {
	iuo.mutation.SetItemImageURL(s)
	return iuo
}

// SetItemCategories sets the "item_categories" field.
func (iuo *ItemUpdateOne) SetItemCategories(s string) *ItemUpdateOne {
	iuo.mutation.SetItemCategories(s)
	return iuo
}

// SetItemEnclosures sets the "item_enclosures" field.
func (iuo *ItemUpdateOne) SetItemEnclosures(s string) *ItemUpdateOne {
	iuo.mutation.SetItemEnclosures(s)
	return iuo
}

// SetCrawlerTitle sets the "crawler_title" field.
func (iuo *ItemUpdateOne) SetCrawlerTitle(s string) *ItemUpdateOne {
	iuo.mutation.SetCrawlerTitle(s)
	return iuo
}

// SetCrawlerAuthor sets the "crawler_author" field.
func (iuo *ItemUpdateOne) SetCrawlerAuthor(s string) *ItemUpdateOne {
	iuo.mutation.SetCrawlerAuthor(s)
	return iuo
}

// SetCrawlerExcerpt sets the "crawler_excerpt" field.
func (iuo *ItemUpdateOne) SetCrawlerExcerpt(s string) *ItemUpdateOne {
	iuo.mutation.SetCrawlerExcerpt(s)
	return iuo
}

// SetCrawlerSiteName sets the "crawler_site_name" field.
func (iuo *ItemUpdateOne) SetCrawlerSiteName(s string) *ItemUpdateOne {
	iuo.mutation.SetCrawlerSiteName(s)
	return iuo
}

// SetCrawlerImage sets the "crawler_image" field.
func (iuo *ItemUpdateOne) SetCrawlerImage(s string) *ItemUpdateOne {
	iuo.mutation.SetCrawlerImage(s)
	return iuo
}

// SetCrawlerContentHTML sets the "crawler_content_html" field.
func (iuo *ItemUpdateOne) SetCrawlerContentHTML(s string) *ItemUpdateOne {
	iuo.mutation.SetCrawlerContentHTML(s)
	return iuo
}

// SetCrawlerContentText sets the "crawler_content_text" field.
func (iuo *ItemUpdateOne) SetCrawlerContentText(s string) *ItemUpdateOne {
	iuo.mutation.SetCrawlerContentText(s)
	return iuo
}

// SetCreatedAt sets the "created_at" field.
func (iuo *ItemUpdateOne) SetCreatedAt(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetCreatedAt(t)
	return iuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableCreatedAt(t *time.Time) *ItemUpdateOne {
	if t != nil {
		iuo.SetCreatedAt(*t)
	}
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *ItemUpdateOne) SetUpdatedAt(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (iuo *ItemUpdateOne) SetFeedID(id uuid.UUID) *ItemUpdateOne {
	iuo.mutation.SetFeedID(id)
	return iuo
}

// SetNillableFeedID sets the "feed" edge to the Feed entity by ID if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableFeedID(id *uuid.UUID) *ItemUpdateOne {
	if id != nil {
		iuo = iuo.SetFeedID(*id)
	}
	return iuo
}

// SetFeed sets the "feed" edge to the Feed entity.
func (iuo *ItemUpdateOne) SetFeed(f *Feed) *ItemUpdateOne {
	return iuo.SetFeedID(f.ID)
}

// AddReadByUserIDs adds the "read_by_users" edge to the User entity by IDs.
func (iuo *ItemUpdateOne) AddReadByUserIDs(ids ...uuid.UUID) *ItemUpdateOne {
	iuo.mutation.AddReadByUserIDs(ids...)
	return iuo
}

// AddReadByUsers adds the "read_by_users" edges to the User entity.
func (iuo *ItemUpdateOne) AddReadByUsers(u ...*User) *ItemUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return iuo.AddReadByUserIDs(ids...)
}

// AddReadIDs adds the "reads" edge to the Read entity by IDs.
func (iuo *ItemUpdateOne) AddReadIDs(ids ...uuid.UUID) *ItemUpdateOne {
	iuo.mutation.AddReadIDs(ids...)
	return iuo
}

// AddReads adds the "reads" edges to the Read entity.
func (iuo *ItemUpdateOne) AddReads(r ...*Read) *ItemUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return iuo.AddReadIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearFeed clears the "feed" edge to the Feed entity.
func (iuo *ItemUpdateOne) ClearFeed() *ItemUpdateOne {
	iuo.mutation.ClearFeed()
	return iuo
}

// ClearReadByUsers clears all "read_by_users" edges to the User entity.
func (iuo *ItemUpdateOne) ClearReadByUsers() *ItemUpdateOne {
	iuo.mutation.ClearReadByUsers()
	return iuo
}

// RemoveReadByUserIDs removes the "read_by_users" edge to User entities by IDs.
func (iuo *ItemUpdateOne) RemoveReadByUserIDs(ids ...uuid.UUID) *ItemUpdateOne {
	iuo.mutation.RemoveReadByUserIDs(ids...)
	return iuo
}

// RemoveReadByUsers removes "read_by_users" edges to User entities.
func (iuo *ItemUpdateOne) RemoveReadByUsers(u ...*User) *ItemUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return iuo.RemoveReadByUserIDs(ids...)
}

// ClearReads clears all "reads" edges to the Read entity.
func (iuo *ItemUpdateOne) ClearReads() *ItemUpdateOne {
	iuo.mutation.ClearReads()
	return iuo
}

// RemoveReadIDs removes the "reads" edge to Read entities by IDs.
func (iuo *ItemUpdateOne) RemoveReadIDs(ids ...uuid.UUID) *ItemUpdateOne {
	iuo.mutation.RemoveReadIDs(ids...)
	return iuo
}

// RemoveReads removes "reads" edges to Read entities.
func (iuo *ItemUpdateOne) RemoveReads(r ...*Read) *ItemUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return iuo.RemoveReadIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, iuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Item)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ItemUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := item.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if v, ok := iuo.mutation.ItemLink(); ok {
		if err := item.ItemLinkValidator(v); err != nil {
			return &ValidationError{Name: "item_link", err: fmt.Errorf(`ent: validator failed for field "Item.item_link": %w`, err)}
		}
	}
	return nil
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.ItemGUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemGUID,
		})
	}
	if value, ok := iuo.mutation.ItemTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemTitle,
		})
	}
	if value, ok := iuo.mutation.ItemDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemDescription,
		})
	}
	if value, ok := iuo.mutation.ItemContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemContent,
		})
	}
	if value, ok := iuo.mutation.ItemLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemLink,
		})
	}
	if value, ok := iuo.mutation.ItemUpdated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemUpdated,
		})
	}
	if value, ok := iuo.mutation.ItemPublished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemPublished,
		})
	}
	if value, ok := iuo.mutation.ItemAuthorName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemAuthorName,
		})
	}
	if value, ok := iuo.mutation.ItemAuthorEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemAuthorEmail,
		})
	}
	if value, ok := iuo.mutation.ItemImageTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemImageTitle,
		})
	}
	if value, ok := iuo.mutation.ItemImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemImageURL,
		})
	}
	if value, ok := iuo.mutation.ItemCategories(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemCategories,
		})
	}
	if value, ok := iuo.mutation.ItemEnclosures(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldItemEnclosures,
		})
	}
	if value, ok := iuo.mutation.CrawlerTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerTitle,
		})
	}
	if value, ok := iuo.mutation.CrawlerAuthor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerAuthor,
		})
	}
	if value, ok := iuo.mutation.CrawlerExcerpt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerExcerpt,
		})
	}
	if value, ok := iuo.mutation.CrawlerSiteName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerSiteName,
		})
	}
	if value, ok := iuo.mutation.CrawlerImage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerImage,
		})
	}
	if value, ok := iuo.mutation.CrawlerContentHTML(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerContentHTML,
		})
	}
	if value, ok := iuo.mutation.CrawlerContentText(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldCrawlerContentText,
		})
	}
	if value, ok := iuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldCreatedAt,
		})
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdatedAt,
		})
	}
	if iuo.mutation.FeedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.FeedTable,
			Columns: []string{item.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.FeedTable,
			Columns: []string{item.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.ReadByUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.ReadByUsersTable,
			Columns: item.ReadByUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		createE := &ReadCreate{config: iuo.config, mutation: newReadMutation(iuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedReadByUsersIDs(); len(nodes) > 0 && !iuo.mutation.ReadByUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.ReadByUsersTable,
			Columns: item.ReadByUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReadCreate{config: iuo.config, mutation: newReadMutation(iuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ReadByUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.ReadByUsersTable,
			Columns: item.ReadByUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReadCreate{config: iuo.config, mutation: newReadMutation(iuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.ReadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ReadsTable,
			Columns: []string{item.ReadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: read.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedReadsIDs(); len(nodes) > 0 && !iuo.mutation.ReadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ReadsTable,
			Columns: []string{item.ReadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: read.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ReadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   item.ReadsTable,
			Columns: []string{item.ReadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: read.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
