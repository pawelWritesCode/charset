// Package charset contains stuff related to charset
// Each exported charset usually does not contain all symbols in it.
package charset

import (
	"math/rand"
	"time"
)

const (
	// Unicode represents set of Unicode characters (contain multi byte runes).
	Unicode = Emoji + Sex + Keyboard + Currencies + Greek +
		MathematicalSymbols + MathematicalFonts + InvertedLetters + Braille + IPA + FullWidthCharacters + Units +
		Latin + Cyrillic + Chinese + Japanese + Korean + Arabic + Ethiopian + Devanagari + Bengali + Tamil + Tibetan +
		Phoenician + Runes

	// ASCII represents set containing only ASCII characters.
	ASCII = " !#$%&()*+,-.0123456789:;=?@ABCDEFGHIJKLMNOPQRSTUVWXYZ^_abcdefghijklmnopqrstuvwxyz|~"

	// Polish represents set of only polish letters.
	Polish = "ĄąĆćĘęŁłŃńÓóŚśŹźŻżabcdefghijklmnoprstuwvxyzABCDEFGHIJKLMNOPRSTUWVXYZ"

	// English represents set of only english letters.
	English = "abcdefghijklmnoprstuwvxyzABCDEFGHIJKLMNOPRSTUWVXYZ"

	// Russian represents set of only russian letters.
	Russian = "АаБбВвГгДдЕеЁёЖжЗзИиЙйКкЛлМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЪъЫыЬьЭэЮюЯя"

	// MathematicalSymbols represents set of mathematical symbols.
	MathematicalSymbols = "∛ℤℝ⋃⋂⋡⪺≠⊐∵∧⦝⊕∯∑∏⊶⫝̸⦕⟪⦅」⟅⦋"

	// MathematicalFonts represents set of different letters and numbers in different mathematical fonts.
	MathematicalFonts = "𝔄𝔞𝕭𝖇𝔻𝕕𝟛𝐅𝐟𝐿𝑙𝑴𝒎𝒩𝓃𝓞𝓸𝖶𝗐𝗫𝘅𝟭𝘠𝘺𝚉𝚣"

	// Emoji represents set of emojis
	Emoji = "🤡🤖🧟🏋🥇☟🚲🚠🐞🐜💄🐲🌓🌪🇵🇱💵"

	// Currencies represents set of financial currencies.
	Currencies = "₿£$¢₰₱₣₳฿₲₭₥₦₱₽₴₮₩"

	// Sex represents set of symbols related with genders.
	Sex = "⚥⚩⚮⚭⚣⚢⚤⚯"

	// Keyboard represents set of keyboard symbols.
	Keyboard = "⌘⎈✄↵✧✲⎇✦✧"

	// Greek represents set of greek characters.
	Greek = "αβδεθλμπφψΩ𝞕𝞖𝞗𝞘𝝽𝝾𝝿𝞀𝞂"

	// InvertedLetters represents set of inverted letters.
	InvertedLetters = "Z⅄XʍɅ∩ꞱSᴚÒԀONƜꞀꞰſIH⅁ℲƎDƆBⱯzʎxʍʌnʇsɹbdouɯꞁʞfᴉɥᵷɟǝpɔqɐ"

	// IPA International Phonetic Alphabet (IPA) is a phonetic (pronunciation) notation designed for all human languages.
	IPA = "āäēĕæɛiːɝo͞o"

	// FullWidthCharacters A full-width character means the character has the same width as a Chinese character,
	// regardless of font choice.
	FullWidthCharacters = "０ＣＤ￦￤ｒｐａｒｔｉｃｕｌａｒ"

	// Units represents set of unit symbols
	Units = "㎛㎠㎢㎖㎲㎏㎆㎑㎷㏀㎃㎨㎮㎪㎉㏑㍱㏖"

	// Braille is used by blind people
	Braille = "⠞⠽⠏⠑⠓⠑⠗⠑"

	// Latin Alphabet (aka Roman alphabet) is used to write most Western languages, including:
	// English, German, French, Spanish, Italian.
	Latin = "ᴁèóƃᶐɷȷᴂɒᴝᴥ"

	// Cyrillic The Cyrillic script is a writing system used for various alphabets used in Eastern Europe,
	// the Caucasus, Central Asia, and North Asia.
	Cyrillic = Russian + "ԆҤ҂ԔԙӜԨԬӦӴѠѼ"

	// Chinese There are many thousands of Chinese characters. A college educated Chinese person typically know
	// 5 thousand character. 3 thousand is minimum to read newspapers.
	Chinese = "的一是在不了有和人这中大为上个国相见欢·林花谢了春红"

	// Japanese represents set of Japanese characters.
	Japanese = "あいうえアイウエオ・ヽヾヿｱｲｳｴｵｶｷｸ"

	// Korean represents set of Korean characters.
	Korean = "ᅒᅓᅔᅕᅖᅗᅘᆪᆫᆬᆭᆮᆯᆰ"

	// Arabic represents set of Arabic characters.
	Arabic = "ﷺ﷽ﵴﴙﲀ۝ﲂ۞؊٩١۲ݶمڮجݗݨݳھڝ"

	// Ethiopian represents set of Ethiopian characters. It is third in Africa, followed by West African Cats and East Africa.
	Ethiopian = "ቷቸቹጟጠኸኹⶹꬨꬩꬖꬠ᎐᎑᎒፣፤፥፪፫፬የኢትዮጵያ፡መ"

	// Devanagari represents set of Indian characters. The most popular language of the Indo-Aryan family is Hindi,
	// and Hindi uses Devanagari as its writing system
	Devanagari = "ऒओॺॻ॥॰षस१२३ळऴ"

	// Bengali represents set of Bengali characters. Bengali script is mainly used by languages Bengali and Assamese, in India.
	Bengali = "ঊঋজঝঞলশ৪৫৵৶ঁংৗৢ"

	// Tamil represents set of Tamil characters. It is an abugida script that is used by Tamils and Tamil speakers
	// in India, Sri Lanka, Malaysia, Singapore, Indonesia.
	Tamil = "ஃஅஆஇணதந௫௬௭ௐ௳௴ோ ௌ"

	// Tibetan represents set of Tibetan characters.
	Tibetan = "གྷངཅ༳༪༫༐༑༒࿄࿅࿇࿈༚༛༜࿐࿑࿒ ཱ ི ཱི༻༼ ༽ྨ ྩ"

	// Phoenician represents set of Phoenican characters. Phoenician is the oldest alphabet. It began around 1200 BC.
	Phoenician = "𐤟𐤛𐤗𐤘𐤒𐤓𐤔𐤕"

	// Runes represents set of old runes. Rune alphabets was used to write various
	// Germanic languages before the adoption of the Latin alphabet.
	Runes = "ᚠᚡᚢᛋᛌᛍ"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// RandomRunes returns slice of random runes of given length.
// Argument length indices length of output slice.
// Argument charset indices input charset from which output slice will be composed.
func RandomRunes(length int, charset []rune) []rune {
	output := make([]rune, 0, length)

	for i := 0; i < length; i++ {
		output = append(output, charset[seededRand.Intn(len(charset))])
	}

	return output
}
