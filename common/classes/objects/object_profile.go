package objects

// Profile 表示对象的视图形式
type Profile int

const (
	ProfileMeta     Profile = 0
	ProfileThumbMin Profile = 1
	ProfileThumb    Profile = 100
	ProfileThumbMax Profile = 1900

	ProfileData Profile = ProfileThumbMax + 1
)

////////////////////////////////////////////////////////////////////////////////

func (p Profile) IsData() bool {
	return (p >= ProfileData)
}

func (p Profile) IsMeta() bool {
	return (p == ProfileMeta)
}

func (p Profile) IsThumbnail() bool {
	return ((ProfileThumbMin <= p) && (p <= ProfileThumbMax))
}
