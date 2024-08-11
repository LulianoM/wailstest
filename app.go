package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GenerateProject(projectName string, projectTeam string) string {
	// Caminho para o arquivo placeholder
	placeholderPath := "./placeholder/"

	// Ler o arquivo placeholder
	content, err := os.ReadFile(placeholderPath)
	if err != nil {
		return fmt.Sprintf("Erro ao ler o arquivo: %s", err.Error())
	}

	// Conteúdo do arquivo como string
	contentStr := string(content)

	// Modificar o conteúdo, substituindo {{module_name}} por empresa/projectName
	moduleName := fmt.Sprintf("empresa/%s", projectName)
	modifiedContent := strings.ReplaceAll(contentStr, "{{module_name}}", moduleName)

	// Nome do novo arquivo
	newFileName := fmt.Sprintf("./%s_%s_module.txt", projectName, projectTeam)

	// Escrever o conteúdo modificado em um novo arquivo
	err = os.WriteFile(newFileName, []byte(modifiedContent), 0644)
	if err != nil {
		return fmt.Sprintf("Erro ao escrever o arquivo: %s", err.Error())
	}

	return fmt.Sprintf("Projeto criado com sucesso: %s", newFileName)
}
