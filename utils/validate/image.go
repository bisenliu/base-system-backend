package validate

import (
	"base-system-backend/enums/errmsg"
	"base-system-backend/global"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

var ext = []string{
	"ase",
	"art",
	"bmp",
	"blp",
	"cd5",
	"cit",
	"cpt",
	"cr2",
	"cut",
	"dds",
	"dib",
	"djvu",
	"egt",
	"exif",
	"gif",
	"gpl",
	"grf",
	"icns",
	"ico",
	"iff",
	"jng",
	"jpeg",
	"jpg",
	"jfif",
	"jp2",
	"jps",
	"lbm",
	"max",
	"miff",
	"mng",
	"msp",
	"nitf",
	"ota",
	"pbm",
	"pc1",
	"pc2",
	"pc3",
	"pcf",
	"pcx",
	"pdn",
	"pgm",
	"PI1",
	"PI2",
	"PI3",
	"pict",
	"pct",
	"pnm",
	"pns",
	"ppm",
	"psb",
	"psd",
	"pdd",
	"psp",
	"px",
	"pxm",
	"pxr",
	"qfx",
	"raw",
	"rle",
	"sct",
	"sgi",
	"rgb",
	"int",
	"bw",
	"tga",
	"tiff",
	"tif",
	"vtf",
	"xbm",
	"xcf",
	"xpm",
	"3dv",
	"amf",
	"ai",
	"awg",
	"cgm",
	"cdr",
	"cmx",
	"dxf",
	"e2d",
	"egt",
	"eps",
	"fs",
	"gbr",
	"odg",
	"svg",
	"stl",
	"vrml",
	"x3d",
	"sxd",
	"v2d",
	"vnd",
	"wmf",
	"emf",
	"art",
	"xar",
	"png",
	"webp",
	"jxr",
	"hdp",
	"wdp",
	"cur",
	"ecw",
	"iff",
	"lbm",
	"liff",
	"nrrd",
	"pam",
	"pcx",
	"pgf",
	"sgi",
	"rgb",
	"rgba",
	"bw",
	"int",
	"inta",
	"sid",
	"ras",
	"sun",
	"tga",
}

func ImageVerify(fileHeader *multipart.FileHeader) (err error, debugInfo interface{}) {
	if fileHeader.Size > global.CONFIG.Avatar.Size*1024*1024 {
		return fmt.Errorf(errmsg.FileSizeRange.Error(), int(global.CONFIG.Avatar.Size)), nil
	}
	buff := make([]byte, 512)
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("头像文件%w", errmsg.Invalid), err.Error()
	}
	if _, err = file.Read(buff); err != nil {
		return fmt.Errorf("头像文件%w", errmsg.ReadFailed), err.Error()
	}
	filetype := http.DetectContentType(buff)
	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:]) {
			return
		}
	}
	return fmt.Errorf("头像文件%w", errmsg.Invalid), nil
}
