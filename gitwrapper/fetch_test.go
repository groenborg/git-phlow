package gitwrapper

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/praqma/git-phlow/testfixture"
)

func TestFetch_Fetch(t *testing.T) {

	Convey("Test fetch from origin", t, func() {

		testfixture.SetupTestRepo()

		Convey("fetch all should fail", func() {
			git := InitGit()
			_, err := git.Fetch().FetchFromOrigin()
			So(err, ShouldNotBeNil)
		})

		testfixture.TearDownTestRepo()
	})
}

func TestFetch_HasRemote(t *testing.T) {

	Convey("Test has remote method", t, func() {

		testfixture.SetupTestRepo()

		Convey("fetch shold return false for no origin", func() {
			hasRemote := InitGit().Fetch().HasRemote()

			So(hasRemote, ShouldBeFalse)
		})

		testfixture.TearDownTestRepo()
	})
}
