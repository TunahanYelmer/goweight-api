# GoWeight API

**GoWeight API** is a lightweight Go-based RESTful API that connects to a Bluetooth-enabled scale (BLE) and retrieves weight measurements in real time. This project is designed for integration into health-tracking systems, smart fitness applications, or IoT environments.

---

## 🚀 Features

- 🔌 Scan and connect to BLE (Bluetooth Low Energy) scales
- ⚖️ Read and parse weight data from supported characteristics
- 🌐 REST API endpoints to trigger scans, connect, and fetch data
- 🧩 Modular and extensible Go project structure

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Bluetooth:** [`tinygo-org/bluetooth`](https://github.com/tinygo-org/bluetooth)
- **Web API:** Go standard `net/http`
- **Platform:** Linux/macOS (BLE support required)

---

## 📦 Installation

1. **Clone the repo:**

```bash
git clone https://github.com/yourusername/goweight-api.git
cd goweight-api
