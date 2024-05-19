package item

type LocationType string

type Location struct {
	LocationType
	Page int
}

const (
	LocationEquipped    LocationType = "equipped"
	LocationStash       LocationType = "stash"
	LocationSharedStash LocationType = "shared_stash"
	LocationBelt        LocationType = "belt"
	LocationInventory   LocationType = "inventory"
	LocationCube        LocationType = "cube"
	LocationVendor      LocationType = "vendor"
	LocationGround      LocationType = "ground"
	LocationSocket      LocationType = "socket"
	LocationUnknown     LocationType = "unknown"
	LocationCursor      LocationType = "cursor"
)
