package main

import (
	"encoding/json"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(generateMarksCmd)
}

// markCmd represents the mark command
var generateMarksCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"purge"},
	Short:   "Generate marks from configuration preferences",
	Long:    `Generate marks from configuration preferences`,
	Run: func(_ *cobra.Command, _ []string) {
		generateMarks()
	},
}

// utility function to pull attributes out from a map
func pluck(key string, input any) interface{} {
	obj := input
	val, ok := obj.(map[string]interface{})[key]
	if !ok {
		log.Tracef("Warning: key %v not found in %v", key, obj)
		return nil
	}
	return val
}

func generateMarks() {
	_, windowsList := getWindows()
	preferences := viper.Get("preferences").(map[string]interface{})
	log.Debugf("Preferences: %+v", preferences)

	for i, windowObj := range windowsList {
		windowName := pluck("app", windowObj).(string)
		windowTitle := pluck("title", windowObj).(string)
		log.Debugf("Considering window %v, %+v", i, windowName)
		for prefName := range preferences {
			if !strings.Contains(strings.ToLower(windowName), strings.ToLower(prefName)) {
				continue
			}
			prefMark := pluck("mark", preferences[prefName]).(string)
			if prefTitle := pluck("title", preferences[prefName]); prefTitle != nil {
				t := prefTitle.(string)
				if !strings.Contains(strings.ToLower(windowTitle), strings.ToLower(t)) {
					continue
				}
			}
			log.Debugf("  [>] Found preferred window %v (app %v), applying mark %v", windowTitle, windowName, prefMark)
			jsonBA, err := json.Marshal(windowObj)
			if err != nil {
				log.Warningf("Warning: Error converting window obj %v to json string", windowObj)
			}
			m := NewWinMarker()
			if err := m.SaveMark([]rune(prefMark)[0], jsonBA); err != nil {
				log.Warningf("Warning: Error saving mark %v: %v", prefMark, err)
			}
		}
	}
}
