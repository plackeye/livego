package av

type StaticsBW struct {
	StreamId               uint32
	PeerIP                 string
	VideoDatainBytes       uint64
	LastVideoDatainBytes   uint64
	VideoSpeedInBytesperMS uint64

	AudioDatainBytes       uint64
	LastAudioDatainBytes   uint64
	AudioSpeedInBytesperMS uint64

	LastTimestamp int64
}

const SAVE_STATICS_INTERVAL = 5000
