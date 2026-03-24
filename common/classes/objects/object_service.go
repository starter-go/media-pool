package objects

type Service interface {

	// fetch

	DownloadFilterChain

	// put

	UploadFilterChain
}

type Client interface {
	Service
}

type Server interface {
	Service
}
