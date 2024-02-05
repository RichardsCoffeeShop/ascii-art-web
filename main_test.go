package main

import (
	"strings"
	"testing"
)

func RunGenerateASCII(input string, expectedOutput string, t *testing.T) {
	actualOutput := strings.TrimSpace(string(GenerateASCII(input, "standard.txt")))
	expectedOutput = strings.TrimSpace(string(expectedOutput))

	if actualOutput != expectedOutput {
		t.Errorf("Expected:\n%s\n\nBut got:\n%s\n", expectedOutput, actualOutput)
	}
}

func TestRunAllTestCases(t *testing.T) {
	testInputs := []string{"hello", "HELLO", "HeLlo HuMaN", "0123456789", "\\!\" #$%&'\"'\"'()*+,-./"}
	expectedOutputs := []string{
		`| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
	`,
		`| |  | | |  ____| | |      | |       / __ \  
| |__| | | |__    | |      | |      | |  | | 
|  __  | |  __|   | |      | |      | |  | | 
| |  | | | |____  | |____  | |____  | |__| | 
|_|  |_| |______| |______| |______|  \____/  
`,
		`| |  | |        | |      | |               | |  | |         |  \/  |         | \ | | 
| |__| |   ___  | |      | |   ___         | |__| |  _   _  | \  / |   __ _  |  \| | 
|  __  |  / _ \ | |      | |  / _ \        |  __  | | | | | | |\/| |  / _` + "`" + ` | | . ` + "`" + ` | 
| |  | | |  __/ | |____  | | | (_) |       | |  | | | |_| | | |  | | | (_| | | |\  | 
|_|  |_|  \___| |______| |_|  \___/        |_|  |_|  \__,_| |_|  |_|  \__,_| |_| \_| 
`,
		`___    _   ____    _____   _  _     ____     __     _____    ___     ___   
 / _ \  / | |___ \  |___ /  | || |   | ___|   / /    |___  |  ( _ )   / _ \  
| | | | | |   __) |   |_ \  | || |_  |___ \  | '_ \     / /   / _ \  | (_) | 
| |_| | | |  / __/   ___) | |__   _|   __) | | (_) |   / /   | (_) |  \__, | 
 \___/  |_| |_____| |____/     |_|   |____/   \___/   /_/     \___/     / /  
                                                                       /_/`,
		`\ \     | | ( | )        _| || |_   | |  (_) / /   ___    ( ) ( | ) ( ) ( | ) ( )  / / \ \   /\| |/\     _                         / / 
 \ \    | |  V V        |_  __  _| / __)    / /   ( _ )   |/   V V  |/   V V  |/  | |   | |  \ ` + "`" + ` ' /   _| |_       ______         / /  
  \ \   | |              _| || |_  \__ \   / /    / _ \/\                         | |   | | |_     _| |_   _|     |______|       / /   
   \ \  |_|             |_  __  _| (   /  / / _  | (_>  <                         | |   | |  / , . \    |_|    _            _   / /    
    \_\ (_)               |_||_|    |_|  /_/ (_)  \___/\/                         | |   | |  \/|_|\/          ( )          (_) /_/     
                                                                                   \_\ /_/                    |/`,
	}

	for i, input := range testInputs {
		RunGenerateASCII(input, expectedOutputs[i], t)
	}
}
