package engine

import (
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/proto/urls"
)

type DatabaseUrls struct {
}

func (d *DatabaseUrls) Create(url string, interval int) int {
	panic("implement me")
}

func (d *DatabaseUrls) Delete(id int) {
	panic("implement me")
}

func (d *DatabaseUrls) Get() []*urls.Url {
	panic("implement me")
}

func (d *DatabaseUrls) History(id int) []*urls.Response {
	panic("implement me")
}
