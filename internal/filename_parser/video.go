package filename_parser

const (
	FileTypeWEBM  string = "webm"
	FileTypeM4V   string = "m4v"
	FileType3GP   string = "3gp"
	FileTypeNSV   string = "nsv"
	FileTypeTY    string = "ty"
	FileTypeSTRM  string = "strm"
	FileTypeRM    string = "rm"
	FileTypeRMVB  string = "rmvb"
	FileTypeM3U   string = "m3u"
	FileTypeIFO   string = "ifo"
	FileTypeMOV   string = "mov"
	FileTypeQT    string = "qt"
	FileTypeDIVX  string = "divx"
	FileTypeXVID  string = "xvid"
	FileTypeBIVX  string = "bivx"
	FileTypeNRG   string = "nrg"
	FileTypePVA   string = "pva"
	FileTypeWMV   string = "wmv"
	FileTypeASF   string = "asf"
	FileTypeASX   string = "asx"
	FileTypeOGM   string = "ogm"
	FileTypeOGV   string = "ogv"
	FileTypeM2V   string = "m2v"
	FileTypeAVI   string = "avi"
	FileTypeBIN   string = "bin"
	FileTypeDAT   string = "dat"
	FileTypeDVRMS string = "dvr-ms"
	FileTypeMPG   string = "mpg"
	FileTypeMPEG  string = "mpeg"
	FileTypeMP4   string = "mp4"
	FileTypeAVC   string = "avc"
	FileTypeVP3   string = "vp3"
	FileTypeSVQ3  string = "svq3"
	FileTypeNUV   string = "nuv"
	FileTypeVIV   string = "viv"
	FileTypeDV    string = "dv"
	FileTypeFLI   string = "fli"
	FileTypeFLV   string = "flv"
	FileTypeWPL   string = "wpl"
	FileTypeIMG   string = "img"
	FileTypeISO   string = "iso"
	FileTypeVOB   string = "vob"
	FileTypeMKV   string = "mkv"
	FileTypeMK3D  string = "mk3d"
	FileTypeTS    string = "ts"
	FileTypeWTV   string = "wtv"
	FileTypeM2TS  string = "m2ts"
)

var fileTypes = []string{
	FileTypeWEBM, FileTypeM4V, FileType3GP,
	FileTypeNSV, FileTypeTY, FileTypeSTRM,
	FileTypeRM, FileTypeRMVB, FileTypeM3U,
	FileTypeIFO, FileTypeMOV, FileTypeQT,
	FileTypeDIVX, FileTypeXVID, FileTypeBIVX,
	FileTypeNRG, FileTypePVA, FileTypeWMV,
	FileTypeASF, FileTypeASX, FileTypeOGM,
	FileTypeOGV, FileTypeM2V, FileTypeAVI,
	FileTypeBIN, FileTypeDAT, FileTypeDVRMS,
	FileTypeMPG, FileTypeMPEG, FileTypeMP4,
	FileTypeAVC, FileTypeVP3, FileTypeSVQ3,
	FileTypeNUV, FileTypeVIV, FileTypeDV,
	FileTypeFLI, FileTypeFLV, FileTypeWPL,
	FileTypeIMG, FileTypeISO, FileTypeVOB,
	FileTypeMKV, FileTypeMK3D, FileTypeTS,
	FileTypeWTV, FileTypeM2TS,
}

var regexpFileTypes = mustCompileRegexpList(fileTypes)

func ParseFileType(s string) string {
	return regexpFileTypes.parseFromPatterns(s, fileTypes)
}
