package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nomionz/currency-converter/internal/html"
	"github.com/nomionz/currency-converter/pkg/ecb"
)

// Server is a http server with simple example of ecb client
// It responses with html, because it utilizes htmx for dynamic content update
type Server struct {
	cli  *ecb.Client
	view *html.ViewData
}

func NewServer() *Server {
	s := &Server{
		cli:  ecb.NewClient(NewMemStore()),
		view: html.NewViewData(),
	}

	go s.cli.Fetch()

	return s
}

func (s *Server) Router(app *fiber.App) {
	app.Get("/", s.Index)
	app.Post("/convert", s.Convert)
}

func (s *Server) Index(c *fiber.Ctx) error {
	s.view.ResetValues()
	s.view.SetCurrencies(s.cli.GetCurrencies())
	return c.Render("index", s.view, "layouts/main")
}

func (s *Server) Convert(c *fiber.Ctx) error {
	req := new(html.Request)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	tag := req.GetAmountTagName()

	if tag == html.AmountTo {
		tmp := req.From
		req.From = req.To
		req.To = tmp
	}

	amount, _ := s.convert(req.From, req.To, req.GetAmount())

	return s.renderInput(c, tag, amount)
}

func (s *Server) renderInput(c *fiber.Ctx, tag, amount string) error {
	switch tag {
	case html.AmountFrom:
		s.view.InputTo.Value = amount
		return c.Render("input", s.view.InputTo)
	case html.AmountTo:
		s.view.InputFrom.Value = amount
		return c.Render("input", s.view.InputFrom)
	default:
		return c.Status(fiber.StatusUnprocessableEntity).SendString("could not determine amount tag name to update")
	}
}

func (s *Server) convert(from, to, amount string) (string, error) {
	amountDec, err := s.cli.Convert(from, to, amount)

	if err != nil {
		return "", err
	}

	return amountDec.String(), nil
}
