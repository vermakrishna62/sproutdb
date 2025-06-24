# 🌱 SproutDB

**SproutDB** is a lightweight, command-line based in-memory key-value database written in Go.  
It is designed as a foundational step toward building a deeper understanding of database internals — with future plans to include B+ Tree indexing, persistent storage, and a custom SQL query layer.

---

## 📌 Features

- In-memory key-value store
- Simple CLI interface
- Supports:
  - `SET <key> <value>` — insert or update key
  - `GET <key>` — retrieve value
  - `DELETE <key>` — remove key
- Compiles to native `.exe` — works directly in terminal

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/vermakrishna62/sproutdb.git
cd sproutdb
