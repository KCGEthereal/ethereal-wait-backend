# Models folder

You can keep type structs of every model used within your app.

### Example

Banner service model struct

```go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Banner struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Type         string `bson:"type"`
	Link         string `bson:"link"`
	Priority     int    `bson:"priority"`
	Disabled     bool   `bson:"isDisabled"`
	DesktopImage string `bson:"desktopImage"`
	MobileImage  string `bson:"mobileImage"`

	CreatedBy string    `bson:"createdBy"`
	CreatedAt time.Time `bson:"createdAt"`

	UpdatedBy string    `bson:"updatedBy,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}
```
