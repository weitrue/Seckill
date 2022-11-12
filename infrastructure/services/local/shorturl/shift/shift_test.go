package shift

import (
	"testing"
)

type Case struct {
	Str    string
	Expect string
}

var cases []Case

func TestGenerator(t *testing.T) {
	cases = []Case{
		{
			Str:    "https://jira.33.cn/browse/SYCZ-237?jql=project%20%3D%20SYCZ%20AND%20text%20~%20%E5%90%8E%E7%AB%AF%20AND%20assignee%20in%20(currentUser())%20ORDER%20BY%20priority%20DESC%2C%20updated%20DESC",
			Expect: "",
		},
		{
			Str:    "https://jira.33.cn/browse/SYCZ-237?jql=project%20%3D%20SYCZ%20AND%20text%20~%20%E5%90%8E%E7%AB%AF%20AND%20assignee%20in%20(currentUser())%20ORDER%20BY%20priority%20DESC%2C%20updated%20DESD",
			Expect: "",
		},
		{
			Str:    "https://jira.33.cn/browse/SYCZ-237?jql=project%20%3D%20SYCZ%20AND%20text%20~%20%E5%90%8E%E7%AB%AF%20AND%20assignee%20in%20(currentUser())%20ORDER%20BY%20priority%20DESC%2C%20updated%20DESE",
			Expect: "",
		},
		{
			Str:    "https://jira.33.cn/browse/SYCZ-237?jql=project%20%3D%20SYCZ%20AND%20text%20~%20%E5%90%8E%E7%AB%AF%20AND%20assignee%20in%20(currentUser())%20ORDER%20BY%20priority%20DESC%2C%20updated%20DESF",
			Expect: "",
		},
	}

	base := NewShortUrlBase(12, 6, "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for _, c := range cases {
		res := base.Generator(c.Str)
		t.Log(res)
	}
}
