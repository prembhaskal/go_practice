package bit

func reverseBits(num uint32) uint32 {
	var rev uint32
	rev = 0
	for n := 0; n < 32; n++ {
		//         fmt.Printf("num:%b\n",num)
		//         fmt.Printf("rev: %0b\n", rev)

		lastbit := num & 1
		rev = rev << 1
		if lastbit == 1 {
			rev = rev | 1
		}
		num = num >> 1
	}
	return rev
}
