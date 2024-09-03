package filename_parser

var regexpAudioChannels = mustCompileRegexpList([]string{
	`(?i)\b7\.?[01]\b`,
	`(?i)\b(6[\W_]0(?:ch)?(?!\d)|5[\W_][01](?:ch)?(?!\d)|5ch|6ch)\b`,
	`(?i)\b(2[\W_]0(?:ch)?(?!\d)|stereo)\b`,
	`(?i)\b(1[\W_]0(?:ch)?(?!\d)|mono|1ch)\b`,
})

const (
	AudioChannelSeven  string = "7.1"
	AudioChannelSix    string = "5.1"
	AudioChannelStereo string = "sterio"
	AudioChannelMono   string = "mono"
)

var audioChannels = []string{
	AudioChannelSeven,
	AudioChannelSix,
	AudioChannelStereo,
	AudioChannelMono,
}

func ParseAudioChannels(s string) string {
	return regexpAudioChannels.parseFromPatterns(s, audioChannels)
}

var regexpAudioCodacs = mustCompileRegexpList([]string{
	`(?i)\b(LAME(?:\d)+-?(?:\d)+|mp3)\b`,
	`(?i)\b(mp2)\b`,
	`(?i)\b(Dolby|Dolby-?Digital|DD|AC3D?)\b`,
	`(?i)\b(Dolby-?Atmos)\b`,
	`(?i)\b(AAC)(\d?.?\d?)(ch)?\b`,
	`(?i)\b(EAC3|DDP|DD\+)\b`,
	`(?i)\b(FLAC)\b`,
	`(?i)\b(DTS)\b`,
	`(?i)\b(DTS-?HD|DTS(?=-?MA)|DTS-X)\b`,
	`(?i)\b(True-?HD)\b`,
	`(?i)\b(Opus)\b`,
	`(?i)\b(Vorbis)\b`,
	`(?i)\b(PCM)\b`,
	`(?i)\b(LPCM)\b`,
})

const (
	AudioCodacMP3        string = "MP3"
	AudioCodacMP2        string = "MP2"
	AudioCodacDolby      string = "Dolby Digital"
	AudioCodacDolbyAtmos string = "Dolby Atmos"
	AudioCodacAAC        string = "AAC"
	AudioCodacEAC3       string = "EAC3"
	AudioCodacFLAC       string = "FLAC"
	AudioCodacDTS        string = "DTS"
	AudioCodacDTSHD      string = "DTS-HD"
	AudioCodacTrueHD     string = "TrueHD"
	AudioCodacOpus       string = "Opus"
	AudioCodacVorbis     string = "Vorbis"
	AudioCodacPCM        string = "PCM"
	AudioCodacLPCM       string = "LPCM"
)

var audioCodacs = []string{
	AudioCodacMP3,
	AudioCodacMP2,
	AudioCodacDolby,
	AudioCodacDolbyAtmos,
	AudioCodacAAC,
	AudioCodacEAC3,
	AudioCodacFLAC,
	AudioCodacDTS,
	AudioCodacDTSHD,
	AudioCodacTrueHD,
	AudioCodacOpus,
	AudioCodacVorbis,
	AudioCodacPCM,
	AudioCodacLPCM,
}

func ParseAudioCodac(s string) string {
	return regexpAudioCodacs.parseFromPatterns(s, audioCodacs)
}

type ParsedAudio struct {
	Codac    string
	Channels string
}

func ParseAudio(filename string) ParsedAudio {
	return ParsedAudio{
		Codac:    ParseAudioCodac(filename),
		Channels: ParseAudioChannels(filename),
	}
}
