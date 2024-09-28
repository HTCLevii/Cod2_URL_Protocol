package main

import (
	"fmt"
	"log"
	"golang.org/x/sys/windows/registry"
)

func main() {
	keyPath := `Software\Wow6432Node\Activision\Call of Duty 2`
	multiEXEString := "MultiEXEString"
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.QUERY_VALUE)
	if err != nil {
		log.Fatalf("Error opening registry key: %v", err)
	}
	defer k.Close()

	exePath, _, err := k.GetStringValue(multiEXEString)
	if err != nil {
		log.Fatalf("Error reading MultiEXEString value: %v", err)
	}
	fmt.Printf("Retrieved executable path: %s\n", exePath)

	baseKeyPath := `cod2`

	cod2Key, _, err := registry.CreateKey(registry.CLASSES_ROOT, baseKeyPath, registry.SET_VALUE)
	if err != nil {
		log.Fatalf("Error creating cod2 key: %v", err)
	}
	defer cod2Key.Close()

	err = cod2Key.SetStringValue("", "URL: cod2 Connect Handler")
	if err != nil {
		log.Fatalf("Error setting default value for cod2: %v", err)
	}

	err = cod2Key.SetStringValue("URL Protocol", "")
	if err != nil {
		log.Fatalf("Error setting URL Protocol value: %v", err)
	}

	iconKey, _, err := registry.CreateKey(registry.CLASSES_ROOT, baseKeyPath+`\DefaultIcon`, registry.SET_VALUE)
	if err != nil {
		log.Fatalf("Error creating DefaultIcon key: %v", err)
	}
	defer iconKey.Close()

	err = iconKey.SetStringValue("", exePath+",0")
	if err != nil {
		log.Fatalf("Error setting DefaultIcon value: %v", err)
	}

	commandKey, _, err := registry.CreateKey(registry.CLASSES_ROOT, baseKeyPath+`\shell\open\command`, registry.SET_VALUE)
	if err != nil {
		log.Fatalf("Error creating command key: %v", err)
	}
	defer commandKey.Close()

	commandValue := fmt.Sprintf("\"%s\" +openurl \"%%1\"", exePath)
	err = commandKey.SetStringValue("", commandValue)
	if err != nil {
		log.Fatalf("Error setting command value: %v", err)
	}

	fmt.Println("Registry updated successfully!")
}
