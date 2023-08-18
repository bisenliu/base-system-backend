package response

type CaptchaInfo struct {
	JigsawImageBase64   interface{} `json:"jigsaw_image_base_64"`
	OriginalImageBase64 interface{} `json:"original_image_base_64"`
	SecretKey           interface{} `json:"secret_key"`
	Token               interface{} `json:"token"`
	WordList            interface{} `json:"word_list"`
}
