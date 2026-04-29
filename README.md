# sky2k8s

A **Go-based tooling utility** for deploying and managing applications on Kubernetes (k8s). Simplifies the workflow of pushing services from source to a running Kubernetes cluster.

## ✨ Features

- 🚀 Streamlined deployment pipeline to Kubernetes
- 🐳 Docker image build and push integration
- ☸️ Kubernetes manifest generation and apply
- ⚡ Single command deploy workflow
- 🔧 Configurable via YAML

## 🛠️ Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?style=flat&logo=kubernetes&logoColor=white)

## ⚙️ Requirements

- Go >= 1.13
- Docker
- `kubectl` configured and connected to a cluster

## 🚀 Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/hamdyelbatal122/sky2k8s-master.git
   cd sky2k8s-master
   ```

2. **Build the binary**
   ```bash
   go build -o sky2k8s .
   ```

3. **Run**
   ```bash
   ./sky2k8s --help
   ```

## 📋 Usage

```bash
# Deploy an application to Kubernetes
./sky2k8s deploy --image myapp:latest --namespace production

# Check deployment status
./sky2k8s status --app myapp
```

## 📁 Project Structure

```
sky2k8s-master/
├── main.go          # Entry point
├── cmd/             # CLI commands
├── pkg/             # Core packages
├── config.yaml      # Default configuration
└── README.md
```

## 📄 License

This project is open source and available under the [MIT License](LICENSE).
