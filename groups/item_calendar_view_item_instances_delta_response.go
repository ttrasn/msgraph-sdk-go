package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarViewItemInstancesDeltaGetResponseable instead.
type ItemCalendarViewItemInstancesDeltaResponse struct {
    ItemCalendarViewItemInstancesDeltaGetResponse
}
// NewItemCalendarViewItemInstancesDeltaResponse instantiates a new ItemCalendarViewItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarViewItemInstancesDeltaResponse()(*ItemCalendarViewItemInstancesDeltaResponse) {
    m := &ItemCalendarViewItemInstancesDeltaResponse{
        ItemCalendarViewItemInstancesDeltaGetResponse: *NewItemCalendarViewItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarViewItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarViewItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarViewItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarViewItemInstancesDeltaGetResponseable instead.
type ItemCalendarViewItemInstancesDeltaResponseable interface {
    ItemCalendarViewItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
