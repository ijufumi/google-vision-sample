package enums

import "golang.org/x/exp/slices"

type ContentType string

const (
	ContentType_Pdf         ContentType = "application/pdf"
	ContentType_Jpeg        ContentType = "image/jpeg"
	ContentType_Png         ContentType = "image/png"
	ContentType_Gif         ContentType = "image/gif"
	ContentType_OctetStream ContentType = "application/octet-stream"
	ContentType_Json        ContentType = "application/json"
)

var imageContentTypes = []ContentType{
	ContentType_Jpeg,
	ContentType_Png,
	ContentType_Gif,
}

func ConvertToContentType(contentType string) ContentType {
	switch contentType {
	case string(ContentType_Pdf):
		return ContentType_Pdf
	case string(ContentType_Jpeg):
		return ContentType_Jpeg
	case string(ContentType_Png):
		return ContentType_Png
	case string(ContentType_Gif):
		return ContentType_Gif
	case string(ContentType_Json):
		return ContentType_Json
	}
	return ContentType_OctetStream
}

func IsImage(contentType ContentType) bool {
	return slices.Contains(imageContentTypes, contentType)
}
