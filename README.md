# ArchGen - CLI para Geração de Arquiteturas

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)  
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`archgen` é uma ferramenta de linha de comando (CLI) escrita em Go que automatiza a criação de estruturas de arquiteturas de software. Atualmente, suporta os padrões **MVC** (Model-View-Controller) e **Clean Architecture**, com suporte para as linguagens/frameworks **Go**, **Node.js**, **Python** e **Java**. A CLI permite que desenvolvedores gerem rapidamente uma estrutura inicial de projeto, economizando tempo e garantindo consistência.

## Funcionalidades

- Geração de arquiteturas **MVC** e **Clean Architecture**.
- Suporte para múltiplas linguagens: **Go**, **Node.js**, **Python** e **Java**.
- Personalização do diretório de saída.
- Interface simples e intuitiva via linha de comando.
- Autocompletar para shells (via comando `completion`).

## Pré-requisitos

- [Go](https://golang.org/dl/) (versão 1.21 ou superior) instalado para compilar a CLI (se você for construir do código-fonte).
- O binário pode ser executado em qualquer sistema operacional (Windows, macOS, Linux).

## Instalação

### Via Binário Pré-compilado
1. Baixe o binário mais recente da seção [Releases](https://github.com/seu-usuario/archgen/releases) no GitHub.
2. Adicione o binário ao seu PATH para usá-lo globalmente:
   - **Windows**: Mova o arquivo para `C:\Program Files\archgen` e adicione ao PATH.
   - **Linux/macOS**: Mova para `/usr/local/bin/` com `sudo mv archgen /usr/local/bin/`.
