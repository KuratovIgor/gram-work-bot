package telegram

import "strings"

func searchAreaByName(name string) string {
	for _, area := range AllAreas {
		if strings.Contains(area.Name, name) {
			return area.Id
		}
	}

	return "113"
}

func getScheduleIdByText(schedule string) string {
	switch schedule {
	case scheduleCommands[0]:
		return "fullDay"
	case scheduleCommands[1]:
		return "shift"
	case scheduleCommands[2]:
		return "flexible"
	case scheduleCommands[3]:
		return "remote"
	case scheduleCommands[4]:
		return "flyInFlyOut"
	}

	return "unknown"
}

func getExperienceIdByText(experience string) string {
	switch experience {
	case experienceCommands[0]:
		return "noExperience"
	case experienceCommands[1]:
		return "between1And3"
	case experienceCommands[2]:
		return "between3And6"
	case experienceCommands[3]:
		return "moreThan6"

	}

	return "unknown"
}
