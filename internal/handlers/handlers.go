package handlers

type Ctx struct {
	Proto      string `json:"proto"`
	StatusCode int    `json:"status"`
}

type Handler interface {
	HandleRequest(ctx *Ctx)
	SetNext(handler Handler)
}

func SetNextReferences(hs []Handler) Handler {
	for i := 0; i < len(hs)-1; i++ {
		hs[i].SetNext(hs[i+1])
	}

	return hs[0]
}
