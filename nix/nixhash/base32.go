package nixhash

const base32Chars = "0123456789abcdfghijklmnpqrsvwxyz"

func base32Encode(data []byte) []byte {
	s := make([]byte, (len(data)*8+4)/5)
	n, i, j := len(data)-1, len(data), len(s)*5%8
	for k := range s {
		j -= 5
		if j < 0 {
			i, j = i-1, j+8
		}
		c := data[i] >> uint8(j)
		if i < n {
			c |= data[i+1] << uint8(8-j)
		}
		s[k] = base32Chars[c&0x1f]
	}
	return s
}
