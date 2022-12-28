package landings

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type landings struct {
	rendered []byte
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	// ropts := opts.Domain.RendererOptions

	return &landings{
		rendered: []byte(`<h1>Todo</h1>`),
	}, nil
}

func (l *landings) ServeEditorFile(ctx *gin.Context, file string) error {
	return nil
}

func (l *landings) Handle(ctx httpx.Context) {
	ctx.Http.Writer.Write(l.rendered)

}

type templateOptions struct {
	BackgroudURL string
	Title        string
	TitleColor   string
	Info         string
	InfoColor    string
	Email        bool
	EmailCtaText string
	Socials      map[string]string
	SocialStype  string
	InjectCSS    string
	InjectJS     string
}
