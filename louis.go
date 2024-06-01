package louis

import "louis/interfaces"

type SupportedLanguage int

type Louis struct {
	Transpilers map[SupportedLanguage]interfaces.Transpiler
}
