package item

type LocationType string

type Location struct {
	LocationType
	BodyLocation LocationType
	Page         int
}

const (
	// Storage locations
	LocationUnknown     LocationType = "unknown"
	LocationInventory   LocationType = "inventory"
	LocationStash       LocationType = "stash"
	LocationSharedStash LocationType = "shared_stash"
	LocationBelt        LocationType = "belt"
	LocationCube        LocationType = "cube"
	LocationVendor      LocationType = "vendor"
	LocationGround      LocationType = "ground"
	LocationSocket      LocationType = "socket"
	LocationCursor      LocationType = "cursor"
	LocationEquipped    LocationType = "equipped"
	LocationMercenary   LocationType = "mercenary"

	// Body locations
	LocNone              LocationType = "none"
	LocHead              LocationType = "head"
	LocNeck              LocationType = "neck"
	LocTorso             LocationType = "torso"
	LocLeftArm           LocationType = "left_arm"
	LocRightArm          LocationType = "right_arm"
	LocLeftRing          LocationType = "left_ring"
	LocRightRing         LocationType = "right_ring"
	LocBelt              LocationType = "belt"
	LocFeet              LocationType = "feet"
	LocGloves            LocationType = "gloves"
	LocLeftArmSecondary  LocationType = "left_arm_secondary"
	LocRightArmSecondary LocationType = "right_arm_secondary"
)
