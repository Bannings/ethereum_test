package sig

import "testing"

func TestVerifySHA256RSA(t *testing.T) {
	const pubPEM = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDjdsshaaj/v7ho9ECXlvTZba2E
5GcbplT3AchiATlcgXr6WdYJvgZRRkv4/QdJanvWitsxgSphzT3yyFToqKOIS2js
9cPc7yj1hB1v+lu4WwWbxLwdHEejKZqYbLIDaqruxNZ7/QZFJBpd8ZNif2gRLqfR
gLmtXNR/vkj2IQeRcQIDAQAB
-----END PUBLIC KEY-----`

	const (
		message = "lalalalala"
		sig     = "MwKQnq2Syy6dYTuWXK3CYRGE/afc+Wi+1+koRXdLROgYj7XF8ZJQ5dEXdgvyOgA0TUozssWXk9s0pPMGfHd9PvYtou5z92wgo5AwD5PCDIC6+OWxG7BeKz/LBrVS8EejbKD9hc84cwlg1+dIxr2xOjjHHwlEwoYqjEL4Fgy+aDA="
	)

	if err := VerifySHA256RSA(pubPEM, []byte(message), sig); err != nil {
		t.Error(err)
	}

}
