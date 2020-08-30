package main

const banner =
`    ______   __  __   __  ___   ______   ____ 
   / ____/  / / / /  /  |/  /  / ____/  / __ \
  / __/    / /_/ /  / /|_/ /  / / __   / /_/ /
 / /___   / __  /  / /  / /  / /_/ /  / _, _/ 
/_____/  /_/ /_/  /_/  /_/   \____/  /_/ |_|  `

func CreateCustomBanner(description, version string) string {
	return banner + "\n\n" + version + " - " + description + "\n"
}