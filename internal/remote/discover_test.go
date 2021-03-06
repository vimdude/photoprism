package remote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscover(t *testing.T) {
	t.Run("webdav", func(t *testing.T) {
		r, err := Discover("http://admin:photoprism@webdav-dummy/", "", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Webdav", r.AccName)
		assert.Equal(t, "webdav", r.AccType)
		assert.Equal(t, "http://webdav-dummy/", r.AccURL)
		assert.Equal(t, "admin", r.AccUser)
		assert.Equal(t, "photoprism", r.AccPass)
	})

	t.Run("webdav password", func(t *testing.T) {
		r, err := Discover("http://admin@webdav-dummy/", "", "photoprism")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Webdav", r.AccName)
		assert.Equal(t, "webdav", r.AccType)
		assert.Equal(t, "http://webdav-dummy/", r.AccURL)
		assert.Equal(t, "admin", r.AccUser)
		assert.Equal(t, "photoprism", r.AccPass)
	})

	t.Run("facebook", func(t *testing.T) {
		r, err := Discover("https://www.facebook.com/ob.boris.palmer", "", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Facebook", r.AccName)
		assert.Equal(t, "facebook", r.AccType)
		assert.Equal(t, "https://www.facebook.com/ob.boris.palmer", r.AccURL)
		assert.Equal(t, "", r.AccUser)
		assert.Equal(t, "", r.AccPass)
	})
}
