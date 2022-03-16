package people

import (
	"context"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/make-go-great/ioe-go"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(ctx context.Context) error {
	people, err := h.service.List(ctx)
	if err != nil {
		return err
	}

	// https://github.com/jedib0t/go-pretty/tree/main/table
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)
	tableWriter.AppendHeader(table.Row{
		"ID",
		"Name",
		"Birthday",
		"Phone",
		"CMND",
		"MST",
		"BHXH",
		"University",
		"VNG",
		"Facebook",
		"Instagram",
		"Tiktok",
	})

	for _, person := range people {
		tableWriter.AppendRow(table.Row{
			person.ID,
			person.Name,
			person.Birthday,
			person.Phone,
			person.CMND,
			person.MST,
			person.BHXH,
			person.University,
			person.VNG,
			person.Facebook,
			person.Instagram,
			person.Tiktok,
		})
	}

	tableWriter.Render()

	return nil
}

func (h *Handler) Add(ctx context.Context) error {
	person := Person{}

	fmt.Printf("Input name: ")
	person.Name = ioe.ReadInput()

	fmt.Printf("Input birthday (1998-01-18 or 1997/04/01, month before day): ")
	person.Birthday = ioe.ReadInputEmpty()

	// TODO: check valid phone
	fmt.Printf("Input phone: ")
	person.Phone = ioe.ReadInputEmpty()

	fmt.Printf("Input CMND: ")
	person.CMND = ioe.ReadInputEmpty()

	fmt.Printf("Input MST: ")
	person.MST = ioe.ReadInputEmpty()

	fmt.Printf("Input BHXH: ")
	person.BHXH = ioe.ReadInputEmpty()

	fmt.Printf("Input university: ")
	person.University = ioe.ReadInputEmpty()

	fmt.Printf("Input VNG: ")
	person.VNG = ioe.ReadInputEmpty()

	fmt.Printf("Input Facebook: ")
	person.Facebook = ioe.ReadInputEmpty()

	fmt.Printf("Input Instagram: ")
	person.Instagram = ioe.ReadInputEmpty()

	fmt.Printf("Input Tiktok: ")
	person.Tiktok = ioe.ReadInputEmpty()

	return h.service.Add(ctx, person)
}

func (h *Handler) Update(ctx context.Context) error {
	fmt.Printf("Input ID: ")
	id := ioe.ReadInput()

	person, err := h.service.Get(ctx, id)
	if err != nil {
		return err
	}

	fmt.Println("!!! Input empty to keep old value")
	var val string

	fmt.Printf("Update name, current is %s: ", person.Name)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Name = val
	}

	fmt.Printf("Update birthday, current is %s: ", person.Birthday)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Birthday = val
	}

	fmt.Printf("Update phone, current is %s: ", person.Phone)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Phone = val
	}

	fmt.Printf("Input CMND, current is: %s", person.CMND)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.CMND = val
	}

	fmt.Printf("Input MST, current is %s: ", person.MST)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.MST = val
	}

	fmt.Printf("Input BHXH, current is %s: ", person.BHXH)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.BHXH = val
	}

	fmt.Printf("Input university, current is %s: ", person.University)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.University = val
	}

	fmt.Printf("Input VNG, current is %s: ", person.VNG)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.VNG = val
	}

	fmt.Printf("Input Facebook, current is %s: ", person.Facebook)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Facebook = val
	}

	fmt.Printf("Input Instagram, current is %s: ", person.Instagram)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Instagram = ioe.ReadInputEmpty()
	}

	fmt.Printf("Input Tiktok, current is %s: ", person.Tiktok)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Tiktok = val
	}

	return h.service.Update(ctx, person)
}

func (h *Handler) Remove(ctx context.Context) error {
	fmt.Printf("Input ID: ")
	id := ioe.ReadInput()

	return h.service.Remove(ctx, id)
}