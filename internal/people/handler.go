package people

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/make-go-great/date-go"
	"github.com/make-go-great/ioe-go"
)

type Handler interface {
	List(ctx context.Context) error
	Add(ctx context.Context) error
	Update(ctx context.Context) error
	Remove(ctx context.Context) error
	Export(ctx context.Context) error
	Import(ctx context.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) List(ctx context.Context) error {
	people, err := h.service.List(ctx)
	if err != nil {
		return err
	}

	for i, person := range people {
		fmt.Printf("%d: %s\n", i, person.Name)
		fmt.Println(person.Pretty("\t"))
	}

	return nil
}

func (h *handler) Add(ctx context.Context) error {
	person := Person{}

	fmt.Printf("Input name: ")
	person.Name = ioe.ReadInput()

	fmt.Printf("Input birthday (example %s): ", date.SupportDateFormats())
	person.Birthday = ioe.ReadInputEmpty()

	// TODO: check valid phone
	fmt.Printf("Input phone: ")
	person.Phone = ioe.ReadInputEmpty()

	fmt.Printf("Input University: ")
	person.University = ioe.ReadInputEmpty()

	fmt.Printf("Input VNCMND: ")
	person.VNCMND = ioe.ReadInputEmpty()

	fmt.Printf("Input VNCCCD: ")
	person.VNCCCD = ioe.ReadInputEmpty()

	fmt.Printf("Input VNMST: ")
	person.VNMST = ioe.ReadInputEmpty()

	fmt.Printf("Input VNBHXH: ")
	person.VNBHXH = ioe.ReadInputEmpty()

	fmt.Printf("Input CompanyVNG: ")
	person.CompanyVNG = ioe.ReadInputEmpty()

	fmt.Printf("Input SocialFacebook: ")
	person.SocialFacebook = ioe.ReadInputEmpty()

	fmt.Printf("Input SocialInstagram: ")
	person.SocialInstagram = ioe.ReadInputEmpty()

	fmt.Printf("Input SocialTiktok: ")
	person.SocialTiktok = ioe.ReadInputEmpty()

	fmt.Printf("Input SocialLinkedin: ")
	person.SocialLinkedin = ioe.ReadInputEmpty()

	return h.service.Add(ctx, &person)
}

func (h *handler) Update(ctx context.Context) error {
	fmt.Printf("Input ID: ")
	id := ioe.ReadInput()

	person, err := h.service.Get(ctx, id)
	if err != nil {
		return err
	}

	fmt.Println("!!! Input empty to keep old value")
	var val string

	fmt.Printf("Update name, current is [%s]: ", person.Name)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Name = val
	}

	fmt.Printf("Update birthday, current is [%s]: ", person.Birthday)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Birthday = val
	}

	fmt.Printf("Update phone, current is [%s]: ", person.Phone)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.Phone = val
	}

	fmt.Printf("Input University, current is [%s]: ", person.University)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.University = val
	}

	fmt.Printf("Input VNCMND, current is: [%s]", person.VNCMND)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.VNCMND = val
	}

	fmt.Printf("Input VNMST, current is [%s]: ", person.VNMST)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.VNMST = val
	}

	fmt.Printf("Input VNBHXH, current is [%s]: ", person.VNBHXH)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.VNBHXH = val
	}

	fmt.Printf("Input CompanyVNG, current is [%s]: ", person.CompanyVNG)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.CompanyVNG = val
	}

	fmt.Printf("Input SocialFacebook, current is [%s]: ", person.SocialFacebook)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.SocialFacebook = val
	}

	fmt.Printf("Input SocialInstagram, current is [%s]: ", person.SocialInstagram)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.SocialInstagram = ioe.ReadInputEmpty()
	}

	fmt.Printf("Input SocialTiktok, current is [%s]: ", person.SocialTiktok)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.SocialTiktok = val
	}

	fmt.Printf("Input SocialLinkedin, current is [%s]: ", person.SocialLinkedin)
	val = ioe.ReadInputEmpty()
	if val != "" {
		person.SocialLinkedin = val
	}

	return h.service.Update(ctx, person)
}

func (h *handler) Remove(ctx context.Context) error {
	fmt.Printf("Input ID: ")
	id := ioe.ReadInput()

	return h.service.Remove(ctx, id)
}

func (h *handler) Export(ctx context.Context) error {
	fmt.Printf("Input filename: ")
	filename := ioe.ReadInput()

	people, err := h.service.List(ctx)
	if err != nil {
		return fmt.Errorf("service: failed to list: %w", err)
	}

	data := WrapPeople{
		People: people,
	}

	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("json: failed to marshal indent: %w", err)
	}

	if err := os.WriteFile(filename, bytes, 0o600); err != nil {
		return fmt.Errorf("os: failed to write file: %w", err)
	}

	return nil
}

func (h *handler) Import(ctx context.Context) error {
	fmt.Printf("Input filename: ")
	filename := ioe.ReadInput()

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("os: failed to read file: %w", err)
	}

	data := WrapPeople{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return fmt.Errorf("json: failed to unmarshal: %w", err)
	}

	for _, person := range data.People {
		if err := h.service.Add(ctx, person); err != nil {
			return fmt.Errorf("service: failed to add: %w", err)
		}
	}

	return nil
}
