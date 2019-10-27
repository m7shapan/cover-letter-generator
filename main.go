package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	data := make(map[string]string)
	missingData := []string{
		"Company Name",
		"Job Role",
		"Website",
		"Email or Form",
	}

	for index := 0; index < len(missingData); index++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("%s ? ", missingData[index])
		scanner.Scan()

		data[missingData[index]] = scanner.Text()
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	pdf.Write(10, fmt.Sprintf("Dear %s,", data[missingData[0]]))
	pdf.Ln(10)
	pdf.Write(10, "To whom it may concern,")
	pdf.Ln(10)
	pdf.Write(5, fmt.Sprintf(`I recently came across your job posting for the %s on %s and I am very excited to apply. After reading the job description and requirements and matching it with my own experiences, I am certain that I would be a valuable asset to your company.`, data[missingData[1]], data[missingData[2]]))
	pdf.Ln(10)
	pdf.Write(5, fmt.Sprintf(`You will find both my resume and my LinkedIn profile attached to this %s. It would be great to hear back from you and talk with you in more detail.`, data[missingData[3]]))
	pdf.Ln(10)
	pdf.Write(10, "Thank you for your time.")
	pdf.Ln(10)
	pdf.Write(10, "Sincerely,")
	pdf.Ln(5)
	pdf.Write(10, "Mohamed Shapan")
	pdf.Ln(5)
	pdf.Write(10, "Backend Engineer")
	pdf.Ln(5)
	pdf.Write(10, "m7shapan@gmail.com")
	pdf.Ln(5)
	pdf.Write(10, "+201145989317")
	pdf.Ln(5)
	pdf.WriteLinkString(10, "Linkedin", "https://www.linkedin.com/in/m7shapan/")

	err := pdf.OutputFileAndClose("Mohamed_Shapan_Cover_letter.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
