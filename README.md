***Test Task***

How to run? 
1. Clone the repository
2. Run the following command in the terminal
```bash
go mod tidy
```
3. Fill the config.yaml file with the token and the url

4. Run the following command in the terminal
```bash
go run cmd/server/main.go
```
5. You are ready to test the server!
```bash
curl -X POST http://localhost:8080/buffer \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "period_start=2024-05-01" \
  -d "period_end=2024-05-31" \
  -d "period_key=month" \
  -d "indicator_to_mo_id=227373" \
  -d "indicator_to_mo_fact_id=0" \
  -d "value=1" \
  -d "fact_time=2024-05-31" \
  -d "is_plan=0" \
  -d "auth_user_id=40" \
  -d "comment=buffer Last_name"
```