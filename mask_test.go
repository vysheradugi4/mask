package mask

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestMask(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("CreateMask testing", func() {

		g.It("Should return empty string and error if code is empty string", func() {
			code := ""
			key := "winter"
			out, err := CreateMask(code, key)
			Expect(out).Should(BeEmpty())
			Expect(err.Error()).Should(Equal("Code string should be not empty"))
		})

		g.It("Should return empty string and error if code shorter than the key", func() {
			code := "win"
			key := "winter"
			out, err := CreateMask(code, key)
			Expect(out).Should(BeEmpty())
			Expect(err.Error()).Should(Equal("Code string should be not shorter than the key"))
		})

		g.It("Should return key string in uppercase cause length key == code", func() {
			code := "WINTER"
			key := "winter"
			out, err := CreateMask(code, key)
			Expect(out).Should(Equal("KEY"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return key string in uppercase for first character cause length key == code", func() {
			code := "Winter"
			key := "winter"
			out, err := CreateMask(code, key)
			Expect(out).Should(Equal("Key"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return key string in lowercase cause length key == code", func() {
			code := "winter"
			key := "winter"
			out, err := CreateMask(code, key)
			Expect(out).Should(Equal("key"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return empty key string and error cause code not contain key", func() {
			code := "autumn"
			key := "sum"
			out, err := CreateMask(code, key)
			Expect(out).Should(Equal(""))
			Expect(err.Error()).Should(Equal("Code must contain the key"))
		})

		g.It("Should return ?", func() {
			code := "autumn55AA"
			key := "Autumn"
			out, err := CreateMask(code, key)
			Expect(out).Should(Equal("key\\d\\d[A-Z][A-Z]"))
			Expect(err).Should(BeNil())
		})
	})

	g.Describe("changeToKey testing", func() {
		g.It("Should return flag lowercase for code: winter", func() {
			code := "winter"
			key := "winter"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal("key"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return flag Capitalize for code: Winter", func() {
			code := "Winter"
			key := "winter"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal("Key"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return flag UPPERCASE for code: WINTER", func() {
			code := "WINTER"
			key := "winter"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal("KEY"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return flag camelCase for code: wINTER", func() {
			code := "wINTER"
			key := "winter"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal("kEy"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return flag camelCase for code: WInTER", func() {
			code := "WInTER"
			key := "winter"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal("kEy"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return flag camelCase for code: WinTER", func() {
			code := "WinTER"
			key := "winter"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal("kEy"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return flag camelCase for code: winTer", func() {
			code := "winTer"
			key := "winter"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal("kEy"))
			Expect(err).Should(BeNil())
		})

		g.It("Should return error cause code not contain key", func() {
			code := "winTer"
			key := "summer"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal(""))
			Expect(err.Error()).Should(Equal("Code must contain key"))
		})

		g.It("Should return error cause code not contain key", func() {
			code := "winTer"
			key := "summer"
			flag, err := changeToKey(code, key)
			Expect(flag).Should(Equal(""))
			Expect(err.Error()).Should(Equal("Code must contain key"))
		})
	})

	g.Describe("containsFold testing", func() {
		g.It("Should return true cause code contain key", func() {
			code := "autumn55AA"
			key := "autUmn"
			flag := containsFold(code, key)
			Expect(flag).Should(BeTrue())
		})

		g.It("Should return false cause code is empty", func() {
			code := ""
			key := "autUmn"
			flag := containsFold(code, key)
			Expect(flag).Should(BeFalse())
		})

		g.It("Should return false cause key is empty", func() {
			code := "asdfasdf"
			key := ""
			flag := containsFold(code, key)
			Expect(flag).Should(BeFalse())
		})
	})

	g.Describe("changeChar testing", func() {
		g.It("Should return mask key55[A-Z][A-Z]", func() {
			code := "key55AA"
			flag := changeChar(code)
			Expect(flag).Should(Equal("key\\d\\d[A-Z][A-Z]"))
		})
	})

	g.Describe("Function generateCodesFromMask", func() {

		g.It("Should return slise of codes like summer01", func() {

			mask := "key\\d\\d"
			keyword := "summer"

			result := GenerateCodesFromMask(mask, keyword, 10)

			g.Assert("summer00").Equal(result[0])
			g.Assert("summer01").Equal(result[1])
			g.Assert("summer99").Equal(result[99])
		})

		g.It("Should return slise of codes like 01winter", func() {

			mask := "\\d\\dkey"
			keyword := "winter"

			result := GenerateCodesFromMask(mask, keyword, 2)

			g.Assert("00winter").Equal(result[0])
			g.Assert("01winter").Equal(result[1])
			g.Assert("99winter").Equal(result[99])
		})

		g.It("Should return slise of codes like winter1", func() {

			mask := "key\\d"
			keyword := "winter"

			result := GenerateCodesFromMask(mask, keyword, 1)

			g.Assert("winter0").Equal(result[0])
			g.Assert("winter1").Equal(result[1])
			g.Assert("winter9").Equal(result[9])
		})

		g.It("Should return slise of codes like winter111 with 3 digits", func() {

			mask := "key\\d\\d\\d"
			keyword := "winter"

			result := GenerateCodesFromMask(mask, keyword, 3)

			g.Assert("winter000").Equal(result[000])
			g.Assert("winter100").Equal(result[100])
			g.Assert("winter999").Equal(result[999])
		})

		g.It("Should return slise of codes like winter1111 with 4 digits", func() {

			mask := "KEY\\d\\d\\d\\d"
			keyword := "winter"

			result := GenerateCodesFromMask(mask, keyword, 4)

			g.Assert("WINTER0000").Equal(result[0])
			g.Assert("WINTER1111").Equal(result[1111])
			g.Assert("WINTER9999").Equal(result[9999])
		})

		g.It("Should return slise of codes with uppercase like WINTER1", func() {

			mask := "KEY\\d"
			keyword := "winter"

			result := GenerateCodesFromMask(mask, keyword, 1)

			g.Assert("WINTER0").Equal(result[0])
			g.Assert("WINTER1").Equal(result[1])
			g.Assert("WINTER9").Equal(result[9])
		})

		g.It("Should return slise of codes with first character uppercase like Winter1", func() {

			mask := "Key\\d"
			keyword := "winter"

			result := GenerateCodesFromMask(mask, keyword, 1)

			g.Assert("Winter0").Equal(result[0])
			g.Assert("Winter1").Equal(result[1])
			g.Assert("Winter9").Equal(result[9])
		})

		g.It("Should return empty slise of codes because count of iterations < needed substitution", func() {

			mask := "Key\\d\\d\\d\\d"
			keyword := "winter"

			result := GenerateCodesFromMask(mask, keyword, 3)

			Expect(result).Should(BeNil())
		})
	})
}
