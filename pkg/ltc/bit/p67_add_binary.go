package bit

func addBinary(a string, b string) string {
	ar := []rune(a)
	br := []rune(b)
	ai := len(ar) - 1
	bi := len(br) - 1
	car := false

	nr := make([]rune, 0)
	for ai >= 0 || bi >= 0 {
		var sm rune
		if ai >= 0 && bi >= 0 {
			sm, car = sum1(ar[ai], br[bi], car)
		} else if ai >= 0 {
			sm, car = sum1(ar[ai], '0', car)
		} else {
			sm, car = sum1(br[bi], '0', car)
		}
		nr = append(nr, sm)
		ai--
		bi--
	}

	if car {
		nr = append(nr, '1')
	}
	return string(reverse(nr))
}

func reverse(rn []rune) []rune {
	i := 0
	j := len(rn) - 1
	for i < j {
		rn[i], rn[j] = rn[j], rn[i]
		i++
		j--
	}
	return rn
}

func sum1(a, b rune, car bool) (rune, bool) {
	sum := 0
	sum = sum + int(a-'0')
	sum = sum + int(b-'0')
	if car {
		sum++
	}
	return rune(sum%2) + '0', sum >= 2
}

// func sum(a , b rune, car bool) (rune, bool) {
//     if a == '1' {
//         if b == '1' {
//             if car {
//                 return '1', true
//             } else {
//                 return '0', true
//             }
//         } else {
//             if car {
//                 return '0', true
//             } else {
//                 return '1', false
//             }
//         }
//     } else {
//         if b == '1' {
//             if car {
//                 return '0', true
//             } else {
//                 return '1', false
//             }
//         } else {
//             if car {
//                 return '1', false
//             } else {
//                 return '0', false
//             }
//         }
//     }
//     panic("sum: should not happen")
// }
