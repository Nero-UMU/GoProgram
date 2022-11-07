package aes

var RC = [10]int{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1b, 0x36}

func keyExpansion(key *[16]int, w *[176]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			w[j*4+i] = key[j*4+i]
		}
	}
	for i := 1; i < 11; i++ {
		for j := 0; j < 4; j++ {
			var temp [4]int
			if j == 0 {
				temp[0] = w[(i-1)*16+3*4+1]
				temp[1] = w[(i-1)*16+3*4+2]
				temp[2] = w[(i-1)*16+3*4+3]
				temp[3] = w[(i-1)*16+3*4+0]
				for k := 0; k < 4; k++ {
					m := temp[k]
					temp[k] = calcBox(m)
					if k == 0 {
						temp[k] = temp[k] ^ RC[i-1]
					}
				}
			} else {
				temp[0] = w[i*16+(j-1)*4+0]
				temp[1] = w[i*16+(j-1)*4+1]
				temp[2] = w[i*16+(j-1)*4+2]
				temp[3] = w[i*16+(j-1)*4+3]
			}
			for l := 0; l < 4; l++ {

				w[i*16+j*4+l] = w[(i-1)*16+j*4+l] ^ temp[l]
			}

		}
	}
}
