// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/item"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ItemGUID holds the value of the "item_guid" field.
	ItemGUID string `json:"item_guid,omitempty"`
	// ItemTitle holds the value of the "item_title" field.
	ItemTitle string `json:"item_title,omitempty"`
	// ItemDescription holds the value of the "item_description" field.
	ItemDescription string `json:"item_description,omitempty"`
	// ItemContent holds the value of the "item_content" field.
	ItemContent string `json:"item_content,omitempty"`
	// ItemLink holds the value of the "item_link" field.
	ItemLink string `json:"item_link,omitempty"`
	// ItemUpdated holds the value of the "item_updated" field.
	ItemUpdated time.Time `json:"item_updated,omitempty"`
	// ItemPublished holds the value of the "item_published" field.
	ItemPublished time.Time `json:"item_published,omitempty"`
	// ItemAuthorName holds the value of the "item_author_name" field.
	ItemAuthorName string `json:"item_author_name,omitempty"`
	// ItemAuthorEmail holds the value of the "item_author_email" field.
	ItemAuthorEmail string `json:"item_author_email,omitempty"`
	// ItemImageTitle holds the value of the "item_image_title" field.
	ItemImageTitle string `json:"item_image_title,omitempty"`
	// ItemImageURL holds the value of the "item_image_url" field.
	ItemImageURL string `json:"item_image_url,omitempty"`
	// ItemCategories holds the value of the "item_categories" field.
	ItemCategories string `json:"item_categories,omitempty"`
	// ItemEnclosures holds the value of the "item_enclosures" field.
	ItemEnclosures string `json:"item_enclosures,omitempty"`
	// CrawlerTitle holds the value of the "crawler_title" field.
	CrawlerTitle string `json:"crawler_title,omitempty"`
	// CrawlerAuthor holds the value of the "crawler_author" field.
	CrawlerAuthor string `json:"crawler_author,omitempty"`
	// CrawlerExcerpt holds the value of the "crawler_excerpt" field.
	CrawlerExcerpt string `json:"crawler_excerpt,omitempty"`
	// CrawlerSiteName holds the value of the "crawler_site_name" field.
	CrawlerSiteName string `json:"crawler_site_name,omitempty"`
	// CrawlerImage holds the value of the "crawler_image" field.
	CrawlerImage string `json:"crawler_image,omitempty"`
	// CrawlerContentHTML holds the value of the "crawler_content_html" field.
	CrawlerContentHTML string `json:"crawler_content_html,omitempty"`
	// CrawlerContentText holds the value of the "crawler_content_text" field.
	CrawlerContentText string `json:"crawler_content_text,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges        ItemEdges `json:"edges"`
	feed_items   *uuid.UUID
	selectValues sql.SelectValues
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Feed holds the value of the feed edge.
	Feed *Feed `json:"feed,omitempty"`
	// ReadByUsers holds the value of the read_by_users edge.
	ReadByUsers []*User `json:"read_by_users,omitempty"`
	// Reads holds the value of the reads edge.
	Reads []*Read `json:"reads,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// FeedOrErr returns the Feed value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) FeedOrErr() (*Feed, error) {
	if e.Feed != nil {
		return e.Feed, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: feed.Label}
	}
	return nil, &NotLoadedError{edge: "feed"}
}

// ReadByUsersOrErr returns the ReadByUsers value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) ReadByUsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.ReadByUsers, nil
	}
	return nil, &NotLoadedError{edge: "read_by_users"}
}

// ReadsOrErr returns the Reads value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) ReadsOrErr() ([]*Read, error) {
	if e.loadedTypes[2] {
		return e.Reads, nil
	}
	return nil, &NotLoadedError{edge: "reads"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldItemGUID, item.FieldItemTitle, item.FieldItemDescription, item.FieldItemContent, item.FieldItemLink, item.FieldItemAuthorName, item.FieldItemAuthorEmail, item.FieldItemImageTitle, item.FieldItemImageURL, item.FieldItemCategories, item.FieldItemEnclosures, item.FieldCrawlerTitle, item.FieldCrawlerAuthor, item.FieldCrawlerExcerpt, item.FieldCrawlerSiteName, item.FieldCrawlerImage, item.FieldCrawlerContentHTML, item.FieldCrawlerContentText:
			values[i] = new(sql.NullString)
		case item.FieldItemUpdated, item.FieldItemPublished, item.FieldCreatedAt, item.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case item.FieldID:
			values[i] = new(uuid.UUID)
		case item.ForeignKeys[0]: // feed_items
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case item.FieldItemGUID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_guid", values[j])
			} else if value.Valid {
				i.ItemGUID = value.String
			}
		case item.FieldItemTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_title", values[j])
			} else if value.Valid {
				i.ItemTitle = value.String
			}
		case item.FieldItemDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_description", values[j])
			} else if value.Valid {
				i.ItemDescription = value.String
			}
		case item.FieldItemContent:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_content", values[j])
			} else if value.Valid {
				i.ItemContent = value.String
			}
		case item.FieldItemLink:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_link", values[j])
			} else if value.Valid {
				i.ItemLink = value.String
			}
		case item.FieldItemUpdated:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field item_updated", values[j])
			} else if value.Valid {
				i.ItemUpdated = value.Time
			}
		case item.FieldItemPublished:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field item_published", values[j])
			} else if value.Valid {
				i.ItemPublished = value.Time
			}
		case item.FieldItemAuthorName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_author_name", values[j])
			} else if value.Valid {
				i.ItemAuthorName = value.String
			}
		case item.FieldItemAuthorEmail:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_author_email", values[j])
			} else if value.Valid {
				i.ItemAuthorEmail = value.String
			}
		case item.FieldItemImageTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_image_title", values[j])
			} else if value.Valid {
				i.ItemImageTitle = value.String
			}
		case item.FieldItemImageURL:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_image_url", values[j])
			} else if value.Valid {
				i.ItemImageURL = value.String
			}
		case item.FieldItemCategories:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_categories", values[j])
			} else if value.Valid {
				i.ItemCategories = value.String
			}
		case item.FieldItemEnclosures:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_enclosures", values[j])
			} else if value.Valid {
				i.ItemEnclosures = value.String
			}
		case item.FieldCrawlerTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawler_title", values[j])
			} else if value.Valid {
				i.CrawlerTitle = value.String
			}
		case item.FieldCrawlerAuthor:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawler_author", values[j])
			} else if value.Valid {
				i.CrawlerAuthor = value.String
			}
		case item.FieldCrawlerExcerpt:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawler_excerpt", values[j])
			} else if value.Valid {
				i.CrawlerExcerpt = value.String
			}
		case item.FieldCrawlerSiteName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawler_site_name", values[j])
			} else if value.Valid {
				i.CrawlerSiteName = value.String
			}
		case item.FieldCrawlerImage:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawler_image", values[j])
			} else if value.Valid {
				i.CrawlerImage = value.String
			}
		case item.FieldCrawlerContentHTML:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawler_content_html", values[j])
			} else if value.Valid {
				i.CrawlerContentHTML = value.String
			}
		case item.FieldCrawlerContentText:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawler_content_text", values[j])
			} else if value.Valid {
				i.CrawlerContentText = value.String
			}
		case item.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case item.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case item.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field feed_items", values[j])
			} else if value.Valid {
				i.feed_items = new(uuid.UUID)
				*i.feed_items = *value.S.(*uuid.UUID)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Item.
// This includes values selected through modifiers, order, etc.
func (i *Item) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryFeed queries the "feed" edge of the Item entity.
func (i *Item) QueryFeed() *FeedQuery {
	return NewItemClient(i.config).QueryFeed(i)
}

// QueryReadByUsers queries the "read_by_users" edge of the Item entity.
func (i *Item) QueryReadByUsers() *UserQuery {
	return NewItemClient(i.config).QueryReadByUsers(i)
}

// QueryReads queries the "reads" edge of the Item entity.
func (i *Item) QueryReads() *ReadQuery {
	return NewItemClient(i.config).QueryReads(i)
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return NewItemClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("item_guid=")
	builder.WriteString(i.ItemGUID)
	builder.WriteString(", ")
	builder.WriteString("item_title=")
	builder.WriteString(i.ItemTitle)
	builder.WriteString(", ")
	builder.WriteString("item_description=")
	builder.WriteString(i.ItemDescription)
	builder.WriteString(", ")
	builder.WriteString("item_content=")
	builder.WriteString(i.ItemContent)
	builder.WriteString(", ")
	builder.WriteString("item_link=")
	builder.WriteString(i.ItemLink)
	builder.WriteString(", ")
	builder.WriteString("item_updated=")
	builder.WriteString(i.ItemUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("item_published=")
	builder.WriteString(i.ItemPublished.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("item_author_name=")
	builder.WriteString(i.ItemAuthorName)
	builder.WriteString(", ")
	builder.WriteString("item_author_email=")
	builder.WriteString(i.ItemAuthorEmail)
	builder.WriteString(", ")
	builder.WriteString("item_image_title=")
	builder.WriteString(i.ItemImageTitle)
	builder.WriteString(", ")
	builder.WriteString("item_image_url=")
	builder.WriteString(i.ItemImageURL)
	builder.WriteString(", ")
	builder.WriteString("item_categories=")
	builder.WriteString(i.ItemCategories)
	builder.WriteString(", ")
	builder.WriteString("item_enclosures=")
	builder.WriteString(i.ItemEnclosures)
	builder.WriteString(", ")
	builder.WriteString("crawler_title=")
	builder.WriteString(i.CrawlerTitle)
	builder.WriteString(", ")
	builder.WriteString("crawler_author=")
	builder.WriteString(i.CrawlerAuthor)
	builder.WriteString(", ")
	builder.WriteString("crawler_excerpt=")
	builder.WriteString(i.CrawlerExcerpt)
	builder.WriteString(", ")
	builder.WriteString("crawler_site_name=")
	builder.WriteString(i.CrawlerSiteName)
	builder.WriteString(", ")
	builder.WriteString("crawler_image=")
	builder.WriteString(i.CrawlerImage)
	builder.WriteString(", ")
	builder.WriteString("crawler_content_html=")
	builder.WriteString(i.CrawlerContentHTML)
	builder.WriteString(", ")
	builder.WriteString("crawler_content_text=")
	builder.WriteString(i.CrawlerContentText)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item
