package pdftemplating

import (
	"log"

	"github.com/go-pdf/fpdf"
	"github.com/go-pdf/fpdf/contrib/gofpdi"
)

type TemplateGenerator struct {
	Orientation          Orientation
	Template             string
	OverrideReplacements map[string]string
	Replacements         map[string]Replacement
	Debug                bool
}

// logPrintf only prints log entries when debugging is enabled
func (g TemplateGenerator) logPrintf(s string, a ...any) {
	if g.Debug {
		log.Printf(s, a...)
	}
}

// Generate generates a single PDF file with the specified replacements
func (g *TemplateGenerator) Generate(fn string, repl map[string]string) error {
	g.logPrintf("Generate()")
	pdf := fpdf.New(g.Orientation.String(), "pt", "Letter", "")

	imp := gofpdi.NewImporter()
	tpl := imp.ImportPage(pdf, g.Template, 1, "/MediaBox")
	pageSizes := imp.GetPageSizes()

	// TODO: FIXME: Specify in configuration
	pdf.SetFontLocation("fonts")

	for _, r := range g.Replacements {
		g.logPrintf("Add font %s : %s", r.FontFamily, r.FontJson)
		pdf.AddFont(r.FontFamily, "", r.FontJson)
	}

	g.logPrintf("Add page")
	pdf.AddPage()

	g.logPrintf("Use imported template")
	imp.UseImportedTemplate(
		pdf,
		tpl,
		0, 0,
		pageSizes[0]["/MediaBox"]["w"],
		pageSizes[0]["/MediaBox"]["h"],
	)

	g.logPrintf("Iterating through replacements")
	//g.logPrintf("repl: %#v", repl)

	for k, v := range g.Replacements {
		g.logPrintf("Replacement %s : %#v", k, v)
		text, ok := g.OverrideReplacements[k]
		if !ok {
			text, ok = repl[k]
			if !ok {
				g.logPrintf("WARN: no value for %s, skipping", k)
			}
		}
		//g.logPrintf("Replacement %s : %s", k, text)
		pdf.SetFont(v.FontFamily, "", v.FontSize)
		wd := pdf.GetStringWidth(text) + 6

		switch v.Alignment {
		case AlignmentCenter:
			if v.PosX == 0 {
				pdf.SetX((pageSizes[0]["/MediaBox"]["w"] - wd) / 2)
			} else {
				pdf.SetX((v.PosX - wd) / 2)
			}
		case AlignmentLeft:
			pdf.SetX(v.PosX)
		case AlignmentRight:
			pdf.SetX(v.PosX - wd)
		default:
		}

		pdf.SetY(v.PosY)
		pdf.WriteAligned(0, v.FontSize, text, v.Alignment.String())
	}

	return pdf.OutputFileAndClose(fn)
}
