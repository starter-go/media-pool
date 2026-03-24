package downloads

import "github.com/starter-go/media-pool/common/classes/objects"

type Service interface {

	// the chain
	objects.DownloadFilterChain
}
