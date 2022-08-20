// Code generated by ent, DO NOT EDIT.

package item

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldItemGUID holds the string denoting the item_guid field in the database.
	FieldItemGUID = "item_guid"
	// FieldItemTitle holds the string denoting the item_title field in the database.
	FieldItemTitle = "item_title"
	// FieldItemDescription holds the string denoting the item_description field in the database.
	FieldItemDescription = "item_description"
	// FieldItemContent holds the string denoting the item_content field in the database.
	FieldItemContent = "item_content"
	// FieldItemLink holds the string denoting the item_link field in the database.
	FieldItemLink = "item_link"
	// FieldItemUpdated holds the string denoting the item_updated field in the database.
	FieldItemUpdated = "item_updated"
	// FieldItemPublished holds the string denoting the item_published field in the database.
	FieldItemPublished = "item_published"
	// FieldItemAuthorName holds the string denoting the item_author_name field in the database.
	FieldItemAuthorName = "item_author_name"
	// FieldItemAuthorEmail holds the string denoting the item_author_email field in the database.
	FieldItemAuthorEmail = "item_author_email"
	// FieldItemImageTitle holds the string denoting the item_image_title field in the database.
	FieldItemImageTitle = "item_image_title"
	// FieldItemImageURL holds the string denoting the item_image_url field in the database.
	FieldItemImageURL = "item_image_url"
	// FieldItemCategories holds the string denoting the item_categories field in the database.
	FieldItemCategories = "item_categories"
	// FieldItemEnclosures holds the string denoting the item_enclosures field in the database.
	FieldItemEnclosures = "item_enclosures"
	// FieldCrawlerTitle holds the string denoting the crawler_title field in the database.
	FieldCrawlerTitle = "crawler_title"
	// FieldCrawlerAuthor holds the string denoting the crawler_author field in the database.
	FieldCrawlerAuthor = "crawler_author"
	// FieldCrawlerExcerpt holds the string denoting the crawler_excerpt field in the database.
	FieldCrawlerExcerpt = "crawler_excerpt"
	// FieldCrawlerSiteName holds the string denoting the crawler_site_name field in the database.
	FieldCrawlerSiteName = "crawler_site_name"
	// FieldCrawlerImage holds the string denoting the crawler_image field in the database.
	FieldCrawlerImage = "crawler_image"
	// FieldCrawlerContentHTML holds the string denoting the crawler_content_html field in the database.
	FieldCrawlerContentHTML = "crawler_content_html"
	// FieldCrawlerContentText holds the string denoting the crawler_content_text field in the database.
	FieldCrawlerContentText = "crawler_content_text"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeFeed holds the string denoting the feed edge name in mutations.
	EdgeFeed = "feed"
	// EdgeReadByUsers holds the string denoting the read_by_users edge name in mutations.
	EdgeReadByUsers = "read_by_users"
	// EdgeReads holds the string denoting the reads edge name in mutations.
	EdgeReads = "reads"
	// Table holds the table name of the item in the database.
	Table = "items"
	// FeedTable is the table that holds the feed relation/edge.
	FeedTable = "items"
	// FeedInverseTable is the table name for the Feed entity.
	// It exists in this package in order to avoid circular dependency with the "feed" package.
	FeedInverseTable = "feeds"
	// FeedColumn is the table column denoting the feed relation/edge.
	FeedColumn = "feed_items"
	// ReadByUsersTable is the table that holds the read_by_users relation/edge. The primary key declared below.
	ReadByUsersTable = "reads"
	// ReadByUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ReadByUsersInverseTable = "users"
	// ReadsTable is the table that holds the reads relation/edge.
	ReadsTable = "reads"
	// ReadsInverseTable is the table name for the Read entity.
	// It exists in this package in order to avoid circular dependency with the "read" package.
	ReadsInverseTable = "reads"
	// ReadsColumn is the table column denoting the reads relation/edge.
	ReadsColumn = "item_id"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldItemGUID,
	FieldItemTitle,
	FieldItemDescription,
	FieldItemContent,
	FieldItemLink,
	FieldItemUpdated,
	FieldItemPublished,
	FieldItemAuthorName,
	FieldItemAuthorEmail,
	FieldItemImageTitle,
	FieldItemImageURL,
	FieldItemCategories,
	FieldItemEnclosures,
	FieldCrawlerTitle,
	FieldCrawlerAuthor,
	FieldCrawlerExcerpt,
	FieldCrawlerSiteName,
	FieldCrawlerImage,
	FieldCrawlerContentHTML,
	FieldCrawlerContentText,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"feed_items",
}

var (
	// ReadByUsersPrimaryKey and ReadByUsersColumn2 are the table columns denoting the
	// primary key for the read_by_users relation (M2M).
	ReadByUsersPrimaryKey = []string{"user_id", "item_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ItemLinkValidator is a validator for the "item_link" field. It is called by the builders before save.
	ItemLinkValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
