package okr

import "fmt"

var (
	// set convention of format results
	convention = `O: This is objective
	KR1: This is key result
	KR2: This is key result
	KR3: This is key result
	KR4: 
	`
)

// ContentGenerator(param string) generate the content for system
func ContentGenerator(language string) string {
	var result string
	content := `Please read this https://about.gitlab.com/company/okrs/ as a reference. 
	Please translate this objective into format OKR's`

	result = fmt.Sprintf(`
	%s.
	Please translate to %s.
	Please follow this format %s as the result.
	`,
		content,
		language,
		convention,
	)

	return result
}
