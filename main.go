package lecture8_encrypt

func Encrypt(text string, key string) string {
	data := []byte(text)
	for i := range data {
		data[i] ^= key[i%len(key)]
	}
	return string(data)
}

func decrypt_text(text string, key string) string {
	return Encrypt(text, key)
}
