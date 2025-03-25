# 🏎️ Go vs. C# - Benchmarking Database Fetch Performance  

## 📌 Overview  
This project benchmarks the performance of **Go** and **C# (JIT)** when fetching **150 rows** from a **PostgreSQL** database. The goal is to analyze response times, request throughput, and overall efficiency between the two implementations.  

## 📊 Benchmark Results  

| Language | Requests/sec | Avg Latency | Max Latency | Min Latency |
|----------|-------------|-------------|-------------|-------------|
| **Go (8083)** | 557.78 | 14.4ms | 401.3ms | 1.2ms |
| **C# (8081, JIT)** | 1148.70 | 8.2ms | 556.2ms | 1.7ms |

### 🔍 Key Findings  
- **C# (JIT) significantly outperforms Go**, handling **2x more requests per second**.  
- **Go's response time varies more**, with occasional spikes.  
- **C# provides more consistent latency**, leading to predictable performance.  

## 🛠️ Tech Stack  
- **Go** (using `github.com/lib/pq` for PostgreSQL connection)  
- **C# .NET (JIT)**  
- **PostgreSQL** (running locally)  
- **Hey (Load Testing Tool)**  

## 🚀 Setup & Run  

### 🏗️ Prerequisites  
- Go 1.20+  
- .NET 6+  
- PostgreSQL  
- Docker (optional)  

### 🔹 Running the Go Server  
```sh
go mod tidy
go run main.go


Runs on http://localhost:8083/

🔹 Running the C# Server
sh
Copy
Edit
dotnet run
Runs on http://localhost:8081/users

🔹 Running the Load Test
sh
Copy
Edit
hey -n 1000 -c 10 http://localhost:8083  
hey -n 1000 -c 10 http://localhost:8081/users  
📈 Future Improvements
Optimize Go’s database connection pooling.

Tune concurrency handling for both implementations.

Test with different database query complexities.

🏆 Contributing
Feel free to fork the repo and submit PRs with optimizations!

📜 License
MIT License