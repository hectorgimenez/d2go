package item

type Location string

const (
	LocationEquipped     Location = "equipped"
	LocationStash        Location = "stash"
	LocationSharedStash1 Location = "shared_stash_1"
	LocationSharedStash2 Location = "shared_stash_2"
	LocationSharedStash3 Location = "shared_stash_3"
	LocationBelt         Location = "belt"
	LocationInventory    Location = "inventory"
	LocationCube         Location = "cube"
	LocationVendor       Location = "vendor"
	LocationGround       Location = "ground"
	LocationSocket       Location = "socket"
	LocationUnknown      Location = "unknown"
)
