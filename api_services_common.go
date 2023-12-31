package broadpeakio

// Misc types start
type GetAllSlotsInput struct {
	Offset     uint
	Limit      uint
	Categories []uint
	From       string
	To         string
}

type AdBreakInsertion struct {
	AdServer  *Identifiable `json:"adServer,omitempty"`
	GapFilter *Identifiable `json:"gapFilter,omitempty"`
}

type ServerSideAdTracking struct {
	Enable                          bool `json:"enable,omitempty"`
	CheckAdMediaSegmentAvailability bool `json:"checkAdMediaSegmentAvailability,omitempty"`
}

type LiveAdPreRoll struct {
	AdServer    *Identifiable `json:"adServer,omitempty"` //required
	MaxDuration uint          `json:"maxDuration,omitempty"`
	Offset      uint          `json:"offset,omitempty"`
}

type LiveAdReplacement struct {
	AdServer  *Identifiable `json:"adServer,omitempty"` //required
	GapFiller *Identifiable `json:"gapFiller,omitempty"`
}

type VodAdInsertion struct {
	AdServer *Identifiable `json:"adServer,omitempty"` //required
}

type AdServer struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Queries string `json:"queries"`
}

type Source struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

type TranscodingProfile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Id      uint   `json:"id"`
}

type LiveAdPreRollOutput struct {
	AdServer    AdServer `json:"adServer,omitempty"`
	MaxDuration string   `json:"maxDuration"`
	Offset      uint     `json:"offset"`
}

type LiveAdReplacementOutput struct {
	AdServer  AdServer `json:"adServer"`
	GapFiller Source   `json:"gapFiller"`
}

type AdBreakInsertionOutput struct {
	AdServer  AdServer `json:"adServer"`
	GapFiller Source   `json:"gapFiller"`
}

type VodAdInsertionOutput struct {
	AdServer AdServer `json:"adServer"`
}

// Misc types end
