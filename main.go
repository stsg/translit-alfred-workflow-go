package main

import (
	"log"
	"os/exec"
	"strings"

	//	"os"
	"os"

	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// Icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}
	redditIcon      = &aw.Icon{Value: "icons/reddit.png"}
	githubIcon      = &aw.Icon{Value: "icons/github.png"}
	translateIcon   = &aw.Icon{Value: "icons/translate.png"}
	forumsIcon      = &aw.Icon{Value: "icons/forums.png"}
	stackIcon       = &aw.Icon{Value: "icons/stack.png"}
	docIcon         = &aw.Icon{Value: "icons/doc.png"}

	query string

	repo = "stsg/translit-alfred-workflow-go"

	// Workflow stuff
	wf       *aw.Workflow
	startDir string // Directory to read
)

func init() {
	startDir = os.Getenv("HOME")
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues")) //Initialise workflow
}

// run executes the Script Filter.
func run() {
	// ----------------------------------------------------------------
	// Handle CLI arguments
	// ----------------------------------------------------------------

	// You should always use wf.Args() in Script Filters. It contains the
	// same as os.Args[1:], but the arguments are first parsed for AwGo's
	// magic actions (i.e. `workflow:*` to allow the user to easily open
	// the log or data/cache directory).
	if args := wf.Args(); len(args) > 0 {
		// If you're using "{query}" or "$1" (with quotes) in your
		// Script Filter, $1 will always be set, even if to an empty
		// string.
		// This guard serves mostly to prevent errors when run on
		// the command line.
		query = args[0]

		log.Printf("query=%s", query)

		issw := os.Getenv("HOME") + "/bin/issw"

		out, err := exec.Command(issw).Output()
		if err != nil {
			log.Printf("err: %s", err)
		}
		cuKl := strings.TrimSpace(string(out))
		if cuKl == enKl {
			_, err := exec.Command(issw, ruKl).Output()
			if err != nil {
				log.Printf("err: %s", err)
			}
		} else {
			_, err := exec.Command(issw, enKl).Output()
			if err != nil {
				log.Printf("err: %s", err)
			}
		}

		wf.NewItem(tr(string(query)))
		wf.SendFeedback()
	}
}

func main() {
	// Call workflow via `Run` wrapper to catch any errors, log them
	// and display an error message in Alfred.
	wf.Run(run)
}
