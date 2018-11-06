package main

import (
	"log"

	aw "github.com/deanishe/awgo"
)

// doUpdate checks for a newer version of the workflow.
func doUpdate() error {
	log.Println("Checking for update...")
	return wf.CheckForUpdate()
}

// checkForUpdate runs "./alsf update" in the background if an update check is due.
func checkForUpdate() error {
	// if !wf.UpdateCheckDue() || aw.IsRunning("update") {
	// 	return nil
	// }
	// cmd := exec.Command(os.Args[0], "update")
	// return aw.RunInBackground("update", cmd)
	return nil
}

// showUpdateStatus adds an "update available!" message to Script Filters if an update is available
// and query is empty.
func showUpdateStatus() {
	if query != "" {
		return
	}

	if wf.UpdateAvailable() {
		wf.Configure(aw.SuppressUIDs(true))
		log.Println("Update available!")
		wf.NewItem("An update is available!").
			Subtitle("⇥ or ↩ to install update").
			Valid(false).
			Autocomplete("workflow:update").
			Icon(updateAvailable)
	}
}
