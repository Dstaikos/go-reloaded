# Go-Reloaded 🚀

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Tests](https://img.shields.io/badge/Tests-Passing-brightgreen?style=for-the-badge)
![Build](https://img.shields.io/badge/Build-Passing-success?style=for-the-badge)
![Architecture](https://img.shields.io/badge/Architecture-Pipeline-blue?style=for-the-badge)

A powerful text processing and auto-correction tool written in Go that applies intelligent transformations to text files. ✨

## ✨ Features

- 🔢 **Number Base Conversion** - Convert hexadecimal and binary numbers to decimal
- 🔤 **Text Case Transformation** - Apply uppercase, lowercase, and capitalization rules
- 📝 **Smart Punctuation** - Automatic spacing and formatting correction
- 💬 **Quote Normalization** - Proper single quote pairing and spacing
- 📖 **Article Correction** - Intelligent "a" to "an" replacements
- ⚡ **Batch Processing** - Apply transformations to multiple words at once

## 🚀 Quick Start

### 📋 Prerequisites
- Go 1.21 or higher

### 📦 Installation

```bash
git clone https://github.com/yourusername/go-reloaded.git
cd go-reloaded
```

### 🎯 Usage

```bash
cd cmd
go run main.go input.txt output.txt
```

## 💡 Examples

### 🔢 Number Conversions
```
Input:  "1E (hex) files were added"
Output: "30 files were added"

Input:  "It has been 10 (bin) years"
Output: "It has been 2 years"
```

### 🔤 Text Transformations
```
Input:  "Ready, set, go (up)!"
Output: "Ready, set, GO!"

Input:  "This is so exciting (up, 2)"
Output: "This is SO EXCITING"
```

### ✨ Smart Corrections
```
Input:  "I was sitting over there ,and then BAMM !!"
Output: "I was sitting over there, and then BAMM!!"

Input:  "There it was. A amazing rock!"
Output: "There it was. An amazing rock!"
```

## 📁 Project Structure

```
go-reloaded/
├── cmd/
│   └── main.go          # Entry point
├── pkg/                 # Core functionality
├── docs/                # Documentation
│   ├── analysis.md      # Detailed specifications
│   └── Pipeline-FSM comparison.md
├── tasks/               # Development process (Zone01 requirement)
└── README.md
```

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

## 📚 Documentation

For detailed transformation rules and technical specifications, see [docs/analysis.md](docs/analysis.md).

## 🛠️ Development Process

This project follows the Zone01 curriculum requirements. Development steps and process documentation can be found in the `tasks/` folder.

## 🏗️ Architecture

This project uses a **Pipeline Architecture** for clear separation of concerns and maintainability. Each transformation stage processes the text sequentially, making the code easy to test, debug, and extend.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Built with Go** 🐹 • **Pipeline Architecture** 🔧 • **Comprehensive Testing** ✅