//go:build !windows

package print

import (
	"fmt"
	"testing"
	"time"

	"github.com/OpenPrinting/goipp"
	"github.com/stretchr/testify/assert"
)

func TestCreateGroupsFromMessage(t *testing.T) {
	_, _ = getResponseGroups(goipp.OpCupsGetPrinters, localCupsURI, "all")
	msg := createTestPrintersResponse()
	groups, err := createGroupsFromMessage(msg)
	assert.Nil(t, err)
	assert.Equal(t, goipp.TagOperationGroup, (*groups)[0].Tag)
	assert.Equal(t, goipp.TagPrinterGroup, (*groups)[1].Tag)
}

// This test assumes that there is no server on port 632.
func TestGetResponseGroups_BadURI(t *testing.T) {
	groups, err := getResponseGroups(goipp.OpCupsGetPrinters,
		"http://localhost:632", "all")
	assert.Equal(t, 0, len(*groups))
	assert.NotNil(t, err)
	fmt.Println(err.Error())

}

func createTestPrintersResponse() *goipp.Message {
	now := time.Now().UTC()
	m := goipp.NewResponse(goipp.DefaultVersion, 200, 1)
	operationAttrs := goipp.Attributes{}
	operationAttrs.Add(goipp.MakeAttribute("attributes-charset",
		goipp.TagCharset, goipp.String("utf-8")))
	operationAttrs.Add(goipp.MakeAttribute("attributes-natural-language",
		goipp.TagLanguage, goipp.String("en-us")))
	opAttributesGroup := goipp.Group{Tag: goipp.TagOperationGroup, Attrs: operationAttrs}
	m.Groups.Add(opAttributesGroup)

	prAttrs := goipp.Attributes{}
	prAttrs.Add(goipp.MakeAttribute("marker-change-time",
		goipp.TagInteger, goipp.Integer(0)))
	prAttrs.Add(goipp.MakeAttribute("printer-config-change-date-time",
		goipp.TagDateTime, goipp.Time{now}))
	prAttrs.Add(goipp.MakeAttribute("printer-config-change-time",
		goipp.TagInteger, goipp.Integer(1739223365)))
	prAttrs.Add(goipp.MakeAttribute("printer-dns-sd-name",
		goipp.TagNoValue, goipp.Void{}))
	prAttrs.Add(goipp.MakeAttribute("printer-error-policy",
		goipp.TagName, goipp.String("retry-job")))
	prAttrs.Add(goipp.MakeAttribute("printer-error-policy-supported",
		goipp.TagName, goipp.String("abort-job retry-current-job retry-job stop-printer")))
	prAttrs.Add(goipp.MakeAttribute("printer-icons",
		goipp.TagURI, goipp.String("http://localhost:631/icons/Printer1.png")))
	prAttrs.Add(goipp.MakeAttribute("printer-is-accepting-jobs",
		goipp.TagBoolean, goipp.Boolean(true)))
	prAttrs.Add(goipp.MakeAttribute("printer-is-shared",
		goipp.TagBoolean, goipp.Boolean(false)))
	prAttrs.Add(goipp.MakeAttribute("printer-is-temporary",
		goipp.TagBoolean, goipp.Boolean(false)))
	prAttrs.Add(goipp.MakeAttribute("printer-more-info",
		goipp.TagURI, goipp.String("http://localhost:631/printers/Printer1")))
	prAttrs.Add(goipp.MakeAttribute("printer-op-policy",
		goipp.TagName, goipp.String("default")))
	prAttrs.Add(goipp.MakeAttribute("printer-state",
		goipp.TagEnum, goipp.Integer(3)))
	prAttrs.Add(goipp.MakeAttribute("printer-state-change-date-time",
		goipp.TagDateTime, goipp.Time{now}))
	prAttrs.Add(goipp.MakeAttribute("printer-change-time",
		goipp.TagInteger, goipp.Integer(1739223363)))
	prAttrs.Add(goipp.MakeAttribute("printer-name",
		goipp.TagName, goipp.String("Printer1")))

	prAttributesGroup := goipp.Group{Tag: goipp.TagPrinterGroup, Attrs: prAttrs}
	m.Groups.Add(prAttributesGroup)

	pr1Attrs := goipp.Attributes{}
	pr1Attrs.Add(goipp.MakeAttribute("marker-change-time",
		goipp.TagInteger, goipp.Integer(0)))
	pr1Attrs.Add(goipp.MakeAttribute("printer-config-change-date-time",
		goipp.TagDateTime, goipp.Time{now}))
	pr1Attrs.Add(goipp.MakeAttribute("printer-config-change-time",
		goipp.TagInteger, goipp.Integer(1739223365)))
	pr1Attrs.Add(goipp.MakeAttribute("printer-dns-sd-name",
		goipp.TagNoValue, goipp.Void{}))
	pr1Attrs.Add(goipp.MakeAttribute("printer-error-policy",
		goipp.TagName, goipp.String("retry-job")))
	pr1Attrs.Add(goipp.MakeAttribute("printer-error-policy-supported",
		goipp.TagName, goipp.String("abort-job retry-current-job retry-job stop-printer")))
	pr1Attrs.Add(goipp.MakeAttribute("printer-icons",
		goipp.TagURI, goipp.String("http://localhost:631/icons/Printer2.png")))
	pr1Attrs.Add(goipp.MakeAttribute("printer-is-accepting-jobs",
		goipp.TagBoolean, goipp.Boolean(true)))
	pr1Attrs.Add(goipp.MakeAttribute("printer-is-shared",
		goipp.TagBoolean, goipp.Boolean(false)))
	pr1Attrs.Add(goipp.MakeAttribute("printer-is-temporary",
		goipp.TagBoolean, goipp.Boolean(false)))
	pr1Attrs.Add(goipp.MakeAttribute("printer-more-info",
		goipp.TagURI, goipp.String("http://localhost:631/printers/Printer2")))
	pr1Attrs.Add(goipp.MakeAttribute("printer-op-policy",
		goipp.TagName, goipp.String("default")))
	pr1Attrs.Add(goipp.MakeAttribute("printer-state",
		goipp.TagEnum, goipp.Integer(3)))
	pr1Attrs.Add(goipp.MakeAttribute("printer-state-change-date-time",
		goipp.TagDateTime, goipp.Time{now}))
	pr1Attrs.Add(goipp.MakeAttribute("printer-change-time",
		goipp.TagInteger, goipp.Integer(1739223363)))
	pr1Attrs.Add(goipp.MakeAttribute("printer-name",
		goipp.TagName, goipp.String("Printer2")))

	pr1AttributesGroup := goipp.Group{Tag: goipp.TagPrinterGroup, Attrs: pr1Attrs}
	m.Groups.Add(pr1AttributesGroup)

	return m
}
