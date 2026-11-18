package models

type Collection struct {
	Title     *string
	SortOrder *string
	Handle    *string
	Sources   []CollectionSource
}

type CollectionSource struct {
	MatchType  *string
	Conditions []map[string]CollectionCondition
}

type CollectionCondition struct {
	Field     *string
	Relation  *string
	MatchType *string
	Value     any
}

type CreateCollectionRequest struct {
	Collection *Collection
}

type CollectionResponse struct{}
